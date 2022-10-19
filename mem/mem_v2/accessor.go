package mem_v2

type DefaultAccessor struct {

}

func (DA *DefaultAccessor) MRead(mem interface{}, addr uint16) byte {
	cells := mem.([]MEMCell)
	return *cells[addr].Val
}

func (DA *DefaultAccessor) MWrite(mem interface{}, addr uint16, value byte) {
	cells := mem.([]MEMCell)
	*cells[addr].Val = value
}

func (DA *DefaultAccessor) MWriteUnder(mem interface{}, addr uint16, value byte) {
	cells := mem.([]MEMCell)
	*cells[addr].Under = value
}