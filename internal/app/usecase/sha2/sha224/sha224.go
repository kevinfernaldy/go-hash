package sha224

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/kevinfernaldy/go-hash/internal/constant"
)

func shr(x uint32, n uint) uint32 {
	return x >> n
}

func rotr(x uint32, n uint) uint32 {
	return (x >> n) | (x << (32 - n))
}

func rotl(x uint32, n uint) uint32 {
	return (x << n) | (x >> (32 - n))
}

func ch(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func maj(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func bsig0(x uint32) uint32 {
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func bsig1(x uint32) uint32 {
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func ssig0(x uint32) uint32 {
	return rotr(x, 7) ^ rotr(x, 18) ^ shr(x, 3)
}

func ssig1(x uint32) uint32 {
	return rotr(x, 17) ^ rotr(x, 19) ^ shr(x, 10)
}

func (sha *SHA224) appendMessage(payload []byte, length int) []byte {
	payload = append(payload, 0x80)

	for len(payload)%64 != 56 {
		payload = append(payload, 0x00)
	}

	payload = binary.BigEndian.AppendUint64(payload, uint64(length*8))

	return payload
}

func (sha *SHA224) Update(payload string) error {
	if sha.isUsed {
		return errors.New("hash is already used and cannot be updated")
	}

	sha.payload = []byte(payload)

	return nil
}

func (sha *SHA224) Digest() string {
	if sha.isUsed {
		return hex.EncodeToString(sha.hash)
	}
	sha.isUsed = true

	message := sha.payload

	// Append message
	message = sha.appendMessage(message, len(message))

	iterations := len(message) / 64
	H := constant.SHA224_H

	for i := 0; i < iterations; i++ {
		blockSlice := message[i*64 : (i+1)*64]
		W := [64]uint32{}

		// Prepare message schedule
		for t := 0; t < 16; t++ {
			W[t] = uint32(binary.BigEndian.Uint32(blockSlice[t*4 : (t+1)*4]))
		}

		for t := 16; t < 64; t++ {
			W[t] = ssig1(W[t-2]) + W[t-7] + ssig0(W[t-15]) + W[t-16]
		}

		// Initialize working variables
		a := H[0]
		b := H[1]
		c := H[2]
		d := H[3]
		e := H[4]
		f := H[5]
		g := H[6]
		h := H[7]

		// Perform hash computation
		for t := 0; t < 64; t++ {
			T1 := h + bsig1(e) + ch(e, f, g) + constant.SHA224_K[t] + W[t]
			T2 := bsig0(a) + maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
		}

		H[0] = a + H[0]
		H[1] = b + H[1]
		H[2] = c + H[2]
		H[3] = d + H[3]
		H[4] = e + H[4]
		H[5] = f + H[5]
		H[6] = g + H[6]
		H[7] = h + H[7]
	}

	// Convert the WORD array to byte array
	for i := 0; i < 7; i++ {
		sha.hash = binary.BigEndian.AppendUint32(sha.hash, uint32(H[i]))
	}

	// Return hexadecimal representation
	return hex.EncodeToString(sha.hash)
}
