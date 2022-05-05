package mos6510_2

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Djoulzy/emutools/mem"
)

const (
	ramSize    = 65536
	kernalSize = 8192
	ioSize     = 4096
)

type TestData struct {
	code       string
	acc        byte
	x          byte
	y          byte
	flag       byte
	destMem    uint16
	res        uint16
	resFlag    byte
	cycles     int
	cyclesDone int
}

type TestSuite struct {
	inst byte
	data []TestData
}

func getAddrName(a addressing) string {
	switch a {
	case implied:
		return "implied"
	case immediate:
		return "immediate"
	case relative:
		return "relative"
	case zeropage:
		return "zeropage"
	case zeropageX:
		return "zeropageX"
	case zeropageY:
		return "zeropageY"
	case absolute:
		return "absolute"
	case absoluteX:
		return "absoluteX"
	case absoluteY:
		return "absoluteY"
	case indirect:
		return "indirect"
	case indirectX:
		return "indirectX"
	case indirectY:
		return "indirectY"
	case Branching:
		return "Branching"
	case CrossPage:
		return "CrossPage"
	}
	return "Unknown"
}

func (TS *TestSuite) Add(td TestData) {
	TS.data = append(TS.data, td)
}

func loadMem(data string) {
	tmp := strings.Split(data, " ")
	addr, _ := strconv.ParseInt(tmp[0], 16, 16)
	proc.PC = uint16(addr)
	for i := 0; i < len(tmp)-1; i++ {
		addr, _ = strconv.ParseInt(tmp[i+1], 16, 16)
		// log.Printf("Str: %s = Hex: %02X\n", tmp[i+1], addr)
		RAM[proc.PC+uint16(i)] = byte(addr)
	}
}

func (TD *TestData) run() {
	// log.Printf("%v\n", TD)
	proc.S = TD.flag
	proc.A = TD.acc
	proc.X = TD.x
	proc.Y = TD.y

	loadMem(TD.code)

	proc.CycleCount = 0
	TD.cyclesDone = 0
	for {
		proc.NextCycle()
		TD.cyclesDone++
		// log.Printf("PC: %04X - val: %02X - cycle %d\n", proc.PC, RAM[proc.PC], proc.CycleCount)
		if proc.CycleCount == 0 {
			break
		}
	}
}

func (TD *TestData) checkBit(t *testing.T, val1, val2 byte, name string) bool {
	if val1 != val2 {
		t.Errorf("[%s] %s %s - Incorrect %s - get: %08b - want: %08b", TD.code, proc.Inst.Name, getAddrName(proc.Inst.addr), name, val1, val2)
		return false
	}
	return true
}

func (TD *TestData) checkByte(t *testing.T, val1, val2 byte, name string) bool {
	if val1 != val2 {
		t.Errorf("%s %s - Incorrect %s - get: %02X - want: %02X", proc.Inst.Name, getAddrName(proc.Inst.addr), name, val1, val2)
		return false
	}
	return true
}

func (TD *TestData) checkWord(t *testing.T, val1, val2 uint16, name string) bool {
	if val1 != val2 {
		t.Errorf("%s %s - Incorrect %s - get: %04X - want: %04X", proc.Inst.Name, getAddrName(proc.Inst.addr), name, val1, val2)
		return false
	}
	return true
}

func (TD *TestData) checkCycles(t *testing.T, name string) bool {
	if uint16(TD.cyclesDone) != uint16(TD.cycles) {
		t.Errorf("[%s] %s %s - Incorrect %s - get: %d - want: %d", TD.code, proc.Inst.Name, getAddrName(proc.Inst.addr), name, TD.cyclesDone, TD.cycles)
		return false
	}
	return true
}

func finalize(name string, allGood bool) {
	if allGood {
		log.Printf("%s OK", name)
	} else {
		log.Printf("%s %c[1;31mECHEC%c[0m", name, 27, 27)
	}
}

var proc CPU
var BankSel byte
var MEM mem.BANK
var RAM []byte
var SystemClock uint16

