package mos6510

///////////////////////////////////////////////////////
//                        ROL                        //
///////////////////////////////////////////////////////

func (C *CPU) ROL_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		carry := C.A&0x80 == 0x80
		C.A <<= 1
		if C.issetC() {
			C.A++
		}
		C.setC(carry)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ROL_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x80 == 0x80
		C.tmpBuff <<= 1
		if C.issetC() {
			C.tmpBuff++
		}
		C.setC(carry)
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROL_zpx() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x80 == 0x80
		C.tmpBuff <<= 1
		if C.issetC() {
			C.tmpBuff++
		}
		C.setC(carry)
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROL_abs() {
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x80 == 0x80
		C.tmpBuff <<= 1
		if C.issetC() {
			C.tmpBuff++
		}
		C.setC(carry)
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROL_abx() {
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
			C.Inst.Cycles++
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x80 == 0x80
		C.tmpBuff <<= 1
		if C.issetC() {
			C.tmpBuff++
		}
		C.setC(carry)
	case 7:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        ROR                        //
///////////////////////////////////////////////////////

func (C *CPU) ROR_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		carry := C.A&0x01 > 0
		C.A >>= 1
		if C.issetC() {
			C.A |= 0x80
		}
		C.setC(carry)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ROR_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x01 > 0
		C.tmpBuff >>= 1
		if C.issetC() {
			C.tmpBuff |= 0x80
		}
		C.setC(carry)
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROR_zpx() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x01 > 0
		C.tmpBuff >>= 1
		if C.issetC() {
			C.tmpBuff |= 0x80
		}
		C.setC(carry)
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROR_abs() {
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x01 > 0
		C.tmpBuff >>= 1
		if C.issetC() {
			C.tmpBuff |= 0x80
		}
		C.setC(carry)
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) ROR_abx() {
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
			C.Inst.Cycles++
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		carry := C.tmpBuff&0x01 > 0
		C.tmpBuff >>= 1
		if C.issetC() {
			C.tmpBuff |= 0x80
		}
		C.setC(carry)
	case 7:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}
