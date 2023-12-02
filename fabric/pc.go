package fabric

import "fmt"

type PC struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPC() Computer {
	return PC{
		Type:    PCType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (comp PC) GetType() string {
	return comp.Type
}

func (comp PC) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory:[%d], Monitor: [%v]\n", comp.Type, comp.Core, comp.Memory, comp.Monitor)
}
