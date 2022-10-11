package mos6510

///////////////////////////////////////////////////////
//                        INX                        //
///////////////////////////////////////////////////////

func (C *CPU) INX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.X++

		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        INY                        //
///////////////////////////////////////////////////////

func (C *CPU) INY_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.Y++

		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}
