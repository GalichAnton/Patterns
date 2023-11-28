package data

type JSONDoc struct {
}

func (doc JSONDoc) ConvertToXml() string {
	return "<xml></xml>"
}

type JSONDocAdapter struct {
	jSONDoc *JSONDoc
}

func (adapter JSONDocAdapter) SendXmlData() {
	adapter.jSONDoc.ConvertToXml()
	println("Отправка xml")
}
