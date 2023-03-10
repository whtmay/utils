package gbytes

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const (
	defaultBufSize = 16
	bit7           = 1 << 7
)

type Buffer struct {
	readPos  int
	writePos int
	capacity int
	buf      []byte
}

func New() *Buffer {
	return &Buffer{
		capacity: defaultBufSize,
		buf:      make([]byte, defaultBufSize),
	}
}
func NewBuffer(arr []byte, rw ...int) *Buffer {
	buf := &Buffer{
		capacity: len(arr),
		buf:      arr,
	}
	if len(rw) == 1 {
		buf.writePos = rw[0]
	} else if len(rw) == 2 {
		buf.writePos, buf.readPos = rw[0], rw[1]
	}
	return buf
}
func (b *Buffer) ReadPos() int {
	return b.readPos
}
func (b *Buffer) WritePos() int {
	return b.writePos
}
func (b *Buffer) Capacity() int {
	return b.capacity
}
func (b *Buffer) ReadPosGrow(n int) {
	b.readPos += n
}
func (b *Buffer) WritePosGrow(n int) {
	b.writePos += n
}
func (b *Buffer) ReSetReadPos() {
	b.readPos = 0
}
func (b *Buffer) ReSetWritePos() {
	b.writePos = 0
}
func (b *Buffer) CanRead() int {
	return b.writePos - b.readPos
}
func (b *Buffer) CanWrite() int {
	return b.capacity - b.writePos
}
func (b *Buffer) Bytes() []byte {
	return b.buf
}
func (b *Buffer) tryGrow(n int) bool {
	if n > b.CanWrite() {
		return false
	}
	return true
}
func (b *Buffer) grow() {
	capacity := 0
	if b.buf == nil {
		b.buf = make([]byte, defaultBufSize)
		return
	}
	if b.Capacity() <= 1024 {
		capacity = utils.Next2(b.Capacity() * 2)
	} else {
		capacity = b.Capacity() + b.Capacity()/4
	}
	data := b.buf[b.readPos:b.writePos]
	b.buf = make([]byte, capacity)
	copy(b.buf, data)
	b.writePos -= b.readPos
	b.readPos = 0
	b.capacity = capacity

}
func (b *Buffer) WriteBytes(data []byte) {
	if !b.tryGrow(len(data)) {
		b.grow()
	}
	b.writePos += copy(b.buf[b.writePos:], data)
}
func (b *Buffer) ReadBytes(n int) []byte {
	if n == -1 {
		e := b.buf[b.readPos:b.writePos]
		b.readPos = b.writePos
		return e
	}
	if b.CanRead() < n {
		panic(fmt.Sprintf("canread has %d,but will read %d", b.CanRead(), n))
	}
	b.readPos += n
	return b.buf[b.readPos-n : b.readPos]
}
func (b *Buffer) ReadLastBytes(n int) []byte {
	if b.CanRead() < n {
		panic(fmt.Sprintf("canread has %d,but will read %d", b.CanRead(), n))
	}
	return b.buf[b.writePos-n : b.writePos]
}

