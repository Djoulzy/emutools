package mos6510

import "fmt"

///////////////////////////////////////////////////////
//                        BRK                        //
///////////////////////////////////////////////////////

func (C *CPU) BRK_imp() {
	// fmt.Printf("BRK\n")
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
		C.PC++
	case 3:
		C.writeStack(byte(C.PC >> 8))
	case 4:
		C.writeStack(byte(C.PC))
	case 5:
		C.writeStack(C.S | ^B_mask | ^U_mask)
		if C.model != "6502" {
			C.setD(false)
		}
		C.setI(true)
	case 6:
		C.IndAddrLO = C.ram.Read(IRQBRK_Vector)
		C.PC = uint16(C.IndAddrLO)
	case 7:
		C.IndAddrHI = C.ram.Read(IRQBRK_Vector + 1)
		C.PC += uint16(C.IndAddrHI) << 8
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        IRQ                        //
///////////////////////////////////////////////////////

func (C *CPU) IRQ_imp() {
	switch C.CycleCount {
	case 1:
		C.IRQ_Raised = false
		C.INT_delay = false
	case 2:
	case 3:
		C.writeStack(byte(C.PC >> 8))
	case 4:
		C.writeStack(byte(C.PC))
	case 5:
		C.writeStack(C.S)
		C.setI(true)
	case 6:
		C.IndAddrLO = C.ram.Read(IRQBRK_Vector)
		C.PC = uint16(C.IndAddrLO)
	case 7:
		C.IndAddrHI = C.ram.Read(IRQBRK_Vector + 1)
		C.PC += uint16(C.IndAddrHI) << 8
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        NMI                        //
///////////////////////////////////////////////////////

func (C *CPU) NMI_imp() {
	fmt.Printf("NMI\n")
	switch C.CycleCount {
	case 1:
		C.NMI_Raised = false
		C.INT_delay = false
	case 2:
	case 3:
		C.writeStack(byte(C.PC >> 8))
	case 4:
		C.writeStack(byte(C.PC))
	case 5:
		C.writeStack(C.S)
	case 6:
		C.IndAddrLO = C.ram.Read(IRQBRK_Vector)
		C.PC = uint16(C.IndAddrLO)
	case 7:
		C.IndAddrHI = C.ram.Read(IRQBRK_Vector + 1)
		C.PC += uint16(C.IndAddrHI) << 8
		C.CycleCount = 0
	}
}
