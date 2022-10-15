package mem2

import "fmt"

const (
	PAGE_DIVIDER = 8
	READWRITE    = false
	READONLY     = true
)

type MEMAccess interface {
	MRead([]MEMCell, uint16) byte
	MWrite([]MEMCell, uint16, byte)
	MWriteUnder([]MEMCell, uint16, byte)
}

type MEMCell struct {
	LayerNum      int
	UnderLayerNum int
	Val           *byte
	Under         *byte
	Accessor      MEMAccess
	UnderAccess   MEMAccess
}

type CONFIG struct {
	StorageRef   [][]byte       // Liste des zones buffer attachées
	LayersName   map[string]int // Nom de la couche
	NameLayers   map[int]string // Nom de la couche
	Start        []uint16       // Addresse de début de la couche
	Size         []int          // Taille de zone buffer
	ReadOnlyMode []bool         // Mode d'accès à la couche
	Disabled     []bool         // Momentanement invisible, on bascule en couche 0 (RAM)
	Accessors    []MEMAccess    // Reader/Writer de la couche
	VisibleMem   []MEMCell      // Pointer compilé
}

func (M *MEMCell) dump() {
	fmt.Printf("== Dumping MemCell ==\n")
	fmt.Printf("LayerNum: %d\n", M.LayerNum)
	fmt.Printf("Value: %02X (%d)\n", *M.Val, M.Val)
	fmt.Printf("Accessor addr: %d\n", M.Accessor)
	fmt.Printf("UnderNum: %d\n", M.UnderLayerNum)
	fmt.Printf("UnderVlue: %02X (%d)\n", *M.Under, M.Under)
}

func (C *CONFIG) InitConfig(size uint) {
	C.StorageRef = make([][]byte, 0, 20)
	C.LayersName = make(map[string]int)
	C.NameLayers = make(map[int]string)
	C.Start = make([]uint16, 0, 20)
	C.Size = make([]int, 0, 20)
	C.ReadOnlyMode = make([]bool, 0, 20)
	C.Disabled = make([]bool, 0, 20)
	C.Accessors = make([]MEMAccess, 0, 20)

	C.VisibleMem = make([]MEMCell, size)
}

func (C *CONFIG) Attach(name string, start uint16, content []byte, mode bool, disabled bool, accessor MEMAccess) {
	C.StorageRef = append(C.StorageRef, content)
	layerNum := len(C.StorageRef) - 1

	C.LayersName[name] = layerNum
	C.NameLayers[layerNum] = name

	C.Start = append(C.Start, start)
	C.Size = append(C.Size, len(content))
	C.ReadOnlyMode = append(C.ReadOnlyMode, mode)
	C.Disabled = append(C.Disabled, disabled)
	if accessor == nil {
		accessor = &DefaultAccessor{}
	}
	C.Accessors = append(C.Accessors, accessor)

	if disabled == false {
		for i := range C.StorageRef[layerNum] {
			startPage := int(start) + i
			C.VisibleMem[startPage].LayerNum = layerNum

			C.VisibleMem[startPage].UnderLayerNum = C.VisibleMem[startPage].LayerNum
			C.VisibleMem[startPage].Under = C.VisibleMem[startPage].Val
			C.VisibleMem[startPage].UnderAccess = C.VisibleMem[startPage].Accessor
			C.VisibleMem[startPage].Val = &C.StorageRef[layerNum][i]
			C.VisibleMem[startPage].Accessor = accessor
		}
	}
}

func (C *CONFIG) Disable(layerName string) {
	actualLayerNum := C.LayersName[layerName]
	if C.Disabled[actualLayerNum] == false {
		// for i := C.Start[actualLayerNum]; i < uint16(C.Size[actualLayerNum]); i++ {
		for i := range C.StorageRef[actualLayerNum] {
			startPage := int(C.Start[actualLayerNum]) + i
			if C.VisibleMem[startPage].LayerNum != actualLayerNum {
				continue
			}
			C.VisibleMem[startPage].LayerNum = C.VisibleMem[startPage].UnderLayerNum
			C.VisibleMem[startPage].Val = C.VisibleMem[startPage].Under
			C.VisibleMem[startPage].Accessor = C.VisibleMem[startPage].UnderAccess
		}
		C.Disabled[actualLayerNum] = true
	}
}

func (C *CONFIG) Enable(layerName string) {
	actualLayerNum := C.LayersName[layerName]
	if C.Disabled[actualLayerNum] == true {
		for i := range C.StorageRef[actualLayerNum] {
			startPage := int(C.Start[actualLayerNum]) + i
			if C.VisibleMem[startPage].LayerNum > actualLayerNum {
				continue
			}
			C.VisibleMem[startPage].LayerNum = actualLayerNum
			C.VisibleMem[startPage].UnderLayerNum = C.VisibleMem[startPage].LayerNum
			C.VisibleMem[startPage].Under = C.VisibleMem[startPage].Val
			C.VisibleMem[startPage].UnderAccess = C.VisibleMem[startPage].Accessor
			C.VisibleMem[startPage].Val = &C.StorageRef[actualLayerNum][i]
			C.VisibleMem[startPage].Accessor = C.Accessors[actualLayerNum]
		}
		C.Disabled[actualLayerNum] = false
	}
}

func (C *CONFIG) ReadOnly(layerName string) {
	C.ReadOnlyMode[C.LayersName[layerName]] = true
}

func (C *CONFIG) ReadWrite(layerName string) {
	C.ReadOnlyMode[C.LayersName[layerName]] = false
}

// func (C *CONFIG) Show() {
// 	var line [128]string
// 	var pos int
// 	var indice float32 = 128 / float32(C.TotalPages)

// 	for layerNum := range C.Layers {
// 		for page := 0; page < C.TotalPages; page++ {
// 			pos = int(indice * float32(page))
// 			if C.PagesUsed[layerNum][page] {
// 				line[pos] = clog.CSprintf("black", "green", " ")
// 			} else {
// 				line[pos] = clog.CSprintf("black", "dark_gray", " ")
// 			}
// 		}
// 		fmt.Printf("%10s:", C.NameLayers[layerNum])
// 		for i := range line {
// 			fmt.Printf("%s", line[i])
// 		}
// 		fmt.Println()
// 	}
// 	clog.CPrintf("dark_gray", "black", "\n%12s", " ")
// 	clog.CPrintf("black", "green", "%s", " Read/Write ")
// 	clog.CPrintf("black", "black", "%s", "  ")
// 	clog.CPrintf("black", "yellow", "%s", " Read Only ")
// 	clog.CPrintf("black", "black", "%s", "  ")
// 	clog.CPrintf("black", "red", "%s", " Write Only ")
// 	clog.CPrintf("black", "black", "%s", "  ")
// 	clog.CPrintf("black", "light_gray", "%s", " Masked ")
// 	clog.CPrintf("black", "black", "%s", " ")
// 	fmt.Printf("\n\n")
// }
