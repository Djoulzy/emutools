package mem_v1

type DefaultAccessor struct {

}

func (DA *DefaultAccessor) MRead(mem MEMCells, translatedAddr uint16) byte {
	// clog.Test("MEM", "MRead", "Addr: %04X -> %02X", translatedAddr, mem[translatedAddr])
	cells := mem.([]byte)
	return byte(cells[translatedAddr])
}

func (DA *DefaultAccessor) MWrite(mem MEMCells, translatedAddr uint16, val byte) {
	// clog.Test("MEM", "MWrite", "Addr: %04X -> %02X", addr, val)
	cells := mem.([]byte)
	cells[translatedAddr] = val
}