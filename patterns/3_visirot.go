package patterns

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Паттерн Посетитель - это паттерн проектирования, который позволяет выполнить какое-то действие над группой объектов
		различных типов, при этом не изменяя их структуру.

	Плюсы:
		+ Операции, выполняемые над объектами, выносятся в отдельные классы, что облегчает поддержку и расширение кода.
		+ Можно добавлять новые операции, не изменяя существующий код объектов.
		+ Позволяет избежать загрязнения классов объектов кодом операций

	Минусы:
		- Усложняет структуру программы
		- Посетитель может получать доступ к приватным членам объектов, что может нарушить инкапсуляцию
		- Дополнительная сложность при добавлении новых типов объектов

	Пример ниже.
*/

// Element - Интерфейс элемента, который может быть посещен посетителем
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElement - Конкретная реализация элемента
type ConcreteElement struct {
	data int
}

// Accept - Метод посещения
func (e *ConcreteElement) Accept(visitor Visitor) {
	visitor.VisitConcreteElement(e)
}

// GetData - Метод получения данных
func (e *ConcreteElement) GetData() int {
	return e.data
}

// Visitor - Интерфейс посетителя
type Visitor interface {
	VisitConcreteElement(element *ConcreteElement)
}

// ConcreteVisitor - Конкретная реализация посетителя
type ConcreteVisitor struct{}

// VisitConcreteElement - Метод посещения конкретного элемента
func (v *ConcreteVisitor) VisitConcreteElement(element *ConcreteElement) {
	fmt.Printf("Visited element with data: %d\n", element.GetData())
}

// VisitorTest - Тест посетителя
func VisitorTest() {
	// Создаем элемент
	element := &ConcreteElement{data: 42}

	// Создаем посетителя
	visitor := &ConcreteVisitor{}

	// Посещаем элемент
	element.Accept(visitor)
}
