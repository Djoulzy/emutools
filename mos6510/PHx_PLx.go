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
		C.writeStack(C.A)
		C.CycleCount = 0
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
		C.writeStack(C.X)
		C.CycleCount = 0
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
		C.writeStack(C.Y)
		C.CycleCount = 0
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
		C.tmpBuff = C.readStack()
	case 4:
		C.A = C.tmpBuff
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
		C.tmpBuff = C.readStack()
	case 4:
		C.X = C.tmpBuff
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
		C.tmpBuff = C.readStack()
	case 4:
		C.Y = C.tmpBuff
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
		// C.StackDebugPt--
	}
}
