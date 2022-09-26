package sha224

import "github.com/kevinfernaldy/go-hash/internal/typedef"

func shr(x typedef.Word32, n uint) typedef.Word32 {
	return x >> n
}

func rotr(x typedef.Word32, n uint) typedef.Word32 {
	return (x >> n) | (x << (typedef.Word32Size - n))
}

func rotl(x typedef.Word32, n uint) typedef.Word32 {
	return (x << n) | (x >> (typedef.Word32Size - n))
}

func ch(x typedef.Word32, y typedef.Word32, z typedef.Word32) typedef.Word32 {
	return (x & y) ^ (^x & z)
}

func maj(x typedef.Word32, y typedef.Word32, z typedef.Word32) typedef.Word32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func bsig0(x typedef.Word32) typedef.Word32 {
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func bsig1(x typedef.Word32) typedef.Word32 {
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func ssig0(x typedef.Word32) typedef.Word32 {
	return rotr(x, 7) ^ rotr(x, 18) ^ shr(x, 3)
}

func ssig1(x typedef.Word32) typedef.Word32 {
	return rotr(x, 17) ^ rotr(x, 19) ^ shr(x, 10)
}