func TestMain(m *testing.M) {
	SystemClock = 0

	RAM = make([]byte, ramSize)

	BankSel = 0
	MEM = mem.InitBanks(1, &BankSel)

	MEM.Layouts[0] = mem.InitConfig(ramSize)
	MEM.Layouts[0].Attach("RAM", 0, RAM, mem.READWRITE)

	proc.Init(&MEM)
	os.Exit(m.Run())
}

func TestStack(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)
	for i := 0; i <= 0xFF; i++ {
		proc.pushByteStack(byte(i))
	}
	for i := 0xFF; i >= 0; i-- {
		if proc.pullByteStack() != byte(i) {
			t.Errorf("Bad stack operation")
			allGood = false
		}
	}

	for i := 0; i <= 0x7F; i++ {
		proc.pushWordStack(uint16(i))
	}
	for i := 0x7F; i >= 0; i-- {
		if proc.pullWordStack() != uint16(i) {
			t.Errorf("Bad stack operation")
			allGood = false
		}
	}
	finalize("Stack", allGood)
}

func TestLDA(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)

	ts := TestSuite{}
	ts.Add(TestData{code: "0200 A9 6E", res: 0x6E, flag: 0b00100000, resFlag: 0b00100000, cycles: 2})
	ts.Add(TestData{code: "0200 A9 FF", res: 0xFF, flag: 0b00100000, resFlag: 0b10100000, cycles: 2})
	ts.Add(TestData{code: "0200 A9 00", res: 0x00, flag: 0b00100000, resFlag: 0b00100010, cycles: 2})
	ts.Add(TestData{code: "0200 A9 81", res: 0x81, flag: 0b00100000, resFlag: 0b10100000, cycles: 2})

	RAM[0x05] = 0xAA
	ts.Add(TestData{code: "0200 B5 04", x: 0x01, res: 0xAA, flag: 0b00110000, resFlag: 0b10110000, cycles: 4})

	RAM[0x0FFF] = 0x01
	RAM[0x100F] = 0xEE
	ts.Add(TestData{code: "0200 BD FF 0F", x: 0x10, res: 0xEE, flag: 0b00110000, resFlag: 0b10110000, cycles: 5})
	ts.Add(TestData{code: "0200 BD FF 0F", x: 0x00, res: 0x01, flag: 0b00110000, resFlag: 0b00110000, cycles: 4})

	RAM[0x2211] = 0xE1
	RAM[0x0C] = 0x11
	RAM[0x0D] = 0x22
	ts.Add(TestData{code: "0200 A1 0A", x: 0x02, res: 0xE1, flag: 0b00110000, resFlag: 0b10110000, cycles: 6})

	RAM[0x2213] = 0xE2
	RAM[0x0A] = 0x11
	RAM[0x0B] = 0x22
	ts.Add(TestData{code: "0200 B1 0A", y: 0x02, res: 0xE2, flag: 0b00110000, resFlag: 0b10110000, cycles: 5})

	RAM[0x200F] = 0xE3
	RAM[0x0E] = 0xFF
	RAM[0x0F] = 0x1F
	ts.Add(TestData{code: "0200 B1 0E", y: 0x10, res: 0xE3, flag: 0b00110000, resFlag: 0b10110000, cycles: 6})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, proc.A, byte(table.res), "Assignement")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Status Flag")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestSTA(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)

	ts := TestSuite{}
	ts.Add(TestData{code: "0200 8D 00 04", acc: 0x01, destMem: 0x0400, flag: 0b00110000, resFlag: 0b00110000, cycles: 4})
	ts.Add(TestData{code: "0200 99 FF 0F", acc: 0x02, y: 0x04, destMem: 0x1003, flag: 0b00110000, resFlag: 0b00110000, cycles: 5})

	RAM[0x0A] = 0xFF
	RAM[0x0B] = 0x1F
	ts.Add(TestData{code: "0200 91 0A", acc: 0x02, y: 0x04, destMem: 0x2003, flag: 0b00110000, resFlag: 0b00110000, cycles: 6})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, RAM[table.destMem], proc.A, "Assignement")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Status Flag")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestEOR(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)
	ts := TestSuite{}
	RAM[0x08] = 0x80
	ts.Add(TestData{code: "0200 55 04", acc: 0x11, x: 0x04, flag: 0b00110000, res: 0x91, resFlag: 0b10110000, cycles: 4})
	ts.Add(TestData{code: "0200 55 04", acc: 0x80, x: 0x04, flag: 0b00110000, res: 0x00, resFlag: 0b00110010, cycles: 4})
	ts.Add(TestData{code: "0200 55 04", acc: 0x0F, x: 0x04, flag: 0b00110001, res: 0x8F, resFlag: 0b10110001, cycles: 4})
	ts.Add(TestData{code: "0200 55 04", acc: 0xFF, x: 0x04, flag: 0b00110001, res: 0x7F, resFlag: 0b00110001, cycles: 4})
	ts.Add(TestData{code: "0200 55 04", acc: 0x00, x: 0x04, flag: 0b00110001, res: 0x80, resFlag: 0b10110001, cycles: 4})

	RAM[0x0A] = 0xFF
	RAM[0x0B] = 0x1F
	RAM[0x2003] = 0x80
	ts.Add(TestData{code: "0200 51 0A", acc: 0x11, y: 0x04, flag: 0b00110000, res: 0x91, resFlag: 0b10110000, cycles: 6})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, proc.A, byte(table.res), "Result")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestASL(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)
	ts := TestSuite{}
	RAM[0x19] = 0xF1
	ts.Add(TestData{code: "0200 0E 19 00", destMem: 0x19, flag: 0b00110000, res: 0xE2, resFlag: 0b10110001, cycles: 6})

	RAM[0x14] = 0x80
	RAM[0x15] = 0x7F
	RAM[0x16] = 0x7F
	RAM[0x17] = 0x80
	RAM[0x18] = 0xFF
	ts.Add(TestData{code: "0200 16 10", x: 0x04, destMem: 0x14, flag: 0b00110000, res: 0x00, resFlag: 0b00110011, cycles: 6})
	ts.Add(TestData{code: "0200 16 11", x: 0x04, destMem: 0x15, flag: 0b00110000, res: 0xFE, resFlag: 0b10110000, cycles: 6})
	ts.Add(TestData{code: "0200 16 12", x: 0x04, destMem: 0x16, flag: 0b00110001, res: 0xFE, resFlag: 0b10110000, cycles: 6})
	ts.Add(TestData{code: "0200 16 13", x: 0x04, destMem: 0x17, flag: 0b00110001, res: 0x00, resFlag: 0b00110011, cycles: 6})
	ts.Add(TestData{code: "0200 16 14", x: 0x04, destMem: 0x18, flag: 0b00110001, res: 0xFE, resFlag: 0b10110001, cycles: 6})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, RAM[table.destMem], byte(table.res), "Result")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestLSR(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)
	ts := TestSuite{}
	RAM[0x14] = 0x80
	RAM[0x15] = 0x0F
	RAM[0x16] = 0x0F
	RAM[0x17] = 0x80
	RAM[0x18] = 0xFF
	ts.Add(TestData{code: "0200 56 10", x: 0x04, destMem: 0x14, flag: 0b00110000, res: 0x40, resFlag: 0b00110000, cycles: 6})
	ts.Add(TestData{code: "0200 56 11", x: 0x04, destMem: 0x15, flag: 0b00110000, res: 0x07, resFlag: 0b00110001, cycles: 6})
	ts.Add(TestData{code: "0200 56 12", x: 0x04, destMem: 0x16, flag: 0b00110001, res: 0x07, resFlag: 0b00110001, cycles: 6})
	ts.Add(TestData{code: "0200 56 13", x: 0x04, destMem: 0x17, flag: 0b00110001, res: 0x40, resFlag: 0b00110000, cycles: 6})
	ts.Add(TestData{code: "0200 56 14", x: 0x04, destMem: 0x18, flag: 0b00110001, res: 0x7F, resFlag: 0b00110001, cycles: 6})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, RAM[table.destMem], byte(table.res), "Result")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestADC(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)

	ts := TestSuite{}
	RAM[0x14] = 0x06
	ts.Add(TestData{code: "0200 75 10", acc: 0x01, x: 0x04, flag: 0b00110000, res: 0x07, resFlag: 0b00110000, cycles: 4})
	ts.Add(TestData{code: "0200 75 10", acc: 0x01, x: 0x04, flag: 0b00110001, res: 0x08, resFlag: 0b00110000, cycles: 4})
	ts.Add(TestData{code: "0200 75 10", acc: 0xFE, x: 0x04, flag: 0b00110000, res: 0x04, resFlag: 0b00110001, cycles: 4})
	ts.Add(TestData{code: "0200 75 10", acc: 0xFE, x: 0x04, flag: 0b00110001, res: 0x05, resFlag: 0b00110001, cycles: 4})

	ts.Add(TestData{code: "0200 69 80", acc: 0x78, flag: 0b00110000, res: 0xF8, resFlag: 0b10110000, cycles: 2})
	ts.Add(TestData{code: "0200 69 12", acc: 0x80, flag: 0b00111000, res: 0x92, resFlag: 0b10111000, cycles: 2})
	ts.Add(TestData{code: "0200 69 46", acc: 0x58, flag: 0b00111001, res: 0x05, resFlag: 0b01111001, cycles: 2})
	ts.Add(TestData{code: "0200 69 01", acc: 0x99, flag: 0b00111000, res: 0x00, resFlag: 0b00111011, cycles: 2})

	RAM[0x14] = 0x06
	RAM[0x15] = 0x02
	RAM[0x0206] = 0x0E
	ts.Add(TestData{code: "0200 61 10", acc: 0x20, x: 0x04, flag: 0b00110000, res: 0x2E, resFlag: 0b00110000, cycles: 6})
	ts.Add(TestData{code: "0200 61 10", acc: 0x01, x: 0x04, flag: 0b00110001, res: 0x10, resFlag: 0b00110000, cycles: 6})
	ts.Add(TestData{code: "0200 61 10", acc: 0xA0, x: 0x04, flag: 0b00110000, res: 0xAE, resFlag: 0b10110000, cycles: 6})
	ts.Add(TestData{code: "0200 61 10", acc: 0xFE, x: 0x04, flag: 0b00110001, res: 0x0D, resFlag: 0b00110001, cycles: 6})

	RAM[0x020A] = 0x0D
	ts.Add(TestData{code: "0200 71 14", acc: 0x20, y: 0x04, flag: 0b00110000, res: 0x2D, resFlag: 0b00110000, cycles: 5})
	ts.Add(TestData{code: "0200 71 14", acc: 0x01, y: 0x04, flag: 0b00110001, res: 0x0F, resFlag: 0b00110000, cycles: 5})
	ts.Add(TestData{code: "0200 71 14", acc: 0xA0, y: 0x04, flag: 0b00110000, res: 0xAD, resFlag: 0b10110000, cycles: 5})
	ts.Add(TestData{code: "0200 71 14", acc: 0xFE, y: 0x04, flag: 0b00110001, res: 0x0C, resFlag: 0b00110001, cycles: 5})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, proc.A, byte(table.res), "Result")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}

