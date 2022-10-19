package mem_v1

import (
	"fmt"

	"github.com/Djoulzy/Tools/clog"
)

const (
	PAGE_DIVIDER = 8
	READWRITE    = false
	READONLY     = true
)

type MEMCells interface{}

type MEMAccess interface {
	MRead(MEMCells, uint16) byte
	MWrite(MEMCells, uint16, byte)
}

type CONFIG struct {
	Layers       [][]byte       // Liste des couches de memoire
	LayersName   map[string]int // Nom de la couche
	NameLayers   map[int]string // Nom de la couche
	Start        []uint16       // Addresse de début de la couche
	PagesUsed    [][]bool       // Pages utilisées par la couche
	ReadOnlyMode []bool         // Mode d'accès à la couche
	Disabled     []bool         // Momentanement invisible, on bascule en couche 0 (RAM)
	LayerByPages [][]int        // Couche active pour la page
	Accessors    []MEMAccess    // Reader/Writer de la couche
	TotalPages   int            // Nb total de pages
}

func (C *CONFIG) InitConfig(size uint) {
	C.Layers = make([][]byte, 0, 20)
	C.LayersName = make(map[string]int)
	C.NameLayers = make(map[int]string)
	C.Start = make([]uint16, 0, 20)
	C.PagesUsed = make([][]bool, 0, 20)
	C.ReadOnlyMode = make([]bool, 0, 20)
	C.Disabled = make([]bool, 0, 20)
	C.Accessors = make([]MEMAccess, 0, 20)

	C.TotalPages = int(size >> PAGE_DIVIDER)
	C.LayerByPages = make([][]int, C.TotalPages)
}

func (C *CONFIG) Attach(name string, start uint16, content []byte, mode bool, disabled bool, accessor MEMAccess) {
	nbPages := len(content) >> PAGE_DIVIDER
	startPage := int(start >> PAGE_DIVIDER)

	C.Layers = append(C.Layers, content)
	layerNum := len(C.Layers) - 1

	C.LayersName[name] = layerNum
	C.NameLayers[layerNum] = name

	C.Start = append(C.Start, start)
	C.ReadOnlyMode = append(C.ReadOnlyMode, mode)
	C.Disabled = append(C.Disabled, disabled)
	if accessor == nil {
		accessor = &DefaultAccessor{}
	}
	C.Accessors = append(C.Accessors, accessor)

	C.PagesUsed = append(C.PagesUsed, make([]bool, C.TotalPages))
	for i := 0; i < C.TotalPages; i++ {
		C.PagesUsed[layerNum][i] = false
	}

	for i := 0; i < nbPages; i++ {
		C.LayerByPages[startPage+i] = append([]int{layerNum}, C.LayerByPages[startPage+i]...)
		C.PagesUsed[layerNum][startPage+i] = true
	}
}

func (C *CONFIG) Disable(layerName string) {
	C.Disabled[C.LayersName[layerName]] = true
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

func (C *CONFIG) Accessor(layerName string, access MEMAccess) {
	C.Accessors[C.LayersName[layerName]] = access
}

func (C *CONFIG) Show() {
	var line [128]string
	var pos int
	var indice float32 = 128 / float32(C.TotalPages)

	for layerNum := range C.Layers {
		for page := 0; page < C.TotalPages; page++ {
			pos = int(indice * float32(page))
			if C.PagesUsed[layerNum][page] {
				line[pos] = clog.CSprintf("black", "green", " ")
			} else {
				line[pos] = clog.CSprintf("black", "dark_gray", " ")
			}
		}
		fmt.Printf("%10s:", C.NameLayers[layerNum])
		for i := range line {
			fmt.Printf("%s", line[i])
		}
		fmt.Println()
	}
	clog.CPrintf("dark_gray", "black", "\n%12s", " ")
	clog.CPrintf("black", "green", "%s", " Read/Write ")
	clog.CPrintf("black", "black", "%s", "  ")
	clog.CPrintf("black", "yellow", "%s", " Read Only ")
	clog.CPrintf("black", "black", "%s", "  ")
	clog.CPrintf("black", "red", "%s", " Write Only ")
	clog.CPrintf("black", "black", "%s", "  ")
	clog.CPrintf("black", "light_gray", "%s", " Masked ")
	clog.CPrintf("black", "black", "%s", " ")
	fmt.Printf("\n\n")
}
