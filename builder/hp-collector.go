package builder

type HPCollector struct {
	Brand   string
	Core    int
	Memory  int
	Monitor bool
	GUI     int
}

func (collector *HPCollector) SetBrand() {
	collector.Brand = "HP"
}

func (collector *HPCollector) SetCore() {
	collector.Core = 8
}

func (collector *HPCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HPCollector) SetGUI() {
	collector.GUI = 3060
}

func (collector *HPCollector) SetMonitor() {
	collector.Monitor = true
}

func (collector *HPCollector) GetComputer() Computer {
	return Computer{
		Brand:   collector.Brand,
		Core:    collector.Core,
		Memory:  collector.Memory,
		Monitor: collector.Monitor,
		GUI:     collector.GUI,
	}
}
