package mos6510

///////////////////////////////////////////////////////
//                        RTS                        //
///////////////////////////////////////////////////////

func (C *CPU) RTS_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.tmpBuff = C.readStack()
	case 4:
		C.IndAddrLO = C.tmpBuff
		C.tmpBuff = C.readStack()
	case 5:
		C.IndAddrHI = C.tmpBuff
	case 6:
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.PC++
		C.CycleCount = 0
		// C.StackDebugPt -= 2
	}
}

///////////////////////////////////////////////////////
//                        RTI                        //
///////////////////////////////////////////////////////

func (C *CPU) RTI_imp() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.tmpBuff = C.readStack()
	case 4:
		C.S = C.tmpBuff
		C.tmpBuff = C.readStack()
	case 5:
		C.IndAddrLO = C.tmpBuff
		C.tmpBuff = C.readStack()
	case 6:
		C.IndAddrHI = C.tmpBuff
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.setB(false)
		C.setU(false)
		C.CycleCount = 0
	}
}
