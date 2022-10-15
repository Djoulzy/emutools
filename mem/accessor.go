package mem

type DefaultAccessor struct {

}

func (DA *DefaultAccessor) MRead(mem []byte, translatedAddr uint16) byte {
	// clog.Test("MEM", "MRead", "Addr: %04X -> %02X", translatedAddr, mem[translatedAddr])
	return mem[translatedAddr]
}

func (DA *DefaultAccessor) MWrite(mem []byte, translatedAddr uint16, val byte) {
	// clog.Test("MEM", "MWrite", "Addr: %04X -> %02X", addr, val)
	mem[translatedAddr] = val
}