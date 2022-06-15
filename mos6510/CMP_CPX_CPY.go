package mos6510

var tmp int
var tmpByte byte
var addr uint16

///////////////////////////////////////////////////////
//                        CMP                        //
///////////////////////////////////////////////////////

func (C *CPU) CMP_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
		tmp = int(C.A) - int(C.OperLO)

		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		addr = uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_zpx() {
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
		addr = uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_abs() {
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
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_abx() {
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
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.setC(C.A >= C.ram.Read(addr))
			C.updateN(tmpByte)
			C.updateZ(tmpByte)
			C.CycleCount = 0
		}
	case 5:
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_aby() {
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
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.setC(C.A >= C.ram.Read(addr))
			C.updateN(tmpByte)
			C.updateZ(tmpByte)
			C.CycleCount = 0
		}
	case 5:
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_inx() {
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
		addr = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		tmpByte = C.A - C.ram.Read(addr)

		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_iny() {
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
		addr = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		tmpByte = C.A - C.ram.Read(addr)
		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.setC(C.A >= C.ram.Read(addr))
			C.updateN(tmpByte)
			C.updateZ(tmpByte)
			C.CycleCount = 0
		}
	case 6:
		addr = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		tmpByte = C.A - C.ram.Read(addr)
		C.setC(C.A >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        CPX                        //
///////////////////////////////////////////////////////

func (C *CPU) CPX_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
		tmp = int(C.X) - int(C.OperLO)

		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CPX_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		addr = uint16(C.OperLO)
		tmpByte = C.X - C.ram.Read(addr)
		C.setC(C.X >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CPX_abs() {
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
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.X - C.ram.Read(addr)
		C.setC(C.X >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        CPY                        //
///////////////////////////////////////////////////////

func (C *CPU) CPY_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
		tmp = int(C.Y) - int(C.OperLO)

		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CPY_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		addr = uint16(C.OperLO)
		tmpByte = C.Y - C.ram.Read(addr)
		C.setC(C.Y >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}

func (C *CPU) CPY_abs() {
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
		addr = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		tmpByte = C.Y - C.ram.Read(addr)
		C.setC(C.Y >= C.ram.Read(addr))
		C.updateN(tmpByte)
		C.updateZ(tmpByte)
		C.CycleCount = 0
	}
}
