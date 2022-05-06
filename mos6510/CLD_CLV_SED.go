package mos6510

///////////////////////////////////////////////////////
//                        CLD                        //
///////////////////////////////////////////////////////

func (C *CPU) CLD_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setD(false)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        SED                        //
///////////////////////////////////////////////////////

func (C *CPU) SED_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setD(true)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        CLV                        //
///////////////////////////////////////////////////////

func (C *CPU) CLV_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setV(false)
		C.CycleCount = 0
	}
}
