package mos6510

import "fmt"

///////////////////////////////////////////////////////
//                        PHP                        //
///////////////////////////////////////////////////////

func (C *CPU) PHP_imp() {
	fmt.Printf("PHP\n")
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.stack[C.SP] = C.S | ^B_mask | ^U_mask
		C.SP--
		C.CycleCount = 0
		C.StackDebugPt++
		C.StackDebug[C.StackDebugPt] = fmt.Sprintf("%04X: PHP %02X -> %02X:%02X\n", C.InstStart, C.S, C.SP+1, C.stack[C.SP+1])
	}
}

///////////////////////////////////////////////////////
//                        PLP                        //
///////////////////////////////////////////////////////

func (C *CPU) PLP_imp() {
	fmt.Printf("PLP\n")
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.ram.Read(C.PC) // Read and forget
	case 3:
		C.SP++
	case 4:
		C.S = C.stack[C.SP] & B_mask & U_mask
		C.CycleCount = 0
		C.StackDebugPt--
	}
}
