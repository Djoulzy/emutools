package mem

type DefaultAccessor struct {

}

func (DA *DefaultAccessor) MRead(mem []MEMCell, translatedAddr uint16) byte {
	// clog.Test("MEM", "MRead", "Addr: %04X -> %02X", translatedAddr, mem[translatedAddr])
	return *mem[translatedAddr].Val
}

func (DA *DefaultAccessor) MWrite(mem []MEMCell, translatedAddr uint16, val byte) {
	// clog.Test("MEM", "MWrite", "Addr: %04X -> %02X", addr, val)
	*mem[translatedAddr].Val = val
}

func (DA *DefaultAccessor) MWriteUnder(mem []MEMCell, addr uint16, value byte) {
}