package mos6510

///////////////////////////////////////////////////////
//                        LDY                        //
///////////////////////////////////////////////////////

func (C *CPU) LDY_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.Y = C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) LDY_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.Y = C.ram.Read(uint16(C.OperLO))
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) LDY_zpx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO = C.OperLO + C.X
	case 4:
		C.Y = C.ram.Read(uint16(C.OperLO))
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) LDY_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC++
	case 4:
		C.Y = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) LDY_abx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.X)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.X
		C.PC++
	case 4:
		C.Y = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.Y)
		C.updateZ(C.Y)
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.CycleCount = 0
		}
	case 5:
		C.Y = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.Y)
		C.updateZ(C.Y)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        STY                        //
///////////////////////////////////////////////////////

func (C *CPU) STY_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Write(uint16(C.OperLO), C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) STY_zpx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO += C.X
	case 4:
		C.ram.Write(uint16(C.OperLO), C.Y)
		C.CycleCount = 0
	}
}

func (C *CPU) STY_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC++
	case 4:
		C.ram.Write((uint16(C.OperHI)<<8)+uint16(C.OperLO), C.Y)
		C.CycleCount = 0
	}
}
