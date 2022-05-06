package mos6510

///////////////////////////////////////////////////////
//                        TXA                        //
///////////////////////////////////////////////////////

func (C *CPU) TXA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.A = C.X

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        TAX                        //
///////////////////////////////////////////////////////

func (C *CPU) TAX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.X = C.A

		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}
