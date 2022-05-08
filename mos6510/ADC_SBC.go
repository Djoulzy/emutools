package mos6510

import (
	"github.com/albenik/bcd"
)

func (C *CPU) doADC(inputVal byte) byte {
	if C.getD() == 1 {
		val_bcd := uint16(bcd.ToUint8(C.A)) + uint16(bcd.ToUint8(inputVal)) + uint16(C.getC())
		vals := bcd.FromUint16(val_bcd)
		C.setC(vals[0] > 0)
		C.setV(val_bcd > 100)
		return vals[1]
	} else {
		val := uint16(C.A) + uint16(inputVal) + uint16(C.getC())
		C.setC(val > 0x00FF)
		C.updateV(C.A, inputVal, byte(val))
		return byte(val)
	}
}

func (C *CPU) doSBC(inputVal byte) byte {
	if C.getD() == 1 {
		val_bcd := uint16(bcd.ToUint8(C.A)) - uint16(bcd.ToUint8(byte(inputVal)))
		if C.getC() == 0 {
			val_bcd -= 1
		}
		vals := bcd.FromUint16(val_bcd)
		C.setC(vals[1] > 0)
		C.setV(val_bcd > 100)
		return vals[1]
	} else {
		val := int(C.A) - int(inputVal)
		if C.getC() == 0 {
			val -= 1
		}
		C.setC(val >= 0x00)
		C.updateV(C.A, ^inputVal, byte(val))
		return byte(val)
	}
}

///////////////////////////////////////////////////////
//                        ADC                        //
///////////////////////////////////////////////////////

func (C *CPU) ADC_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		tmp := C.ram.Read(C.PC)
		C.PC++
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_zpx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO = C.OperLO + C.X
	case 4:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_abs() {
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
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_abx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.X)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.X
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_aby() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.Y
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_inx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.Pointer = C.OperLO + C.X
	case 4:
		C.IndAddrLO = C.ram.Read(uint16(C.Pointer))
	case 5:
		C.IndAddrHI = C.ram.Read(uint16(C.Pointer + 1))
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_iny() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.IndAddrLO = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.IndAddrHI = C.ram.Read(uint16(C.OperLO + 1))
		if (uint16(C.IndAddrLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))

		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        SBC                        //
///////////////////////////////////////////////////////

func (C *CPU) SBC_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		tmp := C.ram.Read(C.PC)
		C.PC++
		C.A = C.doSBC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doSBC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_zpx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO = C.OperLO + C.X
	case 4:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doSBC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_abs() {
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
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doSBC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_abx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.X)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.X
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doSBC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doSBC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_aby() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.Y
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doSBC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doSBC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_inx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.Pointer = C.OperLO + C.X
	case 4:
		C.IndAddrLO = C.ram.Read(uint16(C.Pointer))
	case 5:
		C.IndAddrHI = C.ram.Read(uint16(C.Pointer + 1))
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doSBC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) SBC_iny() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.IndAddrLO = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.IndAddrHI = C.ram.Read(uint16(C.OperLO + 1))
		if (uint16(C.IndAddrLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))

		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A = C.doSBC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doSBC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}
