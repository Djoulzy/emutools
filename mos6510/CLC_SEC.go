package mos6510

///////////////////////////////////////////////////////
//                        CLC                        //
///////////////////////////////////////////////////////

func (C *CPU) CLC_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setC(false)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        SEC                        //
///////////////////////////////////////////////////////

func (C *CPU) SEC_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setC(true)
		C.CycleCount = 0
	}
}
