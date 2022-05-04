package mos6510_2

import (
	"fmt"
	"log"
	"time"

	"github.com/Djoulzy/emutools/mem"
)

var perfStats map[byte][]time.Duration

func (C *CPU) timeTrack(start time.Time, name string) {
	elapsed := time.Now().Sub(start)
	perfStats[C.instCode] = append(perfStats[C.instCode], elapsed)
}

func (C *CPU) Reset() {
	C.A = 0xAA
	C.X = 0
	C.Y = 0
	C.S = 0b00100000
	C.SP = 0xFF

	C.IRQ_pin = 0
	C.NMI_pin = 0
	C.NMI_Raised = false
	C.IRQ_Raised = false
	C.INT_delay = false

	// PLA Settings (Bank switching)
	// C.ram.Write(0x0000, 0x2F)
	// C.ram.Write(0x0001, 0x1F)

	C.State = ReadInstruction
	// Cold Start:
	C.PC = C.readWord(COLDSTART_Vector)
	fmt.Printf("mos6510_2 - PC: %04X\n", C.PC)

	perfStats = make(map[byte][]time.Duration)
	for index := range C.Mnemonic {
		perfStats[index] = make([]time.Duration, 0)
	}
}

func (C *CPU) Init(MEM *mem.BANK) {
	fmt.Printf("mos6510_2 - Init\n")
	C.ram = MEM
	C.stack = MEM.Layouts[0].Layers[0][StackStart : StackStart+256]
	C.ramSize = len(MEM.Layouts[0].Layers[0])
	C.initLanguage()
	C.Reset()
	C.CycleCount = 0
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
//////// Stack Operations ////////
//////////////////////////////////

// Byte
func (C *CPU) pushByteStack(val byte) {
	// if C.SP < 90 {
	// 	os.Exit(1)
	// }
	C.stack[C.SP] = val
	C.SP--
}

func (C *CPU) pullByteStack() byte {
	C.SP++
	// if C.SP == 0x00 {
	// 	C.ram.DumpStack(C.SP)
	// 	log.Fatal("Stack overflow")
	// }
	return C.stack[C.SP]
}

// Word
func (C *CPU) pushWordStack(val uint16) {
	low := byte(val)
	hi := byte(val >> 8)
	C.pushByteStack(hi)
	C.pushByteStack(low)
}

func (C *CPU) pullWordStack() uint16 {
	low := C.pullByteStack()
	hi := uint16(C.pullByteStack()) << 8
	return hi + uint16(low)
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

func (C *CPU) ComputeInstruction() {
	C.FullInst = fmt.Sprintf("%04X: %s", C.InstStart, Disassemble(C.Inst, C.Oper))
	if C.CycleCount != C.Inst.Cycles {
		log.Printf("%s - Wanted: %d - Getting: %d\n", C.FullInst, C.Inst.Cycles, C.CycleCount)
	}
	if C.CycleCount == C.Inst.Cycles {
		if C.NMI_Raised || C.IRQ_Raised {
			if C.Inst.Cycles <= 2 && !C.INT_delay {
				C.INT_delay = true
				C.State = ReadInstruction
			} else {
				if C.IRQ_Raised {
					C.State = IRQ1
				}
				if C.NMI_Raised {
					C.State = NMI1
				}
			}
		} else {
			C.State = ReadInstruction
		}
	}
	C.Inst.action()
}

func (C *CPU) NextCycle() {
	var ok bool

	C.CycleCount++
	switch C.CycleCount {
	case 1:
		C.instCode = C.ram.Read(C.PC)
		if C.Inst, ok = C.Mnemonic[C.instCode]; !ok {
			log.Printf(fmt.Sprintf("Unknown instruction: %02X at %04X\n", C.instCode, C.PC))
		}
		fallthrough
	default:
		C.Inst.action()
	}
}
