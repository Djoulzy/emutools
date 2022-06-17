package mos6510

///////////////////////////////////////////////////////
//                        NOP                        //
///////////////////////////////////////////////////////

func (C *CPU) NOP_1x1() {
	switch C.CycleCount {
	case 1:
		C.PC++
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_1x2() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_2x2() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC)
		C.PC++
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_2x3() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_2x4() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC)
		C.PC++
	case 3:
	case 4:
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_3x4() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(C.PC)
		C.PC++
	case 4:
		C.CycleCount = 0
	}
}

func (C *CPU) NOP_3x8() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(C.PC)
		C.PC++
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
		C.CycleCount = 0
	}
}
