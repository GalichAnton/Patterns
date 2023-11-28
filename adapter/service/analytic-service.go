package service

type AnalyticService interface {
	SendXmlData()
}
type XmlDoc struct {
}

func (doc XmlDoc) SendXmlData() {
	println("Отправка xml")
}
