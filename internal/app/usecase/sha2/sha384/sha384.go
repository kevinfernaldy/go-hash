package sha384

import (
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/kevinfernaldy/go-hash/internal/constant"
)

func shr(x uint64, n uint) uint64 {
	return x >> n
}

func rotr(x uint64, n uint) uint64 {
	return (x >> n) | (x << (64 - n))
}

func rotl(x uint64, n uint) uint64 {
	return (x << n) | (x >> (64 - n))
}

func ch(x uint64, y uint64, z uint64) uint64 {
	return (x & y) ^ (^x & z)
}

func maj(x uint64, y uint64, z uint64) uint64 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func bsig0(x uint64) uint64 {
	return rotr(x, 28) ^ rotr(x, 34) ^ rotr(x, 39)
}

func bsig1(x uint64) uint64 {
	return rotr(x, 14) ^ rotr(x, 18) ^ rotr(x, 41)
}

func ssig0(x uint64) uint64 {
	return rotr(x, 1) ^ rotr(x, 8) ^ shr(x, 7)
}

func ssig1(x uint64) uint64 {
	return rotr(x, 19) ^ rotr(x, 61) ^ shr(x, 6)
}

func (sha *SHA384) appendMessage(payload []byte, length int) []byte {
	payload = append(payload, 0x80)

	for len(payload)%128 != 112 {
		payload = append(payload, 0x00)
	}

	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)
	payload = append(payload, 0x00)

	payload = binary.BigEndian.AppendUint64(payload, uint64(length*8))

	return payload
}

func (sha *SHA384) Update(payload string) error {
	if sha.isUsed {
		return errors.New("hash is already used and cannot be updated")
	}

	sha.payload = []byte(payload)

	return nil
}

func (sha *SHA384) Digest() string {
	if sha.isUsed {
		return hex.EncodeToString(sha.hash)
	}
	sha.isUsed = true

	message := sha.payload

	// Append message
	message = sha.appendMessage(message, len(message))

	iterations := len(message) / 128
	H := constant.SHA384_H

	for i := 0; i < iterations; i++ {
		blockSlice := message[i*128 : (i+1)*128]
		W := [80]uint64{}

		// Prepare message schedule
		for t := 0; t < 16; t++ {
			W[t] = uint64(binary.BigEndian.Uint64(blockSlice[t*8 : (t+1)*8]))
		}

		for t := 16; t < 80; t++ {
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
		for t := 0; t < 80; t++ {
			T1 := h + bsig1(e) + ch(e, f, g) + constant.SHA384_K[t] + W[t]
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
	for i := 0; i < 6; i++ {
		sha.hash = binary.BigEndian.AppendUint64(sha.hash, uint64(H[i]))
	}

	// Return hexadecimal representation
	return hex.EncodeToString(sha.hash)
}
