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

func (C *CPU) codeHexa(inst Instruction, addr uint16) string {
	var i uint16
	res := fmt.Sprintf("%02X", C.ram.DirectRead(addr-1))
	for i = 0; i < uint16(inst.bytes-1); i++ {
		res = fmt.Sprintf("%s %02X", res, C.ram.DirectRead(addr+i))
	}
	return res
}

func (C *CPU) codeLine(inst Instruction, addr uint16) string {
	switch {
	case inst.addr == 0:
		return inst.Name
	case (inst.addr > indirect):
		return fmt.Sprintf("%s "+InstTemplate[inst.addr], inst.Name, C.ram.DirectRead(addr+1), C.ram.DirectRead(addr))
	default:
		return fmt.Sprintf("%s "+InstTemplate[inst.addr], inst.Name, C.ram.DirectRead(addr))
	}
}

func (C *CPU) Disassemble(addr uint16, nblines int) {
	var count int = 0
	var inst Instruction
	var ok bool
	var code byte

	for count < nblines {
		count++
		code = C.ram.DirectRead(addr)
		if inst, ok = C.Mnemonic[code]; !ok {
			fmt.Printf("%04X: ???\n", addr)
			addr++
		} else {
			fmt.Printf("%04X: %-10s %s\n", addr, C.codeHexa(inst, addr+1), C.codeLine(inst, addr+1))
			addr += uint16(inst.bytes)
		}
	}
}

func (C *CPU) Trace() string {
	// return fmt.Sprintf("%d  %s   A:%c[1;33m%02X%c[0m X:%c[1;33m%02X%c[0m Y:%c[1;33m%02X%c[0m SP:%c[1;33m%02X%c[0m  %c[1;30m(%d)%c[0m %c[1;37m%-10s%c[0m",
	// 	C.GlobalCycles, C.registers(), 27, C.A, 27, 27, C.X, 27, 27, C.Y, 27, 27, C.SP, 27, 27, C.Inst.Cycles, 27, 27, C.FullInst, 27)

	template := "%s   A:%02X X:%02X Y:%02X SP:%02X  (%d) %04X: %-10s %s"
	return fmt.Sprintf(template, C.registers(), C.A, C.X, C.Y, C.SP, C.Inst.Cycles, C.InstStart, C.codeHexa(C.Inst, C.PC), C.codeLine(C.Inst, C.PC))
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
