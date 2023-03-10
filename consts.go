package utils

import (
	"hash/crc32"
)

const (
	B  uint64 = iota + 1
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
)

var CastagnoliTable = crc32.MakeTable(crc32.Castagnoli)


