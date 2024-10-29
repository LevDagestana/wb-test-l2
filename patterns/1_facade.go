package patterns

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Паттерн Фасад позволяет делегировать массовое выполнение запросов другим объектам.

	Плюсы:
		+ позволяет легко расширить существующую бизнес-логику или мигрировать на другую
		+ позволяет свести обработку разных типов к одной функции, которая затем свитчится по дженерикам (если это функциональный фасад)
		+ позволяет делегировать массовое выполнение запросов другим функциям, что улучшает читаемость кода и уменьшает вероятность ошибок
		+ позволяет организовать процесс разработки небольшими командами, каждая из которых может не углубляться в то, что находится под чужим Фасадом

	Минусы:
		- сводит работу программы к "бутылочному горлышку", которое становится критично уязвимым местом
		- если модули системы работают на конкретных процессорах, нагрузка на те, что обрабатывают Фасад может быть значительно выше
*/

// SubSystemA - Сложный класс A
type SubSystemA struct{}

// OperationA - Операция A
func (s *SubSystemA) OperationA() {
	fmt.Println("SubSystemA: OperationA")
}

// SubSystemB - Сложный класс B
type SubSystemB struct{}

// OperationB - Операция B
func (s *SubSystemB) OperationB() {
	fmt.Println("SubSystemB: OperationB")
}

// Facade - Фасад над A и B
type Facade struct {
	subSystemA *SubSystemA
	subSystemB *SubSystemB
}

// NewFacade - Фасад над A и B
func NewFacade() *Facade {
	return &Facade{
		subSystemA: &SubSystemA{},
		subSystemB: &SubSystemB{},
	}
}

// Operation - Операция с фасадом
func (f *Facade) Operation() {
	fmt.Println("Facade: Операция запущена")
	f.subSystemA.OperationA()
	f.subSystemB.OperationB()
	fmt.Println("Facade: Операция завершена")
}

// FacadeTest - Тест фасада
func FacadeTest() {
	facade := NewFacade()
	facade.Operation()
}
