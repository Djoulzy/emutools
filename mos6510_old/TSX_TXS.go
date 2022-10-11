package mos6510

///////////////////////////////////////////////////////
//                        TSX                        //
///////////////////////////////////////////////////////

func (C *CPU) TSX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.X = C.SP

		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        TXS                        //
///////////////////////////////////////////////////////

func (C *CPU) TXS_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.SP = C.X
		C.CycleCount = 0
	}
}
