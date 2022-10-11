package mos6510

///////////////////////////////////////////////////////
//                        BCC                        //
///////////////////////////////////////////////////////

func (C *CPU) BCC_rel() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.instCode = C.ram.Read(C.PC)
		if !C.issetC() {
			dest := uint16(int(C.PC) + int(int8(C.OperLO)))
			if C.PC&0xFF00 != dest&0xFF00 {
				C.pageCrossed = true
				C.Inst.Cycles++
			} else {
				C.pageCrossed = false
				C.PC = dest
			}
		} else {
			C.CycleCount = 1
			C.firstCycle()
		}
	case 4:
		C.instCode = C.ram.Read(C.PC)
		if C.pageCrossed {
			C.PC = uint16(int(C.PC) + int(int8(C.OperLO)))
		} else {
			C.CycleCount = 1
			C.firstCycle()
		}
	case 5:
		C.CycleCount = 1
		C.firstCycle()
	}
}

///////////////////////////////////////////////////////
//                        BCS                        //
///////////////////////////////////////////////////////

func (C *CPU) BCS_rel() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.instCode = C.ram.Read(C.PC)
		if C.issetC() {
			dest := uint16(int(C.PC) + int(int8(C.OperLO)))
			if C.PC&0xFF00 != dest&0xFF00 {
				C.pageCrossed = true
				C.Inst.Cycles++
			} else {
				C.pageCrossed = false
				C.PC = dest
			}
		} else {
			C.CycleCount = 1
			C.firstCycle()
		}
	case 4:
		C.instCode = C.ram.Read(C.PC)
		if C.pageCrossed {
			C.PC = uint16(int(C.PC) + int(int8(C.OperLO)))
		} else {
			C.CycleCount = 1
			C.firstCycle()
		}
	case 5:
		C.CycleCount = 1
		C.firstCycle()
	}

}
