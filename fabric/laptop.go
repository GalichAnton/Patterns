package fabric

import "fmt"

type Laptop struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewLaptop() Computer {
	return Laptop{
		Type:    NoteBookType,
		Core:    6,
		Memory:  8,
		Monitor: true,
	}
}

func (comp Laptop) GetType() string {
	return comp.Type
}

func (comp Laptop) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory:[%d], Monitor[%v]\n", comp.Type, comp.Core, comp.Memory, comp.Monitor)
}
