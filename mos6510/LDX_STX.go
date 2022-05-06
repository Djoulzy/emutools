package mos6510

///////////////////////////////////////////////////////
//                        LDX                        //
///////////////////////////////////////////////////////

func (C *CPU) LDX_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.X = C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) LDX_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.X = C.ram.Read(uint16(C.OperLO))
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) LDX_zpy() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO = C.OperLO + C.Y
	case 4:
		C.X = C.ram.Read(uint16(C.OperLO))
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) LDX_abs() {
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
		C.X = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) LDX_aby() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.Y
		C.PC++
	case 4:
		C.X = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.X)
		C.updateZ(C.X)
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.CycleCount = 0
		}
	case 5:
		C.X = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.X)
		C.updateZ(C.X)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        STX                        //
///////////////////////////////////////////////////////

func (C *CPU) STX_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Write(uint16(C.OperLO), C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) STX_zpy() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO += C.Y
	case 4:
		C.ram.Write(uint16(C.OperLO), C.X)
		C.CycleCount = 0
	}
}

func (C *CPU) STX_abs() {
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
		C.ram.Write((uint16(C.OperHI)<<8)+uint16(C.OperLO), C.X)
		C.CycleCount = 0
	}
}
