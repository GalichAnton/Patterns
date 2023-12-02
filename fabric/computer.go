package fabric

import "fmt"

const (
	ServerType   = "server"
	PCType       = "personal"
	NoteBookType = "laptop"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case NoteBookType:
		return NewLaptop()
	case PCType:
		return NewPC()
	default:
		fmt.Printf("%s Несуществующий тип\n", typeName)
		return nil
	}
}
