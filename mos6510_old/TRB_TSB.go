package mos6510

///////////////////////////////////////////////////////
//                        TRB                        //
///////////////////////////////////////////////////////

func (C *CPU) TRB_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
		C.updateZ(C.A & C.tmpBuff)
	case 4:
		C.tmpBuff &= ^C.A
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) TRB_abs() {
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
	case 5:
		C.tmpBuff &= ^C.A
	case 6:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        TSB                        //
///////////////////////////////////////////////////////

func (C *CPU) TSB_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
		C.updateZ(C.A & C.tmpBuff)
	case 4:
		C.tmpBuff |= C.A
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) TSB_abs() {
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
	case 5:
		C.tmpBuff |= C.A
	case 6:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}