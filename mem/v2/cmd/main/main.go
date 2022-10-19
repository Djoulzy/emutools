package main

import (
	"fmt"

	mem "github.com/Djoulzy/emutools/mem/v2"
)

const (
	ramSize = 80
	romSize = 40

	DISABLED = true
	ENABLED  = false
)

var (
	MEM     *mem.BANK
	BankSel byte
	RAM     []byte
)

type TestAccessor struct {
}

func (DA *TestAccessor) MRead(mem_access []mem.MEMCell, addr uint16) byte {
	return *mem_access[addr].Val
}

func (DA *TestAccessor) MWrite(mem_access []mem.MEMCell, addr uint16, value byte) {
	*mem_access[addr].Val = value
}

func (DA *TestAccessor) MWriteUnder(mem_access []mem.MEMCell, addr uint16, value byte) {
	fmt.Printf("Write Under\n")
	*mem_access[addr].Under = value
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
	MEM = mem.Init(1, ramSize, &BankSel)

	ZONE1 := make([]byte, ramSize)
	MEM.Clear(ZONE1, ramSize, 0xAA)
	ZONE2 := make([]byte, romSize)
	MEM.Clear(ZONE2, romSize, 0xBB)
	ZONE3 := make([]byte, romSize)
	MEM.Clear(ZONE3, romSize, 0xCC)

	MEM.Attach(0, "ZONE1", 0x0000, ZONE1, mem.READWRITE, ENABLED, nil)
	MEM.Attach(0, "ZONE2", 20, ZONE2, mem.READONLY, ENABLED, &TestAccessor{})
	MEM.Attach(0, "ZONE3", 40, ZONE3, mem.READWRITE, ENABLED, nil)
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
