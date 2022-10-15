package mem

import (
	"fmt"

	"github.com/Djoulzy/Tools/clog"
	"github.com/Djoulzy/emutools/charset"
)

const StackStart = 0x0100

var layout *CONFIG
var cpt, page, layerNum int

type BANK struct {
	Selector *byte
	Layouts  []CONFIG
}

func Init(nbMemLayout int, layoutsSize uint, bankSelector *byte) *BANK {
	B := BANK{}
	B.Layouts = make([]CONFIG, nbMemLayout)
	for i := 0; i < nbMemLayout; i++ {
		B.Layouts[i].InitConfig(layoutsSize)
	}
	B.Selector = bankSelector
	return &B
}

func (B *BANK) Attach(layoutNum int, name string, start uint16, content []byte, mode bool, disabled bool, accessor MEMAccess) {
	B.Layouts[layoutNum].Attach(name, start, content, mode, disabled, accessor)
}

func (B *BANK) GetFullSize() int {
	return len(B.Layouts[0].Layers[0])
}

func (B *BANK) GetStack(start uint16, length uint16) []byte {
	return B.Layouts[0].Layers[0][start : start+length]
}

func (B *BANK) Disable(layerName string) {
	for _, config := range B.Layouts {
		config.Disable(layerName)
	}
}

func (B *BANK) Enable(layerName string) {
	for _, config := range B.Layouts {
		config.Enable(layerName)
	}
}

func (B *BANK) ReadOnly(layerName string) {
	for _, config := range B.Layouts {
		config.ReadOnly(layerName)
	}
}

func (B *BANK) ReadWrite(layerName string) {
	for _, config := range B.Layouts {
		config.ReadWrite(layerName)
	}
}

func (B *BANK) Read(addr uint16) byte {
	layout = &B.Layouts[*B.Selector]
	page = int(addr >> PAGE_DIVIDER)

	for cpt = 0; layout.Disabled[layout.LayerByPages[page][cpt]]; cpt++ {
	}
	layerNum = layout.LayerByPages[page][cpt]
	return layout.Accessors[layerNum].MRead(layout.Layers[layerNum], addr-layout.Start[layerNum])
}

func (B *BANK) Write(addr uint16, value byte) {
	layout = &B.Layouts[*B.Selector]
	page = int(addr >> PAGE_DIVIDER)

	for cpt = 0; layout.Disabled[layout.LayerByPages[page][cpt]] || layout.ReadOnlyMode[layout.LayerByPages[page][cpt]]; cpt++ {
	}
	layerNum = layout.LayerByPages[page][cpt]
	layout.Accessors[layerNum].MWrite(layout.Layers[layerNum], addr-layout.Start[layerNum], value)
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

func (B *BANK) CheckLayoutForAddr(addr uint16) {
	layout = &B.Layouts[*B.Selector]
	page = int(addr >> PAGE_DIVIDER)

	for cpt = 0; layout.Disabled[layout.LayerByPages[page][cpt]]; cpt++ {
	}
	layerNum = layout.LayerByPages[page][cpt]

	for index, layer := range layout.LayerByPages[page] {
		if index == cpt {
			fmt.Printf("[X] ")
		} else {
			fmt.Printf("[-] ")
		}
		fmt.Printf("Pos: %d - Name: %s - Disabled: %v\n", index, layout.NameLayers[layer], layout.Disabled[layer])
	}
}
