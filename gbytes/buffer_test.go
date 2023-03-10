package gbytes

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {

	b := New()
	b.WriteVarUInt32(uint32(32095))
	b.WriteVarUInt16(uint16(32095))
	b.WriteVarUInt8(uint8(56))
	b.WriteBytes([]byte("的方法"))
	b.WriteUInt32(uint32(9032))
	b.writeVarInt(32243)
	fmt.Println(b.ReadVarUInt16())
	fmt.Println(b.ReadVarUInt16())
	fmt.Println(b.ReadUInt8())
	fmt.Println(string(b.ReadBytes(9)))
	fmt.Println(b.ReadUInt32())
	fmt.Println(b.readVarInt())
	b.ReSetReadPos()
	fmt.Println(b.ReadBytes(-1))
}
