package fabric

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (comp Server) GetType() string {
	return comp.Type
}

func (comp Server) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory:[%d]\n", comp.Type, comp.Core, comp.Memory)
}
