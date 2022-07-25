package mem2

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
	LayerNum    int
	Val         *byte
	Under       *byte
	ReadOnly    bool
	Accessor    MEMAccess
	UnderAccess MEMAccess
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

func InitConfig(size int) CONFIG {
	C := CONFIG{}

	C.StorageRef = make([][]byte, 0, 20)
	C.LayersName = make(map[string]int)
	C.NameLayers = make(map[int]string)
	C.Start = make([]uint16, 0, 20)
	C.Size = make([]int, 0, 20)
	C.ReadOnlyMode = make([]bool, 0, 20)
	C.Disabled = make([]bool, 0, 20)
	C.Accessors = make([]MEMAccess, 0, 20)

	C.VisibleMem = make([]MEMCell, size)
	return C
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
	C.Accessors = append(C.Accessors, accessor)

	if disabled == false {
		if accessor == nil {
			accessor = &DefaultAccessor{}
		}
		for i := range C.StorageRef[layerNum] {
			C.VisibleMem[int(start)+i].LayerNum = layerNum
			C.VisibleMem[int(start)+i].Under = C.VisibleMem[int(start)+i].Val
			C.VisibleMem[int(start)+i].UnderAccess = C.VisibleMem[int(start)+i].Accessor

			C.VisibleMem[int(start)+i].Val = &C.StorageRef[layerNum][i]
			C.VisibleMem[int(start)+i].ReadOnly = mode
			C.VisibleMem[int(start)+i].Accessor = accessor
		}
	}
}

func (C *CONFIG) Disable(layerName string) {
	var underLayerNum, cpt int
	actualLayerNum := C.LayersName[layerName]

}

func (C *CONFIG) Enable(layerName string) {
	C.Disabled[C.LayersName[layerName]] = false
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
