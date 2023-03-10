package utils

type Int interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}
type UInt interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type Integer interface {
	Int | UInt
}