func TestSBC(t *testing.T) {
	var allGood bool = true
	mem.Clear(RAM)

	ts := TestSuite{}
	ts.Add(TestData{code: "0200 E9 08", acc: 0x03, flag: 0b00110000, res: 0xFA, resFlag: 0b10110000, cycles: 2})
	ts.Add(TestData{code: "0200 E9 08", acc: 0x03, flag: 0b00110001, res: 0xFB, resFlag: 0b10110000, cycles: 2})
	ts.Add(TestData{code: "0200 E9 46", acc: 0x58, flag: 0b00111000, res: 0x11, resFlag: 0b00111001, cycles: 2})

	RAM[0x14] = 0x06
	ts.Add(TestData{code: "0200 F5 10", acc: 0x01, x: 0x04, flag: 0b00110000, res: 0xFA, resFlag: 0b10110000, cycles: 4})
	ts.Add(TestData{code: "0200 F5 10", acc: 0x20, x: 0x04, flag: 0b00110000, res: 0x19, resFlag: 0b00110001, cycles: 4})

	RAM[0x15] = 0x02
	RAM[0x0206] = 0x0E
	ts.Add(TestData{code: "0200 E1 10", acc: 0x20, x: 0x04, flag: 0b00110000, res: 0x11, resFlag: 0b00110001, cycles: 6})
	ts.Add(TestData{code: "0200 E1 10", acc: 0x01, x: 0x04, flag: 0b00110001, res: 0xF3, resFlag: 0b10110000, cycles: 6})
	ts.Add(TestData{code: "0200 E1 10", acc: 0xA0, x: 0x04, flag: 0b00110000, res: 0x91, resFlag: 0b10110001, cycles: 6})
	ts.Add(TestData{code: "0200 E1 10", acc: 0xFE, x: 0x04, flag: 0b00110001, res: 0xF0, resFlag: 0b10110001, cycles: 6})

	RAM[0x020A] = 0x0E
	ts.Add(TestData{code: "0200 F1 14", acc: 0x20, y: 0x04, flag: 0b00110000, res: 0x11, resFlag: 0b00110001, cycles: 5})
	ts.Add(TestData{code: "0200 F1 14", acc: 0x01, y: 0x04, flag: 0b00110001, res: 0xF3, resFlag: 0b10110000, cycles: 5})
	ts.Add(TestData{code: "0200 F1 14", acc: 0xA0, y: 0x04, flag: 0b00110000, res: 0x91, resFlag: 0b10110001, cycles: 5})
	ts.Add(TestData{code: "0200 F1 14", acc: 0xFE, y: 0x04, flag: 0b00110001, res: 0xF0, resFlag: 0b10110001, cycles: 5})
	ts.Add(TestData{code: "0200 F1 14", acc: 0xFE, y: 0x04, flag: 0b00110011, res: 0xF0, resFlag: 0b10110001, cycles: 5})
	RAM[0x020C] = 0x08
	ts.Add(TestData{code: "0200 F1 14", acc: 0x03, y: 0x06, flag: 0b00110001, res: 0xFB, resFlag: 0b10110000, cycles: 5})
	ts.Add(TestData{code: "0200 F1 14", acc: 0x03, y: 0x06, flag: 0b00110000, res: 0xFA, resFlag: 0b10110000, cycles: 5})

	for _, table := range ts.data {
		table.run()
		allGood = allGood && table.checkByte(t, proc.A, byte(table.res), "Result")
		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
		allGood = allGood && table.checkCycles(t, "Cycles")
	}
	finalize(proc.Inst.Name, allGood)
}


