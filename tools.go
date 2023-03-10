package utils

func LowBit[T Integer](x T) T {
	return x & -x
}

func Next2[T Integer](x T) T {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	x++
	return x
}
func Next2Int(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	x++
	return x
}
func SizeVarInt(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func Log2[I Integer](x I) I {
	if x == 0 {
		return 0
	}
	return Log2(x>>1) + 1
}
func Log2Int(x int) int {
	if x == 0 {
		return 0
	}
	return Log2Int(x>>1) + 1
}
