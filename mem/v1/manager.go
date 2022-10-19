package mem

type MEMAccess interface {
	MRead([]byte, uint16) byte
	MWrite([]byte, uint16, byte)
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