// func TestBNE(t *testing.T) {
// 	var allGood bool = true
// 	mem.Clear(RAM)
// 	ts := TestSuite{proc: &proc, inst: 0xD0}

// 	ts.Add(TestData{flag: 0b00000000, pc: 0xBC16, oper: 0xF9, res: 0xBC11})
// 	ts.Add(TestData{flag: 0b00000010, pc: 0xBC16, oper: 0xF9, res: 0xBC18})

// 	for _, table := range ts.data {
// 		table.run()
// 		allGood = allGood && table.checkWord(t, proc.PC, table.res, "Address")
// 	}
// 	finalize(proc.Mnemonic[ts.inst].Name, allGood)
// }


// func TestCMP(t *testing.T) {
// 	var allGood bool = true
// 	mem.Clear(RAM)

// 	ts := TestSuite{proc: &proc, inst: 0xC9}
// 	ts.Add(TestData{acc: 0x50, oper: 0x20, flag: 0b00110000, resFlag: 0b00110001})
// 	ts.Add(TestData{acc: 0xF0, oper: 0x20, flag: 0b00110000, resFlag: 0b10110001})
// 	ts.Add(TestData{acc: 0x00, oper: 0x20, flag: 0b00110000, resFlag: 0b10110000})
// 	ts.Add(TestData{acc: 0x20, oper: 0x20, flag: 0b00110000, resFlag: 0b00110011})
// 	ts.Add(TestData{acc: 0x01, oper: 0x20, flag: 0b00110000, resFlag: 0b10110000})
// 	ts.Add(TestData{acc: 0x00, oper: 0x00, flag: 0b00110000, resFlag: 0b00110011})
// 	ts.Add(TestData{acc: 0xFF, oper: 0xFF, flag: 0b00110000, resFlag: 0b00110011})

