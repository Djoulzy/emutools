package mos6510_2

///////////////////////////////////////////////////////
//                        BMI                        //
///////////////////////////////////////////////////////

func (C *CPU) BMI_rel() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.instCode = C.ram.Read(C.PC)
		if C.issetN() {
			dest := uint16(int(C.PC) + int(int8(C.OperLO)))
			if C.PC&0xFF00 != dest&0xFF00 {
				C.pageCrossed = true
			} else {
				C.pageCrossed = false
				C.PC = dest
			}
		} else {
			C.Inst = C.Mnemonic[C.instCode]
			C.PC++
			C.CycleCount = 1
		}
	case 4:
		C.instCode = C.ram.Read(C.PC)
		if C.pageCrossed {
			C.PC = uint16(int(C.PC) + int(int8(C.OperLO)))
		} else {
			C.Inst = C.Mnemonic[C.instCode]
			C.PC++
			C.CycleCount = 1
		}
	case 5:
		C.instCode = C.ram.Read(C.PC)
		C.Inst = C.Mnemonic[C.instCode]
		C.PC++
		C.CycleCount = 1
	}
}

///////////////////////////////////////////////////////
//                        BPL                        //
///////////////////////////////////////////////////////

func (C *CPU) BPL_rel() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.instCode = C.ram.Read(C.PC)
		if !C.issetN() {
			dest := uint16(int(C.PC) + int(int8(C.OperLO)))
			if C.PC&0xFF00 != dest&0xFF00 {
				C.pageCrossed = true
			} else {
				C.pageCrossed = false
				C.PC = dest
			}
		} else {
			C.Inst = C.Mnemonic[C.instCode]
			C.PC++
			C.CycleCount = 1
		}
	case 4:
		C.instCode = C.ram.Read(C.PC)
		if C.pageCrossed {
			C.PC = uint16(int(C.PC) + int(int8(C.OperLO)))
		} else {
			C.Inst = C.Mnemonic[C.instCode]
			C.PC++
			C.CycleCount = 1
		}
	case 5:
		C.instCode = C.ram.Read(C.PC)
		C.Inst = C.Mnemonic[C.instCode]
		C.PC++
		C.CycleCount = 1
	}

}
