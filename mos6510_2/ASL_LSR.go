package mos6510_2

///////////////////////////////////////////////////////
//                        ASL                        //
///////////////////////////////////////////////////////

func (C *CPU) ASL_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setC(C.A&0b10000000 > 1)
		C.A <<= 1

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ASL_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff <<= 1
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ASL_zpx() {
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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff <<= 1
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ASL_abs() {
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff <<= 1
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ASL_abx() {
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
		C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		}
	case 5:
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff <<= 1
	case 7:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        LSR                        //
///////////////////////////////////////////////////////

func (C *CPU) LSR_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.setC(C.A&0x01 == 0x01)
		C.A >>= 1

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LSR_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0x01 == 0x01)
		C.RMWBuff >>= 1
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) LSR_zpx() {
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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0x01 == 0x01)
		C.RMWBuff >>= 1
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) LSR_abs() {
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0x01 == 0x01)
		C.RMWBuff >>= 1
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) LSR_abx() {
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
		C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		}
	case 5:
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0x01 == 0x01)
		C.RMWBuff >>= 1
	case 7:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}
