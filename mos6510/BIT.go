package mos6510

///////////////////////////////////////////////////////
//                        BIT                        //
///////////////////////////////////////////////////////

func (C *CPU) BIT_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
		C.updateZ(C.A & C.tmpBuff)
		C.setV(C.tmpBuff&0b01000000 == 0b01000000)
		C.setN(C.tmpBuff&0b10000000 == 0b10000000)
		C.CycleCount = 0
	}
}

func (C *CPU) BIT_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC++
	case 4:
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateZ(C.A & C.tmpBuff)
		C.setV(C.tmpBuff&0b01000000 == 0b01000000)
		C.setN(C.tmpBuff&0b10000000 == 0b10000000)
		C.CycleCount = 0
	}
}
