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

type TestAccessor struct {

}

func (DA *TestAccessor) MRead(mem []mem2.MEMCell, addr uint16) byte {
	return *mem[addr].Val
}

func (DA *TestAccessor) MWrite(mem []mem2.MEMCell, addr uint16, value byte) {
	*mem[addr].Val = value
}

func (DA *TestAccessor) MWriteUnder(mem []mem2.MEMCell, addr uint16, value byte) {
	fmt.Printf("Write Under\n")
	*mem[addr].Under = value
}

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
	MEM.Layouts[0].Attach("ZONE1", 0x0000, ZONE1, mem2.READWRITE, ENABLED, nil)
	MEM.Layouts[0].Attach("ZONE2", 20, ZONE2, mem2.READONLY, ENABLED, &TestAccessor{})
	MEM.Layouts[0].Attach("ZONE3", 40, ZONE3, mem2.READWRITE, ENABLED, nil)

	// ZONE1[1] = 0xFF
	// ZONE2[1] = 0xDD
	// ZONE3[1] = 0xEE

	MEM.Write(2, 0xFF)
	MEM.Write(22, 0xDD)
	MEM.Write(42, 0xEE)

	cpt := 0
	for _, val := range MEM.Layouts[0].VisibleMem {
		fmt.Printf("%02X ", *val.Val)
		cpt++
		if cpt == 10 {
			fmt.Printf("\n")
			cpt = 0
		}
	}

	fmt.Printf("%02X - %02X - %02X\n", ZONE1[2], ZONE2[2], ZONE3[2])
	fmt.Printf("%02X - %02X - %02X\n", ZONE1[2], ZONE1[22], ZONE1[42])

	// var TEST []*byte
	// TEST = make([]*byte, ramSize)
	
	// for i := range ZONE1 {
	// 	TEST[i] = &ZONE1[i]
	// }

	// ZONE1[1] = 0xFF
	// ZONE2[1] = 0xDD
	// ZONE3[1] = 0xEE

	// tmp := TEST[3]
	// *tmp = 0xBB

	// cpt := 0
	// for _, val := range TEST {
	// 	fmt.Printf("%02X ", *val)
	// 	cpt++
	// 	if cpt == 10 {
	// 		fmt.Printf("\n")
	// 		cpt = 0
	// 	}
	// }
	// fmt.Printf("\n")
	// cpt = 0
	// for _, val := range ZONE1 {
	// 	fmt.Printf("%02X ", val)
	// 	cpt++
	// 	if cpt == 10 {
	// 		fmt.Printf("\n")
	// 		cpt = 0
	// 	}
	// }

	// var ZONE []byte = []byte{0x11, 0x22, 0x33}
	// var TEST []*byte

	// TEST = make([]*byte, 3)

	// TEST[0] = &ZONE[0]

	// fmt.Printf("%v - %02X\n", TEST[0], *TEST[0])

	// ZONE[0] = 0x22

	// fmt.Printf("%v - %02X\n", TEST[0], *TEST[0])

	// *TEST[0] = 0x33

	// fmt.Printf("%v - %02X\n", &ZONE[0], ZONE[0])
}
