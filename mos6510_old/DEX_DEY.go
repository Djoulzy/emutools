package mos6510

func (C *CPU) DEX_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.X--

		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) DEY_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.Y--

		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}
