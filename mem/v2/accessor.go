package mem_v2

type DefaultAccessor struct {

}

func (DA *DefaultAccessor) MRead(mem []MEMCell, addr uint16) byte {
	return *mem[addr].Val
}

func (DA *DefaultAccessor) MWrite(mem []MEMCell, addr uint16, value byte) {
	*mem[addr].Val = value
}

func (DA *DefaultAccessor) MWriteUnder(mem []MEMCell, addr uint16, value byte) {
	*mem[addr].Under = value
}