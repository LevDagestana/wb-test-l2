package patterns

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Паттерн Стратегия позволяет определить семейство алгоритмов, инкапсулировать каждый из них и делать их взаимозаменяемыми.
		Это позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются.

	Плюсы:
		+ Можно легко добавлять новые алгоритмы или изменять существующие без изменения структуры Бизнес-логики
		+ Каждый алгоритм инкапсулирован в собственном классе, это способствует изоляции и уменьшению сложности кода

	Минусы:
		- Введение дополнительных классов для каждого алгоритма может привести к увеличению сложности структуры программы
		- В некоторых случаях каждый алгоритм требует своей собственной реализации в отдельном классе, это может привести к разрастанию кода

	Пример ниже.
*/

// Strategy - Интерфейс стратегии
type Strategy interface {
	DoOperation(int, int) int
}

// ConcreteStrategyAdd - Конкретная стратегия сложения
type ConcreteStrategyAdd struct{}

// DoOperation - Метод выполнения
func (s *ConcreteStrategyAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// ConcreteStrategySubtract - Конкретная стратегия вычитания
type ConcreteStrategySubtract struct{}

// DoOperation - Метод выполнения
func (s *ConcreteStrategySubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// Context - Контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

// SetStrategy - Метод установки стратегии
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy - Метод выполнения
func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

// StrategyTest - Тест стратегии
func StrategyTest() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем стратегию сложения и выполняем операцию
	context.SetStrategy(&ConcreteStrategyAdd{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	// Устанавливаем стратегию вычитания и выполняем операцию
	context.SetStrategy(&ConcreteStrategySubtract{})
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}
