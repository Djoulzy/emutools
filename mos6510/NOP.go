package mos6510

///////////////////////////////////////////////////////
//                        NOP                        //
///////////////////////////////////////////////////////

func (C *CPU) NOP_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.CycleCount = 0
	}
}
