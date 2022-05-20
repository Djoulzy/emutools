package mos6510

///////////////////////////////////////////////////////
//                        PHA                        //
///////////////////////////////////////////////////////

func (C *CPU) PHA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.stack[C.SP] = C.A
		C.SP--
		C.CycleCount = 0
		// C.StackDebugPt++
		// C.StackDebug[C.StackDebugPt] = fmt.Sprintf("%02X=%02X - %04X: PHA #$%02X\n", C.SP+1, C.stack[C.SP+1], C.InstStart, C.A)
	}
}

///////////////////////////////////////////////////////
//                        PLA                        //
///////////////////////////////////////////////////////

func (C *CPU) PLA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.A = C.stack[C.SP]
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
		// C.StackDebugPt--
	}
}
