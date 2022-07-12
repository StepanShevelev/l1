package main

import "fmt"

//Сервис отправки xml документов
type dataService interface {
	SendXmlData()
}

type user struct {
}

func (u *user) send(d dataService) {
	fmt.Println("Пользователь хочет отправить xml документ")
	d.SendXmlData()
}

//Структура у которой есть возможность отправить документ в формате xml
type xmlDocument struct {
}

func (x *xmlDocument) SendXmlData() {

	fmt.Println("Отправка xml документа")
}

//Структура которой нужен адаптер
type jsonDocument struct {
}

func (j *jsonDocument) ConvertToXml() {

	fmt.Println("Документ преобразован в xml ")
}

// Адаптер для структуры без метода отправки xml
type jsonDocumentAdapter struct {
	jsonDoc *jsonDocument
}

func (adapter *jsonDocumentAdapter) SendXmlData() {
	adapter.jsonDoc.ConvertToXml()
	fmt.Println("Отправка xml через адаптер")
}

func main() {
	Stepan := &user{}
	//структура с возможностью отправки
	xmlSender := &xmlDocument{}
	Stepan.send(xmlSender)

	//структура без возможности отправки
	jsonToXml := &jsonDocument{}
	Adapter := &jsonDocumentAdapter{
		jsonDoc: jsonToXml,
	}
	// Теперь устройство без метода через адаптер смогло воспользоваться методом
	Stepan.send(Adapter)
}
