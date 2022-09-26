package sha512

import "github.com/kevinfernaldy/go-hash/internal/typedef"

func shr(x typedef.Word64, n uint) typedef.Word64 {
	return x >> n
}

func rotr(x typedef.Word64, n uint) typedef.Word64 {
	return (x >> n) | (x << (typedef.Word64Size - n))
}

func rotl(x typedef.Word64, n uint) typedef.Word64 {
	return (x << n) | (x >> (typedef.Word64Size - n))
}

func ch(x typedef.Word64, y typedef.Word64, z typedef.Word64) typedef.Word64 {
	return (x & y) ^ (^x & z)
}

func maj(x typedef.Word64, y typedef.Word64, z typedef.Word64) typedef.Word64 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func bsig0(x typedef.Word64) typedef.Word64 {
	return rotr(x, 28) ^ rotr(x, 34) ^ rotr(x, 39)
}

func bsig1(x typedef.Word64) typedef.Word64 {
	return rotr(x, 14) ^ rotr(x, 18) ^ rotr(x, 41)
}

func ssig0(x typedef.Word64) typedef.Word64 {
	return rotr(x, 1) ^ rotr(x, 8) ^ shr(x, 7)
}

func ssig1(x typedef.Word64) typedef.Word64 {
	return rotr(x, 19) ^ rotr(x, 61) ^ shr(x, 6)
}
