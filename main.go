package main

import (
	"patterns/builder"
	"patterns/fabric"
)

var computers = []string{"personal", "laptop", "server"}

func main() {
	for _, comp := range computers {
		fabric.New(comp).PrintDetails()
	}

	asusCollector := builder.GetCollector("asus")
	hpCollector := builder.GetCollector("hp")

	factory := builder.NewFactory(asusCollector)

	asusComp := factory.CreateComputer()
	asusComp.Print()

	factory.SetCollector(hpCollector)

	hpComp := factory.CreateComputer()
	hpComp.Print()
}
