package mos6510

///////////////////////////////////////////////////////
//                        PHP                        //
///////////////////////////////////////////////////////

func (C *CPU) PHP_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.writeStack(C.S | ^B_mask | ^U_mask)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        PLP                        //
///////////////////////////////////////////////////////

func (C *CPU) PLP_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.tmpBuff = C.readStack()
	case 4:
		C.S = C.tmpBuff & B_mask & U_mask
		C.CycleCount = 0
		// C.StackDebugPt--
	}
}
