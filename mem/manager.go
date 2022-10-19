package mem

import (
	"mem_v1"
	"mem_v2"
)

type Manager interface {
	Attach(int, string, uint16, []byte, bool, bool, interface{})

	Read(uint16) byte
	Write(uint16, byte)
	GetFullSize() int
	GetStack(uint16, uint16) []byte
	Enable(string)
	Disable(string)
	ReadOnly(string)
	ReadWrite(string)

	Clear([]byte, int, byte)
	LoadROM(int, string) []byte
	Dump(uint16)
}

func GetMemoryManager_v2(nbMemLayout int, layoutsSize uint, bankSelector *byte) interface{} {
	return mem_v2.Init(nbMemLayout, layoutsSize, bankSelector)
}

func GetMemoryManager_v1(nbMemLayout int, layoutsSize uint, bankSelector *byte) interface{} {
	return mem_v1.Init(nbMemLayout, layoutsSize, bankSelector)
}
