package sha512

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
