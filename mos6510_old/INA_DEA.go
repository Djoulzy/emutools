package mos6510

///////////////////////////////////////////////////////
//                        INA                        //
///////////////////////////////////////////////////////

func (C *CPU) INA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.A++

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        DEA                        //
///////////////////////////////////////////////////////

func (C *CPU) DEA_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.A--

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}
