package mos6510

import "fmt"

///////////////////////////////////////////////////////
//                        RTS                        //
///////////////////////////////////////////////////////

func (C *CPU) RTS_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.IndAddrLO = C.stack[C.SP]
		C.SP++
	case 5:
		C.IndAddrHI = C.stack[C.SP]
	case 6:
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.PC++
		C.CycleCount = 0
		C.StackDebugPt--
	}
}

///////////////////////////////////////////////////////
//                        RTI                        //
///////////////////////////////////////////////////////

func (C *CPU) RTI_imp() {
	fmt.Printf("RTI\n")
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.S = C.stack[C.SP]
		C.SP++
	case 5:
		C.IndAddrLO = C.stack[C.SP]
		C.SP++
	case 6:
		C.IndAddrHI = C.stack[C.SP]
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.setB(false)
		C.setU(false)
		C.CycleCount = 0
	}
}
