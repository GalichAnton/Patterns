package builder

import "fmt"

type Computer struct {
	Brand   string
	Core    int
	Memory  int
	Monitor bool
	GUI     int
}

func (c Computer) Print() {
	fmt.Printf("%s Core[%d], Memory[%d], GUI[%d], Monitor[%v]\n", c.Brand, c.Core, c.Memory, c.GUI, c.Monitor)
}
