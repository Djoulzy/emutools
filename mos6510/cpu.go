package mos6510

import (
	"fmt"
	"log"
	"time"

	"github.com/Djoulzy/Tools/confload"

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
	fmt.Printf("mos6510 - PC: %04X\n", C.PC)

	perfStats = make(map[byte][]time.Duration)
	for index := range C.Mnemonic {
		perfStats[index] = make([]time.Duration, 0)
	}
}

func (C *CPU) Init(MEM *mem.BANK, conf *confload.ConfigData) {
	fmt.Printf("mos6510 - Init\n")
	C.conf = conf
	C.ram = MEM
	C.stack = MEM.Layouts[0].Layers[0][StackStart : StackStart+256]
	C.ramSize = len(MEM.Layouts[0].Layers[0])
	fmt.Printf("%d\n", C.ramSize)
	C.initLanguage()
	C.Reset()
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
	if C.conf.RunPerfStats {
		defer C.timeTrack(time.Now(), "ComputeInstruction")
	}
	C.FullInst = Disassemble(C.Inst, C.Oper)
	if C.cycleCount != C.Inst.Cycles {
		log.Printf("%s - Wanted: %d - Getting: %d\n", C.FullInst, C.Inst.Cycles, C.cycleCount)
	}
	if C.cycleCount == C.Inst.Cycles {
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

	C.cycleCount++
	// fmt.Printf("%d - %d\n", C.cycleCount, C.State)
	switch C.State {
	case Idle:
		C.cycleCount = 0
		C.State++

	////////////////////////////////////////////////
	// Cycle 1
	////////////////////////////////////////////////
	case ReadInstruction:
		C.cycleCount = 1
		C.InstStart = C.PC
		C.instCode = C.ram.Read(C.PC)
		C.instDump = fmt.Sprintf("%02X", C.instCode)

		if C.Inst, ok = C.Mnemonic[C.instCode]; !ok {
			log.Printf(fmt.Sprintf("Unknown instruction: %02X at %04X\n", C.instCode, C.PC))
			// C.State = Idle
		}
		if C.Inst.addr == implied {
			C.State = Compute
			C.PC += 1
			C.CheckInterrupts()
		} else {
			C.State = ReadOperLO
			if C.Inst.Cycles <= 3 {
				C.CheckInterrupts()
			}
		}

	////////////////////////////////////////////////
	// Cycle 2
	////////////////////////////////////////////////
	case ReadOperLO:
		C.Oper = uint16(C.ram.Read(C.PC + 1))
		C.instDump += fmt.Sprintf(" %02X", C.Oper)

		switch C.Inst.addr {
		case relative:
			fallthrough
		case immediate:
			C.State = Compute
			C.PC += 2
			if C.Inst.Cycles == 2 {
				C.ComputeInstruction()
			}
		case absolute:
			fallthrough
		case indirect:
			fallthrough
		case absoluteX:
			fallthrough
		case absoluteY:
			C.State = ReadOperHI
		case zeropage:
			fallthrough
		case zeropageX:
			fallthrough
		case zeropageY:
			fallthrough
		case indirectX:
			fallthrough
		case indirectY:
			C.State = ReadZP
		default:
			log.Fatal("Erreur de cycle")
		}
		if C.Inst.Cycles == 4 {
			C.CheckInterrupts()
		}

	////////////////////////////////////////////////
	// Cycle 3
	////////////////////////////////////////////////
	case ReadZP:
		C.PC += 2
		switch C.Inst.addr {
		case zeropage:
			C.State = Compute
			if C.Inst.Cycles == 3 {
				C.ComputeInstruction()
			}
		case zeropageX:
			fallthrough
		case zeropageY:
			C.State = ReadZP_XY
		case indirectX:
			fallthrough
		case indirectY:
			C.State = ReadIndXY_LO
		default:
			log.Fatal("Erreur de cycle")
		}
		if C.Inst.Cycles == 5 {
			C.CheckInterrupts()
		}

	case ReadOperHI: // Cycle 3
		tmp := C.ram.Read(C.PC + 2)
		C.Oper += uint16(tmp) << 8
		C.instDump += fmt.Sprintf(" %02X", tmp)

		C.PC += 3
		switch C.Inst.addr {
		case absolute:
			// C.ram.Write(C.Oper, C.ram.Read(C.Oper)) // Pour Bruce Lee mais pourquoi ?
			C.State = Compute
			if C.Inst.Cycles == 3 {
				C.ComputeInstruction()
			}
		case absoluteX:
			fallthrough
		case absoluteY:
			C.State = ReadAbsXY
		case indirect:
			C.State = ReadIndirect
		default:
			log.Fatal("Erreur de cycle")
		}
		if C.Inst.Cycles == 5 {
			C.CheckInterrupts()
		}

	////////////////////////////////////////////////
	// Cycle 4
	////////////////////////////////////////////////
	case ReadZP_XY: // Cycle 4
		switch C.Inst.addr {
		case zeropageX:
			fallthrough
		case zeropageY:
			C.State = Compute
			if C.Inst.Cycles == 4 {
				C.ComputeInstruction()
			}
		default:
			log.Fatal("Erreur de cycle")
		}

	case ReadIndXY_LO: // Cycle 4
		switch C.Inst.addr {
		case indirectX:
			C.State = ReadIndXY_HI
		case indirectY:
			C.State = ReadIndXY_HI
		default:
			log.Fatal("Erreur de cycle")
		}
		if C.Inst.Cycles == 6 {
			C.CheckInterrupts()
		}

	case ReadIndirect: // Cycle 4
		C.State = Compute

	case ReadAbsXY: // Cycle 4
		switch C.Inst.addr {
		case absoluteX:
			fallthrough
		case absoluteY:
			C.State = Compute
			if C.Inst.Cycles == 4 {
				C.ComputeInstruction()
			}
		default:
			log.Fatal("Erreur de cycle")
		}

	////////////////////////////////////////////////
	// Cycle 5
	////////////////////////////////////////////////

	case ReadIndXY_HI:
		switch C.Inst.addr {
		case indirectX:
			C.State = Compute
		case indirectY:
			C.State = Compute
			if C.Inst.Cycles == 5 {
				C.ComputeInstruction()
			}
		default:
			log.Fatal("Erreur de cycle")
		}
		if C.Inst.Cycles > 6 {
			C.CheckInterrupts()
		}

	////////////////////////////////////////////////
	// Cycle 6-7-8
	////////////////////////////////////////////////

	case Compute:
		if C.Inst.Cycles == C.cycleCount {
			C.ComputeInstruction()
		}

	////////////////////////////////////////////////
	// Interrupt
	////////////////////////////////////////////////

	case NMI1:
		C.NMI_Raised = false
		C.INT_delay = false
		C.State = NMI2
	case NMI2:
		C.pushWordStack(C.PC)
		C.State = NMI3
	case NMI3:
		C.State = NMI4
	case NMI4:
		C.pushByteStack(C.S)
		C.State = NMI5
	case NMI5:
		C.State = NMI6
	case NMI6:
		C.State = NMI7
	case NMI7:
		C.PC = C.readWord(0xFFFA)
		C.State = ReadInstruction

	case IRQ1:
		C.IRQ_Raised = false
		C.INT_delay = false
		C.State = IRQ2
	case IRQ2:
		C.pushWordStack(C.PC)
		C.State = IRQ3
	case IRQ3:
		C.State = IRQ4
	case IRQ4:
		C.pushByteStack(C.S)
		C.State = IRQ5
	case IRQ5:
		C.State = IRQ6
	case IRQ6:
		C.setI(true)
		C.State = IRQ7
	case IRQ7:
		C.PC = C.readWord(0xFFFE)
		C.State = ReadInstruction

	default:
		log.Fatal("Unknown CPU state\n")
	}
}
