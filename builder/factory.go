package builder

type Factory struct {
	Collector Collector
}

func NewFactory(col Collector) *Factory {
	return &Factory{Collector: col}
}

func (factory *Factory) SetCollector(col Collector) {
	factory.Collector = col
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMonitor()
	factory.Collector.SetMonitor()
	factory.Collector.SetGUI()
	factory.Collector.SetBrand()
	return factory.Collector.GetComputer()
}
