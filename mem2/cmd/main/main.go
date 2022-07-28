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

func showLayout() {
	cpt := 0
	for _, val := range MEM.Layouts[0].VisibleMem {
		fmt.Printf("%02X ", *val.Val)
		cpt++
		if cpt == 10 {
			fmt.Printf("\n")
			cpt = 0
		}
	}
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
	showLayout()

	MEM.Write(2, 0xFF)
	MEM.Write(22, 0xDD)
	MEM.Write(42, 0xEE)

	MEM.Disable("ZONE2")
	MEM.Disable("ZONE2")
	MEM.Enable("ZONE2")
	MEM.Enable("ZONE2")
	MEM.Enable("ZONE3")
	MEM.Enable("ZONE1")

	showLayout()

	MEM.Write(22, 0x00)

	fmt.Printf("%02X - %02X - %02X\n", ZONE1[2], ZONE2[2], ZONE3[2])
	fmt.Printf("%02X - %02X - %02X\n", ZONE1[2], ZONE1[22], ZONE1[42])
}
