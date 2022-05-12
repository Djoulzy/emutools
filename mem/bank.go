package mem

import (
	"fmt"

	"github.com/Djoulzy/Tools/clog"
	"github.com/Djoulzy/emutools/charset"
)

const StackStart = 0x0100

var bank byte

type BANK struct {
	Selector *byte
	Layouts  []CONFIG
}

func InitBanks(nbMemLayout int, sel *byte) BANK {
	B := BANK{}
	B.Layouts = make([]CONFIG, nbMemLayout)
	B.Selector = sel
	return B
}

func (B *BANK) Read(addr uint16) byte {
	layerNum := B.Layouts[*B.Selector].LayerByPages[int(addr>>PAGE_DIVIDER)]
	return B.Layouts[*B.Selector].Accessors[layerNum].MRead(B.Layouts[*B.Selector].Layers[layerNum], addr-B.Layouts[*B.Selector].Start[layerNum])
}

func (B *BANK) Write(addr uint16, value byte) {
	layerNum := B.Layouts[*B.Selector].LayerByPages[int(addr>>PAGE_DIVIDER)]
	if B.Layouts[*B.Selector].ReadOnly[layerNum] {
		layerNum = 0
	}
	// if layerNum == 0 {
	// 	clog.Test("MEM", "Write", "Addr: %04X, Page: %d, Data: %02X", addr, int(addr>>PAGE_DIVIDER), value)
	// }
	B.Layouts[*B.Selector].Accessors[layerNum].MWrite(B.Layouts[*B.Selector].Layers[layerNum], addr-B.Layouts[*B.Selector].Start[layerNum], value)
}

func (B *BANK) Dump(startAddr uint16) {
	var val byte
	var line string
	var ascii string

	cpt := startAddr
	for j := 0; j < 16; j++ {
		fmt.Printf("%04X : ", cpt)
		line = ""
		ascii = ""
		for i := 0; i < 16; i++ {
			val = B.Read(cpt)
			if val != 0x00 && val != 0xFF {
				line = line + clog.CSprintf("white", "black", "%02X", val) + " "
			} else {
				line = fmt.Sprintf("%s%02X ", line, val)
			}
			if _, ok := charset.PETSCII[val]; ok {
				ascii += fmt.Sprintf("%s", string(charset.PETSCII[val]))
			} else {
				ascii += "."
			}
			if i == 7 {
				line = fmt.Sprintf("%s  ", line)
			}
			cpt++
		}
		fmt.Printf("%s - %s\n", line, ascii)
	}
}

func (B *BANK) Show() {
	B.Layouts[*B.Selector].Show()
}

func (B *BANK) DumpStack(sp byte) {
	cpt := uint16(0x0100)
	fmt.Printf("\n")
	for j := 0; j < 16; j++ {
		fmt.Printf("%04X : ", cpt)
		for i := 0; i < 16; i++ {
			if cpt == StackStart+uint16(sp) {
				clog.CPrintf("white", "red", "%02X", B.Read(cpt))
				fmt.Print(" ")
				// fmt.Printf("%c[41m%c[0m[0;31m%02X%c[0m ", 27, 27, P.Mem[RAM].Val[cpt], 27)
			} else {
				fmt.Printf("%02X ", B.Read(cpt))
			}
			cpt++
		}
		fmt.Println()
	}
}
