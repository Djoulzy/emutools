package mos6510

///////////////////////////////////////////////////////
//                        CLI                        //
///////////////////////////////////////////////////////

func (C *CPU) CLI_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setI(false)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        SEI                        //
///////////////////////////////////////////////////////

func (C *CPU) SEI_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setI(true)
		C.CycleCount = 0
	}
}
