package third_party

import "fmt"

type Soldier interface {
	Attack() int
	Defence() int
}

type BasicSoldier struct {
}

func (b BasicSoldier) Attack() int {
	return 1
}

func (b BasicSoldier) Defence() int {
	return 2
}

func DisplaySoldier(s Soldier) {
	fmt.Printf("Soildr stats: atack [%d], defence [%d]\n", s.Attack(), s.Defence())
}
