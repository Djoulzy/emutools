package mos6510_2

import (
	"log"
)

///////////////////////////////////////////////////////
//                        CMP                        //
///////////////////////////////////////////////////////

func (C *CPU) CMP_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
		tmp := int(C.A) - int(C.OperLO)

		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		tmp := int(C.A) - int(C.ram.Read(uint16(C.OperLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_zpx() {
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
		tmp := int(C.A) - int(C.ram.Read(uint16(C.OperLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_abs() {
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
		tmp := int(C.A) - int(C.ram.Read((uint16(C.OperHI)<<8)+uint16(C.OperLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_abx() {
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
		tmp := int(C.A) - int(C.ram.Read((uint16(C.OperHI)<<8)+uint16(C.OperLO)))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.setC(tmp >= 0)
			C.updateN(byte(tmp))
			C.updateZ(byte(tmp))
			C.CycleCount = 0
		}
	case 5:
		tmp := int(C.A) - int(C.ram.Read((uint16(C.OperHI)<<8)+uint16(C.OperLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_aby() {
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
		tmp := int(C.A) - int(C.ram.Read((uint16(C.OperHI)<<8)+uint16(C.OperLO)))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.setC(tmp >= 0)
			C.updateN(byte(tmp))
			C.updateZ(byte(tmp))
			C.CycleCount = 0
		}
	case 5:
		tmp := int(C.A) - int(C.ram.Read((uint16(C.OperHI)<<8)+uint16(C.OperLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_inx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.Pointer = C.OperLO + C.X
	case 4:
		C.IndAddrLO = C.ram.Read(uint16(C.Pointer))
	case 5:
		C.IndAddrHI = C.ram.Read(uint16(C.Pointer + 1))
	case 6:
		tmp := int(C.A) - int(C.ram.Read((uint16(C.IndAddrHI)<<8)+uint16(C.IndAddrLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

func (C *CPU) CMP_iny() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.IndAddrLO = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.IndAddrHI = C.ram.Read(uint16(C.OperLO + 1))
		if (uint16(C.IndAddrLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := int(C.A) - int(C.ram.Read((uint16(C.IndAddrHI)<<8)+uint16(C.IndAddrLO)))
		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.setC(tmp >= 0)
			C.updateN(byte(tmp))
			C.updateZ(byte(tmp))
			C.CycleCount = 0
		}
	case 6:
		tmp := int(C.A) - int(C.ram.Read((uint16(C.IndAddrHI)<<8)+uint16(C.IndAddrLO)))
		C.setC(tmp >= 0)
		C.updateN(byte(tmp))
		C.updateZ(byte(tmp))
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        CPX                        //
///////////////////////////////////////////////////////

func (C *CPU) cpx() {
	var val int

	switch C.Inst.addr {
	case immediate:
		val = int(C.X) - int(C.Oper)
	case zeropage:
		val = int(C.X) - int(C.ram.Read(C.Oper))
	case absolute:
		val = int(C.X) - int(C.ram.Read(C.Oper))
	default:
		log.Fatal("Bad addressing mode")
	}
	C.setC(val >= 0)
	C.updateN(byte(val))
	C.updateZ(byte(val))

}

///////////////////////////////////////////////////////
//                        CPY                        //
///////////////////////////////////////////////////////

func (C *CPU) cpy() {
	var val int

	switch C.Inst.addr {
	case immediate:
		val = int(C.Y) - int(C.Oper)
	case zeropage:
		val = int(C.Y) - int(C.ram.Read(C.Oper))
	case absolute:
		val = int(C.Y) - int(C.ram.Read(C.Oper))
	default:
		log.Fatal("Bad addressing mode")
	}
	C.setC(val >= 0)
	C.updateN(byte(val))
	C.updateZ(byte(val))

}
