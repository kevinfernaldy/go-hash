package md5

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/kevinfernaldy/go-hash/internal/constant"
)

type MD5 struct {
	isUsed  bool
	payload []byte
	hash    []byte
}

func (md5 *MD5) appendMessage(payload []byte, length int) []byte {
	payload = append(payload, 0x80)

	for len(payload)%64 != 56 {
		payload = append(payload, 0x00)
	}

	payload = binary.LittleEndian.AppendUint64(payload, uint64(length*8))

	return payload
}

func (md5 *MD5) Update(payload string) error {
	if md5.isUsed {
		return errors.New("hash is already used and cannot be updated")
	}

	md5.payload = []byte(payload)

	return nil
}

func (md5 *MD5) Digest() string {
	if md5.isUsed {
		return hex.EncodeToString(md5.hash)
	}

	md5.isUsed = true

	message := md5.payload

	// Append message
	message = md5.appendMessage(message, len(message))

	iterations := len(message) / 64
	a0 := constant.MD5_A
	b0 := constant.MD5_B
	c0 := constant.MD5_C
	d0 := constant.MD5_D

	for i := 0; i < iterations; i++ {
		blockSlice := message[i*64 : (i+1)*64]

		chunks := []uint32{}

		for j := 0; j < 16; j++ {
			chunks = append(chunks, binary.LittleEndian.Uint32(blockSlice[j*4:(j+1)*4]))
		}

		A := a0
		B := b0
		C := c0
		D := d0
		for j := 0; j < 64; j++ {
			var F, g uint32

			switch {
			case 0 <= j && j <= 15:
				{
					F = (B & C) | (^B & D)
					g = uint32(j)
				}
			case 16 <= j && j <= 31:
				{
					F = (D & B) | (^D & C)
					g = uint32((5*j + 1) % 16)
				}
			case 32 <= j && j <= 47:
				{
					F = B ^ C ^ D
					g = uint32((3*j + 5) % 16)
				}
			case 48 <= j && j <= 63:
				{
					F = C ^ (B | ^D)
					g = uint32((7 * j) % 16)
				}
			}

			F += A + constant.MD5_K[j] + chunks[g]
			A = D
			D = C
			C = B
			B += rotl(F, constant.MD5_s[j])
		}

		a0 += A
		b0 += B
		c0 += C
		d0 += D
	}

	md5.hash = binary.LittleEndian.AppendUint32(md5.hash, a0)
	md5.hash = binary.LittleEndian.AppendUint32(md5.hash, b0)
	md5.hash = binary.LittleEndian.AppendUint32(md5.hash, c0)
	md5.hash = binary.LittleEndian.AppendUint32(md5.hash, d0)

	return hex.EncodeToString(md5.hash)
}
