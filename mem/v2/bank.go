package mem

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

func Init(nbMemLayout int, layoutsSize uint, bankSelector *byte) *BANK  {
	tmp := BANK{}
	tmp.Layouts = make([]CONFIG, nbMemLayout)
	for i := 0; i < nbMemLayout; i++ {
		tmp.Layouts[i].InitConfig(layoutsSize)
	}
	tmp.Selector = bankSelector
	return &tmp
}

func (B *BANK) Attach(layoutNum int, name string, start uint16, content []byte, mode bool, disabled bool, accessor interface{}) {
	if accessor == nil {
		B.Layouts[layoutNum].Attach(name, start, content, mode, disabled, nil)
	} else {
		B.Layouts[layoutNum].Attach(name, start, content, mode, disabled, accessor.(MEMAccess))
	}
}

func (B *BANK) Accessor(layoutNum int, layerName string, access MEMAccess) {
	B.Layouts[layoutNum].Accessors[B.Layouts[layoutNum].LayersName[layerName]] = access
}

func (B *BANK) GetFullSize() int {
	return len(B.Layouts[0].StorageRef[0])
}

func (B *BANK) GetStack(start uint16, length uint16) []MEMCell {
	return B.Layouts[0].VisibleMem[start : start+length]
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
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic - Read Memory: %04X\n", addr)
			layout.VisibleMem[addr].dump()
			os.Exit(1)
		}
	}()

	layout = &B.Layouts[*B.Selector]
	return layout.VisibleMem[addr].Accessor.MRead(layout.VisibleMem, addr)
}

func (B *BANK) Write(addr uint16, value byte) {
	layout = &B.Layouts[*B.Selector]

	layerNum := layout.VisibleMem[addr].LayerNum
	if layout.ReadOnlyMode[layerNum] == false {
		layout.VisibleMem[addr].Accessor.MWrite(layout.VisibleMem, addr, value)
	} else {
		layout.VisibleMem[addr].Accessor.MWriteUnder(layout.VisibleMem, addr, value)
	}
}

func (B *BANK) Clear(zone []byte, interval int, startWith byte) {
	// interval: 0x40 pour C64
	//           0x1000 pour Apple
	// startWith: 0x00 pour C64
	//            0xFF pour Apple
	cpt := 0
	fill := byte(startWith)
	for i := range zone {
		zone[i] = fill
		cpt++
		if cpt == interval {
			fill = ^fill
			cpt = 0
		}
	}
}

func (B *BANK) LoadROM(size int, file string) []byte {
	val := make([]byte, size)
	if len(file) > 0 {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if len(data) != size {
			panic("Bad ROM Size")
		}
		for i := 0; i < size; i++ {
			val[i] = byte(data[i])
		}
	}
	return val
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

// func (B *BANK) Show() {
// 	B.Layouts[*B.Selector].Show()
// }

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

// func (B *BANK) CheckLayoutForAddr(addr uint16) {
// 	layout = &B.Layouts[*B.Selector]
// 	page = int(addr >> PAGE_DIVIDER)

// 	for cpt = 0; layout.Disabled[layout.LayerByPages[page][cpt]]; cpt++ {
// 	}
// 	layerNum = layout.LayerByPages[page][cpt]

// 	for index, layer := range layout.LayerByPages[page] {
// 		if index == cpt {
// 			fmt.Printf("[X] ")
// 		} else {
// 			fmt.Printf("[-] ")
// 		}
// 		fmt.Printf("Pos: %d - Name: %s - Disabled: %v\n", index, layout.NameLayers[layer], layout.Disabled[layer])
// 	}
// }
