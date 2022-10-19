package mem

import (
	"fmt"
	"io/ioutil"

	"github.com/Djoulzy/Tools/clog"
)

func LoadData(zone []byte, file string, memStart uint16) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	for i, val := range content {
		// mem.Write(memStart+uint16(i), val)
		zone[memStart+uint16(i)] = val
	}
	return nil
}

func Fill(zone []byte, val byte) {
	for i := range zone {
		zone[i] = val
	}
}

func dispBin(bin byte) []byte {
	var mask byte = 0b10000000
	var res string = ""
	for i := 0; i < 8; i++ {
		if (bin & mask) > 0 {
			res += clog.CSprintf("black", "white", " ")
		} else {
			res += clog.CSprintf("black", "dark_gray", " ")
		}
		mask >>= 1
	}
	return []byte(res)
}

func DisplayCharRom(zone []byte, bytePerLine int, nbLines int, nbDispPerLine int) {
	var charStartAddr int
	var charLine []byte

	size := len(zone)
	charSize := bytePerLine * nbLines
	nbChar := size / charSize

	clog.CPrintf("light_gray", "black", "- Nb Char found: %d\n", nbChar)
	for y := 0; y < nbChar; y += nbDispPerLine {
		for l := 0; l < nbLines; l++ {
			for x := 0; x < nbDispPerLine; x++ {
				charStartAddr = (y * charSize) + (x * nbLines) + l
				charLine = dispBin(zone[charStartAddr])
				if l == 0 {
					clog.CPrintf("light_gray", "black", "%03X", y+x)
				} else {
					fmt.Printf("   ")
				}
				clog.CPrintf("light_gray", "black", "%s ", charLine)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
