package builder

type AsusCollector struct {
	Brand   string
	Core    int
	Memory  int
	Monitor bool
	GUI     int
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetGUI() {
	collector.GUI = 4090
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = true
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Brand:   collector.Brand,
		Core:    collector.Core,
		Memory:  collector.Memory,
		Monitor: collector.Monitor,
		GUI:     collector.GUI,
	}
}
