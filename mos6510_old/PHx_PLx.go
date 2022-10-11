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
//                        PHX                        //
///////////////////////////////////////////////////////

func (C *CPU) PHX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.stack[C.SP] = C.X
		C.SP--
		C.CycleCount = 0
		// C.StackDebugPt++
		// C.StackDebug[C.StackDebugPt] = fmt.Sprintf("%02X=%02X - %04X: PHA #$%02X\n", C.SP+1, C.stack[C.SP+1], C.InstStart, C.A)
	}
}

///////////////////////////////////////////////////////
//                        PHY                        //
///////////////////////////////////////////////////////

func (C *CPU) PHY_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.stack[C.SP] = C.Y
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

///////////////////////////////////////////////////////
//                        PLX                        //
///////////////////////////////////////////////////////

func (C *CPU) PLX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.X = C.stack[C.SP]
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
		// C.StackDebugPt--
	}
}

///////////////////////////////////////////////////////
//                        PLY                        //
///////////////////////////////////////////////////////

func (C *CPU) PLY_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.Y = C.stack[C.SP]
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
		// C.StackDebugPt--
	}
}
