package mos6510

///////////////////////////////////////////////////////
//                        RMB                        //
///////////////////////////////////////////////////////

func (C *CPU) RMB0_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11111110
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB1_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11111101
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB2_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11111011
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB3_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11110111
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB4_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11101111
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB5_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b11011111
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB6_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b10111111
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) RMB7_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff &= 0b01111111
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        SMB                        //
///////////////////////////////////////////////////////

func (C *CPU) SMB0_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00000001
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB1_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00000010
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB2_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00000100
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB3_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00001000
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB4_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00010000
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB5_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b00100000
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB6_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b01000000
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) SMB7_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.tmpBuff |= 0b10000000
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.CycleCount = 0
	}
}
