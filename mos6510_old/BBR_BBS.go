package mos6510

///////////////////////////////////////////////////////
//                       BBR0                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR0_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000001 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR1                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR1_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000010 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR2                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR2_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000100 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR3                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR3_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00001000 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR4                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR4_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00010000 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR5                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR5_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00100000 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR6                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR6_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b01000000 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBR7                        //
///////////////////////////////////////////////////////

func (C *CPU) BBR7_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b10000000 == 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS0                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS0_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000001 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS1                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS1_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000010 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS2                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS2_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00000100 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS3                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS3_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00001000 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        BBS4                       //
///////////////////////////////////////////////////////

func (C *CPU) BBS4_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00010000 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS5                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS5_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b00100000 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS6                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS6_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b01000000 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                       BBS7                        //
///////////////////////////////////////////////////////

func (C *CPU) BBS7_rel() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		if C.tmpBuff&0b10000000 != 0 {
			C.PC = uint16(int(C.PC) + int(int8(C.OperHI)))
		}
		C.CycleCount = 0
	}
}