// 	ts.inst = 0xD1
// 	ts.Add(TestData{acc: 0x50, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110000})
// 	ts.Add(TestData{acc: 0xF0, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110001})
// 	ts.Add(TestData{acc: 0x00, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110000})
// 	ts.Add(TestData{acc: 0x20, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110000})
// 	ts.Add(TestData{acc: 0xEE, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110011})
// 	ts.Add(TestData{acc: 0xFF, y: 0x08, oper: 0xC1, flag: 0b00110000, resFlag: 0b00110001})

// 	for _, table := range ts.data {
// 		proc.ram.Write(0x0408, 0xEE)
// 		proc.ram.Write(0xC1, 0x00)
// 		proc.ram.Write(0xC2, 0x04)
// 		table.run()
// 		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
// 	}
// 	finalize(proc.Mnemonic[ts.inst].Name, allGood)
// }

// func TestROR(t *testing.T) {
// 	var allGood bool = true
// 	mem.Clear(RAM)
// 	ts := TestSuite{proc: &proc, inst: 0x76}

// 	ts.Add(TestData{mem: 0x06, x: 0x04, oper: 0x10, flag: 0b00110000, res: 0x03, resFlag: 0b00110000})
// 	ts.Add(TestData{mem: 0x06, x: 0x04, oper: 0x10, flag: 0b00110001, res: 0x83, resFlag: 0b10110000})

