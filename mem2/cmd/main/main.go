package main

import (
	"fmt"

	"github.com/Djoulzy/emutools/mem2"
)

const (
	ramSize = 80
	romSize = 40

	DISABLED = true
	ENABLED  = false
)

var (
	MEM     mem2.BANK
	BankSel byte
	RAM     []byte
)

func main() {
	BankSel = 0
	MEM = mem2.InitBanks(1, &BankSel)

	ZONE1 := make([]byte, ramSize)
	mem2.Clear(ZONE1, ramSize, 0xAA)
	ZONE2 := make([]byte, romSize)
	mem2.Clear(ZONE2, romSize, 0xBB)
	ZONE3 := make([]byte, romSize)
	mem2.Clear(ZONE3, romSize, 0xCC)

	MEM.Layouts[0] = mem2.InitConfig(ramSize)
	MEM.Layouts[0].Attach("ZONE1", 0x0000, ZONE1, mem2.READWRITE, ENABLED)
	MEM.Layouts[0].Attach("ZONE2", 20, ZONE2, mem2.READWRITE, ENABLED)
	MEM.Layouts[0].Attach("ZONE3", romSize, ZONE3, mem2.READWRITE, ENABLED)

	// RAM = ZONE1[:]
	// copy(RAM[20:], ZONE2)
	// copy(RAM[romSize:], ZONE3)

	// RAM = append(RAM, ZONE2...)

	// ZONE1[1] = 0xFF
	// ZONE2[1] = 0xDD
	// ZONE3[1] = 0xEE

	// cpt := 0
	// for _, val := range RAM {
	// 	fmt.Printf("%02X ", val)
	// 	cpt++
	// 	if cpt == 10 {
	// 		fmt.Printf("\n")
	// 		cpt = 0
	// 	}
	// }

	var TEST []*byte
	TEST = make([]*byte, ramSize)
	
	for i,val := range ZONE1 {
		TEST[i] = &val
	}

	ZONE1[1] = 0xFF
	ZONE2[1] = 0xDD
	ZONE3[1] = 0xEE

	cpt := 0
	for _, val := range TEST {
		fmt.Printf("%02X ", *val)
		cpt++
		if cpt == 10 {
			fmt.Printf("\n")
			cpt = 0
		}
	}
}
