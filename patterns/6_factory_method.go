package patterns

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Паттерн Фабричный Метод используется для создания объектов без явного указания их классов.
		Вместо этого создание объектов делегируется подклассам, которые решают, какой класс создавать.

	Плюсы:
		+ Новые типы объектов можно легко добавлять, создавая новые подклассы и переопределяя фабричный метод
		+ Бизнес-логика не зависит от конкретных классов создаваемых объектов, это позволяет избежать жестких зависимостей

	Минусы:
		- Если в программе присутствует большое количество типов создаваемых объектов, это может привести к усложнению структуры кода
		- Бизнес-логика теряет контроль над процессом создания объектов, что может привести к нежелательным последствиям в случае
			неправильного использования

	Пример ниже.
*/

// FabricProduct - Интерфейс продукта
type FabricProduct interface {
	Use() string
}

// ConcreteFabricProductA - Конкретная реализация продукта A
type ConcreteFabricProductA struct{}

// Use - Использование продукта A
func (p *ConcreteFabricProductA) Use() string {
	return "FabricProduct A used"
}

// ConcreteFabricProductB - Конкретная реализация продукта B
type ConcreteFabricProductB struct{}

// Use - Использование продукта B
func (p *ConcreteFabricProductB) Use() string {
	return "FabricProduct B used"
}

// Factory - Интерфейс фабрики
type Factory interface {
	CreateFabricProduct() FabricProduct
}

// ConcreteFactoryA - Конкретная реализация фабрики для создания продукта A
type ConcreteFactoryA struct{}

// CreateFabricProduct - Фабричный метод для создания продукта A
func (f *ConcreteFactoryA) CreateFabricProduct() FabricProduct {
	return &ConcreteFabricProductA{}
}

// ConcreteFactoryB - Конкретная реализация фабрики для создания продукта B
type ConcreteFactoryB struct{}

// CreateFabricProduct - Фабричный метод для создания продукта B
func (f *ConcreteFactoryB) CreateFabricProduct() FabricProduct {
	return &ConcreteFabricProductB{}
}

// FactoryMethodTest - Тест фабричного метода
func FactoryMethodTest() {
	// Создание фабрики для продукта A
	factoryA := &ConcreteFactoryA{}
	FabricProductA := factoryA.CreateFabricProduct()
	fmt.Println(FabricProductA.Use())

	// Создание фабрики для продукта B
	factoryB := &ConcreteFactoryB{}
	FabricProductB := factoryB.CreateFabricProduct()
	fmt.Println(FabricProductB.Use())
}
