package mos6510

import (
	"fmt"
)

func (C *CPU) registers() string {
	var i, mask byte
	res := ""
	for i = 0; i < 8; i++ {
		mask = 1 << i
		if C.S&mask == mask {
			res = regString[i] + res
		} else {
			res = "-" + res
		}
	}
	return res
}

func (C *CPU) disassemble() string {
	// switch C.Inst.addr {
	// case implied:
	// 	token = fmt.Sprintf("")
	// case immediate:
	// 	token = fmt.Sprintf("#$%02X", C.ram.Read(C.PC+1))
	// case relative:
	// 	token = fmt.Sprintf("$%02X", C.ram.Read(C.PC+1))
	// case zeropage:
	// 	token = fmt.Sprintf("$%02X", C.ram.Read(C.PC+1))
	// case zeropageX:
	// 	token = fmt.Sprintf("$%02X,X", C.ram.Read(C.PC+1))
	// case zeropageY:
	// 	token = fmt.Sprintf("$%02X,Y", C.ram.Read(C.PC+1))
	// case Branching:
	// 	fallthrough
	// case CrossPage:
	// 	fallthrough
	// case absolute:
	// 	token = fmt.Sprintf("$%02X%02X", C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	// case absoluteX:
	// 	token = fmt.Sprintf("$%02X%02X,X", C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	// case absoluteY:
	// 	token = fmt.Sprintf("$%02X%02X,Y", C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	// case indirect:
	// 	token = fmt.Sprintf("($%02X%02X)", C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	// case indirectX:
	// 	token = fmt.Sprintf("($%02X,X)", C.ram.Read(C.PC+1))
	// case indirectY:
	// 	token = fmt.Sprintf("($%02X),Y", C.ram.Read(C.PC+1))
	// }
	template := "%04X: " + C.Inst.Name + InstTemplate[C.Inst.addr]
	if C.Inst.addr < indirect {
		return fmt.Sprintf(template, C.InstStart, C.ram.Read(C.PC+1))
	} else {
		return fmt.Sprintf(template, C.InstStart, C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	}
}

func (C *CPU) Trace() string {
	// return fmt.Sprintf("%d  %s   A:%c[1;33m%02X%c[0m X:%c[1;33m%02X%c[0m Y:%c[1;33m%02X%c[0m SP:%c[1;33m%02X%c[0m  %c[1;30m(%d)%c[0m %c[1;37m%-10s%c[0m",
	// 	C.GlobalCycles, C.registers(), 27, C.A, 27, 27, C.X, 27, 27, C.Y, 27, 27, C.SP, 27, 27, C.Inst.Cycles, 27, 27, C.FullInst, 27)

	template := "%s   A:%02X X:%02X Y:%02X SP:%02X  (%d) %04X: " + C.Inst.Name + " " + InstTemplate[C.Inst.addr]
	if C.Inst.addr == 0 {
		return fmt.Sprintf(template, C.registers(), C.A, C.X, C.Y, C.SP, C.Inst.Cycles, C.InstStart)
	}
	if C.Inst.addr < indirect {
		return fmt.Sprintf(template, C.registers(), C.A, C.X, C.Y, C.SP, C.Inst.Cycles, C.InstStart, C.ram.Read(C.PC+1))
	} else {
		return fmt.Sprintf(template, C.registers(), C.A, C.X, C.Y, C.SP, C.Inst.Cycles, C.InstStart, C.ram.Read(C.PC+2), C.ram.Read(C.PC+1))
	}
}

// func ColVal(val time.Duration) string {
// 	if val > time.Microsecond {
// 		return clog.CSprintf("white", "red", "%10s", val)
// 	} else {
// 		return fmt.Sprintf("%10s", val)
// 	}
// }

// func (C *CPU) DumpStats() {
// 	var min time.Duration
// 	var max time.Duration

// 	for index, val := range perfStats {
// 		total := 0
// 		cpt := 0
// 		hicount := 0
// 		min = time.Minute
// 		max = 0
// 		for _, duree := range val {
// 			cpt++
// 			total += int(duree)
// 			if duree > time.Microsecond {
// 				hicount++
// 			}
// 			if duree > max {
// 				max = duree
// 			}
// 			if duree < min {
// 				min = duree
// 			}
// 		}
// 		if cpt > 0 {
// 			moy := time.Duration(total / cpt)
// 			hiPercent := float32(hicount) / float32(cpt) * 100
// 			fmt.Printf("$%02X: (%s) Moy: %s - Max: %s - Min: %s - NbHi: %5d = %6.2f%% - Nb Samples: %d \n", index, C.Mnemonic[index].Name, ColVal(moy), ColVal(max), ColVal(min), hicount, hiPercent, cpt)
// 		}
// 	}
// }

// func (C *CPU) DumpStackDebug() {
// 	for i := 0; i <= C.StackDebugPt; i++ {
// 		fmt.Printf("%s", C.StackDebug[i])
// 	}
// }
