package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

type Node struct {
	Name     string  `json:"Name"`
	Pid      int32   `json:"Pid"`
	MemUsage float32 `json:"MemUsage"`
}

type ProcessesTree struct {
	Childs []*ProcessesTree `json:"Childs"`
	Node   *Node            `json:"Node"`
}

// Para convertir un nodo process.Process a un nodo parseable a json
func convertNode(nodo *process.Process) *Node {
	name, _ := nodo.Name()
	mem, _ := nodo.MemoryPercent()
	return &Node{
		Name:     name,
		Pid:      nodo.Pid,
		MemUsage: mem,
	}
}

// Inserta un nuevo nodo dentro del arbol de nodos
func (p *ProcessesTree) insertarNodo(nodo *process.Process, parentId int32) {
	newNode := ProcessesTree{
		Node:   convertNode(nodo),
		Childs: []*ProcessesTree{}}

	// los procesos con padre 0 penden de la raiz
	if parentId == 0 {
		p.Childs = append(p.Childs, &newNode)
		return
	}
	// revisar al padre de manera recursiva
	for _, e := range p.Childs {
		// buscar el padre entre entre los hijos
		if e.Node.Pid == parentId {
			e.Childs = append(e.Childs, &newNode)
			return
		}
		// si no se encuentra hacerlo de manera recursiva
		e.insertarNodo(nodo, parentId)
	}
}

// Crea un arbol de nodos para luego poder mandarlo de manera mas organizada
func createNewTree() (*ProcessesTree, error) {
	ps, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("No se puede traer los procesos")
	}

	// Como los nodos ya vienen ordenados, insertamos directamente
	raiz := ProcessesTree{Node: nil, Childs: nil}
	for _, e := range ps {
		parent, _ := e.Ppid()
		raiz.insertarNodo(e, parent)
	}
	return &raiz, nil
}

func GetData() (*ProcessesTree, error) {
    res, err := createNewTree()
    if err != nil {
        return nil, err
    }
    return res, nil
}

// Actualiza la lista de procesos cada 200 Millisegundos
/* func InicializarSistema(ch chan *ProcessesTree) {
	for {
		time.Sleep(time.Millisecond * 200)
		tree, err := createNewTree()
		if err == nil {
			ch <- tree
		}
	}
} */
