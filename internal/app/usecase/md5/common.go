package md5

func rotl(x uint32, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}
