package builder

const (
	AsusBrand = "asus"
	HPBrand   = "hp"
)

type Collector interface {
	SetBrand()
	SetCore()
	SetMemory()
	SetGUI()
	GetComputer() Computer
	SetMonitor()
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusBrand:
		return &AsusCollector{}
	case HPBrand:
		return &HPCollector{}
	}
}
