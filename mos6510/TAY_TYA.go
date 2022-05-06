package mos6510

///////////////////////////////////////////////////////
//                        TAY                        //
///////////////////////////////////////////////////////

func (C *CPU) TAY_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.Y = C.A

		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        TYA                        //
///////////////////////////////////////////////////////

func (C *CPU) TYA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.A = C.Y

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}
