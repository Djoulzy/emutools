package mos6510

///////////////////////////////////////////////////////
//                        JMP                        //
///////////////////////////////////////////////////////

func (C *CPU) JMP_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		C.CycleCount = 0
	}
}

func (C *CPU) JMP_ind() {
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
		C.IndAddrLO = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		if C.OperLO == 0xFF {
			if C.model == "6502" {
				C.IndAddrHI = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO+1))
			} else {
				C.IndAddrHI = C.ram.Read((uint16(C.OperHI+1) << 8))
			}
		} else {
			C.IndAddrHI = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO+1))
		}
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.CycleCount = 0
	}
}

func (C *CPU) JMP_inx() {
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
		if (uint16(C.OperLO) + uint16(C.X)) > 0x00FF {
			C.OperHI++
		}
		C.OperLO += C.X
	case 5:
		C.IndAddrLO = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		if C.OperLO == 0xFF {
			C.IndAddrHI = C.ram.Read((uint16(C.OperHI+1) << 8))
		} else {
			C.IndAddrHI = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO+1))
		}
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        JSR                        //
///////////////////////////////////////////////////////

func (C *CPU) JSR_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
	case 4:
		C.writeStack(byte(C.PC >> 8))
	case 5:
		C.writeStack(byte(C.PC))
	case 6:
		C.OperHI = C.ram.Read(C.PC)
		C.PC = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		C.CycleCount = 0
	}
}
