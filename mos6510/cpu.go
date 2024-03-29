package mos6510

import (
	"fmt"
	"log"
	"time"

	"github.com/Djoulzy/mmu"
)

// var perfStats map[byte][]time.Duration
var start time.Time
var is_debug bool

// func (C *CPU) timeTrack(start time.Time, name string) {
// 	elapsed := time.Now().Sub(start)
// 	perfStats[C.instCode] = append(perfStats[C.instCode], elapsed)
// }

func (C *CPU) Reset() {
	C.A = 0xAA
	C.X = 0
	C.Y = 0
	C.S = 0b00110000
	C.SP = 0xFF

	C.IRQ_pin = 0
	C.NMI_pin = 0
	C.NMI_Raised = false
	C.IRQ_Raised = false
	C.INT_delay = false

	// Cold Start:
	C.setI(true)
	C.PC = C.readWord(COLDSTART_Vector)
	// C.PC = 0xFA62

	// perfStats = make(map[byte][]time.Duration)
	// for index := range C.Mnemonic {
	// 	perfStats[index] = make([]time.Duration, 0)
	// }
}

func (C *CPU) Init(Model string, MEM *mmu.MMU, debug bool) {
	C.model = Model
	fmt.Printf("%s - Init\n", Model)
	C.ram = MEM
	C.initLanguage()
	C.Reset()
	C.CycleCount = 0
	is_debug = debug

	// start = time.Now()
}

//////////////////////////////////
////// Addressage Indirect ///////
//////////////////////////////////

func (C *CPU) ReadIndirectX(addr uint16) byte {
	dest := addr + uint16(C.X)
	return C.ram.Read((uint16(C.ram.Read(dest+1)) << 8) + uint16(C.ram.Read(dest)))
}

func (C *CPU) ReadIndirectY(addr uint16) byte {
	dest := (uint16(C.ram.Read(addr+1)) << 8) + uint16(C.ram.Read(addr))
	return C.ram.Read(dest + uint16(C.Y))
}

func (C *CPU) GetIndirectYAddr(addr uint16, pagecrossed *bool) uint16 {
	base := (uint16(C.ram.Read(addr+1)) << 8) + uint16(C.ram.Read(addr))
	dest := base + uint16(C.Y)
	*pagecrossed = (base&0xFF00 == dest&0xFF00)
	return dest
}

func (C *CPU) WriteIndirectX(addr uint16, val byte) {
	dest := addr + uint16(C.X)
	C.ram.Write((uint16(C.ram.Read(dest+1))<<8)+uint16(C.ram.Read(dest)), val)
}

func (C *CPU) WriteIndirectY(addr uint16, val byte) {
	dest := (uint16(C.ram.Read(addr+1)) << 8) + uint16(C.ram.Read(addr))
	C.ram.Write(dest+uint16(C.Y), val)
}

//////////////////////////////////
/////// Addressage Relatif ///////
//////////////////////////////////

func (C *CPU) getRelativeAddr(dist uint16) uint16 {
	return uint16(int(C.PC) + int(int8(dist)))
}

//////////////////////////////////
//////////// Read Word ///////////
//////////////////////////////////

func (C *CPU) readWord(addr uint16) uint16 {
	return (uint16(C.ram.Read(addr+1)) << 8) + uint16(C.ram.Read(addr))
}

//////////////////////////////////
////////// Stack Access //////////
//////////////////////////////////

func (C *CPU) readStack() byte {
	C.SP++
	return C.ram.Read(StackStart + uint16(C.SP))
}

func (C *CPU) writeStack(value byte) {
	C.ram.Write(StackStart+uint16(C.SP), value)
	C.SP--
}

//////////////////////////////////
/////////// Interrupts ///////////
//////////////////////////////////

func (C *CPU) CheckInterrupts() {
	if C.NMI_pin > 0 {
		C.NMI_Raised = true
	}
	if (C.IRQ_pin > 0) && (C.S & ^I_mask) == 0 {
		C.IRQ_Raised = true
	}
}

//////////////////////////////////
///////////// Running ////////////
//////////////////////////////////

func (C *CPU) GoTo(addr uint16) {
	C.PC = addr
}

// func (C *CPU) ComputeInstruction() {
// 	if C.CycleCount == C.Inst.Cycles {
// 		if C.NMI_Raised || C.IRQ_Raised {
// 			if C.Inst.Cycles <= 2 && !C.INT_delay {
// 				C.INT_delay = true
// 				C.State = ReadInstruction
// 			} else {
// 				if C.IRQ_Raised {
// 					C.State = IRQ1
// 				}
// 				if C.NMI_Raised {
// 					C.State = NMI1
// 				}
// 			}
// 		} else {
// 			C.State = ReadInstruction
// 		}
// 	}
// 	C.Inst.action()
// }

func (C *CPU) firstCycle() {
	var ok bool
	C.InstStart = C.PC
	// if C.NMI_Raised || C.IRQ_Raised {
	// 	if C.IRQ_Raised {
	// 		C.instCode = 0x6F
	// 	}
	// 	if C.NMI_Raised {
	// 		C.instCode = 0x7F
	// 	}
	// } else {
	C.instCode = C.ram.Read(C.PC)
	// }
	if C.Inst, ok = C.Mnemonic[C.instCode]; !ok {
		log.Printf(fmt.Sprintf("Unknown instruction: %02X at %04X\n", C.instCode, C.PC))
	}
	C.Inst.action()
}

func (C *CPU) NextCycle() {
	C.CycleCount++
	switch C.CycleCount {
	case 1:
		C.firstCycle()
	default:
		C.Inst.action()
	}
	// C.CheckInterrupts()
	// if C.CycleCount == 0 {
	// 	if C.NMI_Raised || C.IRQ_Raised {
	// 		if C.Inst.Cycles <= 2 && !C.INT_delay {
	// 			C.INT_delay = true
	// 		} else {
	// 			if C.IRQ_Raised {
	// 				C.instCode = 0x6F
	// 			}
	// 			if C.NMI_Raised {
	// 				C.instCode = 0x7F
	// 			}
	// 			C.Inst = C.Mnemonic[C.instCode]
	// 			C.Inst.action()
	// 			C.CycleCount = 1
	// 		}
	// 	}
	// }
}