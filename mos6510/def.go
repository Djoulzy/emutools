package mos6510

import "github.com/Djoulzy/mmu"

const (
	C_mask byte = 0b11111110
	Z_mask byte = 0b11111101
	I_mask byte = 0b11111011
	D_mask byte = 0b11110111
	B_mask byte = 0b11101111
	U_mask byte = 0b11011111
	V_mask byte = 0b10111111
	N_mask byte = 0b01111111

	NMI_Vector       = 0xFFFA
	COLDSTART_Vector = 0xFFFC // Go to 0xFCE2
	IRQBRK_Vector    = 0xFFFE
)

var regString [8]string = [8]string{"C", "Z", "I", "D", "B", "U", "V", "N"}
var StackStart uint16 = 0x0100

type addressing int

const (
	implied addressing = iota
	immediate
	relative
	zeropage
	zeropageX
	zeropageY
	indirectX
	indirectY
	indirectzp
	indirect
	absolute
	absoluteX
	absoluteY
	Branching
	CrossPage
)

var InstTemplate [15]string = [15]string{
	"",
	"#$%02X",
	"$%02X",
	"$%02X",
	"$%02X,X",
	"$%02X,Y",
	"($%02X,X)",
	"($%02X),Y",
	"($%02X)",
	"($%02X%02X)",
	"$%02X%02X",
	"$%02X%02X,X",
	"$%02X%02X,Y",
	"$%02X%02X",
	"$%02X%02X",
}

type Instruction struct {
	Name   string
	addr   addressing
	bytes  int
	Cycles int
	action func()
}

type cpuState int

const (
	Idle cpuState = iota
	IRQ1
	IRQ2
	IRQ3
	IRQ4
	IRQ5
	IRQ6
	IRQ7
	NMI1
	NMI2
	NMI3
	NMI4
	NMI5
	NMI6
	NMI7
)

// type Manager interface {
// 	Attach(int, string, uint16, []byte, bool, bool, interface{})
// 	Read(uint16) byte
// 	Write(uint16, byte)
// 	GetFullSize() int
// 	GetStack(uint16, uint16) []mem.MEMCell
// 	Enable(string)
// 	Disable(string)
// 	ReadOnly(string)
// 	ReadWrite(string)
// 	Clear([]byte, int, byte)
// 	LoadROM(int, string) []byte
// 	Dump(uint16)
// }

// CPU :
type CPU struct {
	model   string
	PC      uint16
	SP      byte
	A       byte
	X       byte
	Y       byte
	S       byte
	IRQ_pin int
	NMI_pin int

	Mnemonic map[byte]Instruction

	ram       *mmu.MMU
	InstStart uint16
	instDump  string
	instCode  byte
	FullInst  string
	FullDebug string
	Inst      Instruction

	OperHI      byte
	OperLO      byte
	Pointer     byte
	IndAddrLO   byte
	IndAddrHI   byte
	tmpBuff     byte
	pageCrossed bool

	Oper         uint16
	cross_oper   uint16
	val_zp_lo    byte
	val_zp_hi    byte
	val_absolute byte
	val_absXY    byte
	comp_result  byte

	CycleCount int

	NMI_Raised bool
	IRQ_Raised bool
	INT_delay  bool
}
