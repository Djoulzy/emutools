package mem_v2

type MEMAccess interface {
	MRead([]MEMCell, uint16) byte
	MWrite([]MEMCell, uint16, byte)
	MWriteUnder([]MEMCell, uint16, byte)
}

type MEMCell struct {
	LayerNum      int
	UnderLayerNum int
	Val           *byte
	Under         *byte
	Accessor      MEMAccess
	UnderAccess   MEMAccess
}

func GetMemoryManager(nbMemLayout int, layoutsSize uint, bankSelector *byte) *BANK {
	return Init(nbMemLayout, layoutsSize, bankSelector)
}