// 	for _, table := range ts.data {
// 		proc.ram.Write(0x0014, table.mem)
// 		table.run()
// 		allGood = allGood && table.checkByte(t, proc.ram.Read(0x0014), byte(table.res), "Result")
// 		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
// 	}
// 	finalize(proc.Mnemonic[ts.inst].Name, allGood)
// }

// func TestROL(t *testing.T) {
// 	var allGood bool = true
// 	mem.Clear(RAM)
// 	ts := TestSuite{proc: &proc, inst: 0x36}
// 	ts.Add(TestData{mem: 0x06, x: 0x04, oper: 0x10, flag: 0b00110000, res: 0x0C, resFlag: 0b00110000})
// 	ts.Add(TestData{mem: 0x06, x: 0x04, oper: 0x10, flag: 0b00110001, res: 0x0D, resFlag: 0b00110000})
// 	ts.Add(TestData{mem: 0x80, x: 0x04, oper: 0x10, flag: 0b00110001, res: 0x01, resFlag: 0b00110001})
// 	ts.Add(TestData{mem: 0xF0, x: 0x04, oper: 0x10, flag: 0b00110001, res: 0xE1, resFlag: 0b10110001})
// 	ts.Add(TestData{mem: 0xF0, x: 0x04, oper: 0x10, flag: 0b00110000, res: 0xE0, resFlag: 0b10110001})

// 	for _, table := range ts.data {
// 		proc.ram.Write(0x0014, table.mem)
// 		table.run()
// 		allGood = allGood && table.checkByte(t, proc.ram.Read(0x0014), byte(table.res), "Result")
// 		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
// 	}
// 	finalize(proc.Mnemonic[ts.inst].Name, allGood)
// }

// func TestBIT(t *testing.T) {
// 	var allGood bool = true
// 	// LDA #$80
// 	// STA $14
// 	// CLC
// 	// LDA #$11
// 	// BIT $14
// 	mem.Clear(RAM)
// 	ts := TestSuite{proc: &proc, inst: 0x24}

// 	ts.Add(TestData{mem: 0x80, acc: 0x11, oper: 0x14, flag: 0b00110000, resFlag: 0b10110010})
// 	ts.Add(TestData{mem: 0x80, acc: 0x80, oper: 0x14, flag: 0b00110000, resFlag: 0b10110000})
// 	ts.Add(TestData{mem: 0x80, acc: 0x0F, oper: 0x14, flag: 0b00110001, resFlag: 0b10110011})
// 	ts.Add(TestData{mem: 0x80, acc: 0xFF, oper: 0x14, flag: 0b00110001, resFlag: 0b10110001})
// 	ts.Add(TestData{mem: 0x80, acc: 0x00, oper: 0x14, flag: 0b00110011, resFlag: 0b10110011})

// 	for _, table := range ts.data {
// 		proc.ram.Write(0x0014, table.mem)
// 		table.run()
// 		allGood = allGood && table.checkBit(t, proc.S, table.resFlag, "Flags")
// 	}
// 	finalize(proc.Mnemonic[ts.inst].Name, allGood)
// }
