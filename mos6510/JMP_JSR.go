package mos6510

import "fmt"

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
		C.IndAddrHI = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO+1))
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
		C.stack[C.SP] = byte(C.PC >> 8)
		C.SP--
	case 5:
		C.stack[C.SP] = byte(C.PC)
		C.SP--
	case 6:
		C.OperHI = C.ram.Read(C.PC)
		C.PC = (uint16(C.OperHI) << 8) + uint16(C.OperLO)
		C.CycleCount = 0
		C.StackDebugPt++
		C.StackDebug[C.StackDebugPt] = fmt.Sprintf("%04X: JSR %02X%02X -> %02X:%02X %02X:%02X\n", C.InstStart, C.OperHI, C.OperLO, C.SP+1, C.stack[C.SP+1], C.SP+2, C.stack[C.SP+2])
	}
}
