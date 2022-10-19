package mem

type MEMAccess interface {
	MRead([]MEMCell, uint16) byte
	MWrite([]MEMCell, uint16, byte)
}

type MEMCell struct {
	Val *byte
}

func GetMemoryManager(nbMemLayout int, layoutsSize uint, bankSelector *byte) *BANK {
	return Init(nbMemLayout, layoutsSize, bankSelector)
}
