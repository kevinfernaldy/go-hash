package sha224

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