func (b *Buffer) WriteVarUInt8(x uint8) int {
	return b.writeVarUInt(uint64(x))
}
func (b *Buffer) WriteVarUInt16(x uint16) int {
	return b.writeVarUInt(uint64(x))
}
func (b *Buffer) WriteVarUInt32(x uint32) int {
	return b.writeVarUInt(uint64(x))
}
func (b *Buffer) WriteVarUInt64(x uint64) int {
	return b.writeVarUInt(x)
}
func (b *Buffer) writeVarUInt(x uint64) int {
	if !b.tryGrow(12) {
		b.grow()
	}
	off := b.writePos
	b.writePos += binary.PutUvarint(b.buf[off:], x)
	return b.writePos - off
}
func (b *Buffer) WriteVarInt(x int) int {
	return b.writeVarInt(int64(x))
}
func (b *Buffer) WriteVarInt8(x int8) int {
	return b.writeVarInt(int64(x))
}
func (b *Buffer) WriteVarInt16(x int16) int {
	return b.writeVarInt(int64(x))
}
func (b *Buffer) WriteVarInt32(x int32) int {
	return b.writeVarInt(int64(x))
}
func (b *Buffer) WriteVarInt64(x int64) int {
	return b.writeVarInt(x)
}
func (b *Buffer) writeVarInt(x int64) int {
	if !b.tryGrow(12) {
		b.grow()
	}
	off := b.writePos
	b.writePos += binary.PutVarint(b.buf[off:], x)
	return b.writePos - off
}
func (b *Buffer) ReadVarUInt8() uint8 {
	return uint8(b.readVarUInt())
}
func (b *Buffer) ReadVarUInt16() uint16 {
	return uint16(b.readVarUInt())
}
func (b *Buffer) ReadVarUInt32() uint32 {
	return uint32(b.readVarUInt())
}
func (b *Buffer) ReadVarUInt64() uint64 {
	return b.readVarUInt()
}
func (b *Buffer) readVarUInt() uint64 {
	x, sz := binary.Uvarint(b.buf[b.readPos:])
	b.readPos += sz
	return x
}
func (b *Buffer) ReadVarInt8() int8 {
	return int8(b.readVarInt())
}
func (b *Buffer) ReadVarInt16() int16 {
	return int16(b.readVarInt())
}
func (b *Buffer) ReadVarInt32() int32 {
	return int32(b.readVarInt())
}
func (b *Buffer) ReadVarInt64() int64 {
	return b.readVarInt()
}
func (b *Buffer) readVarInt() int64 {
	x, sz := binary.Varint(b.buf[b.readPos:])
	b.readPos += sz
	return x
}

func (b *Buffer) WriteUInt8(x uint8) {
	if !b.tryGrow(8) {
		b.grow()
	}
	b.writePos++
	b.buf[b.writePos-1] = x
}
func (b *Buffer) WriteUInt16(x uint16) {
	if !b.tryGrow(8) {
		b.grow()
	}
	b.writePos += 2
	binary.BigEndian.PutUint16(b.buf[b.writePos-2:], x)
}
func (b *Buffer) WriteUInt32(x uint32) {
	if !b.tryGrow(8) {
		b.grow()
	}
	b.writePos += 4
	binary.BigEndian.PutUint32(b.buf[b.writePos-4:], x)
}
func (b *Buffer) WriteUInt64(x uint64) {
	if !b.tryGrow(8) {
		b.grow()
	}
	b.writePos += 8
	binary.BigEndian.PutUint64(b.buf[b.writePos-8:], x)
}
func (b *Buffer) WriteUInt(x uint) {
	if !b.tryGrow(8) {
		b.grow()
	}
	if unsafe.Sizeof(uint(0)) == 4 {
		b.WriteUInt32(uint32(x))
	}
	b.WriteUInt64(uint64(x))
}
func (b *Buffer) ReadUInt8() uint8 {
	b.readPos++
	return b.buf[b.readPos-1]
}
func (b *Buffer) ReadUInt16() uint16 {
	b.readPos += 2
	return binary.BigEndian.Uint16(b.buf[b.readPos-2:])
}
func (b *Buffer) ReadUInt32() uint32 {
	b.readPos += 4
	return binary.BigEndian.Uint32(b.buf[b.readPos-4:])
}
func (b *Buffer) ReadUInt64() uint64 {
	b.readPos += 8
	return binary.BigEndian.Uint64(b.buf[b.readPos-8:])
}
func (b *Buffer) ReadUInt() int {
	if unsafe.Sizeof(uint(0)) == 4 {
		return int(b.ReadUInt32())
	}
	return int(b.ReadUInt64())
}
func (b *Buffer) ReadString(n int) string {
	if b.CanRead() < n {
		panic(fmt.Sprintf("canread has %d,but will read %d", b.CanRead(), 1))
	}
	b.readPos += n
	return string(b.buf[b.readPos-n : b.readPos])
}
