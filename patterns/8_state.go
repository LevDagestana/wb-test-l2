package patterns

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Паттерн Состояние позволяет объекту изменять свое поведение в зависимости от своего внутреннего состояния.
		Он применяется, когда объект должен изменять свое поведение в зависимости от внутреннего состояния,
		при этом поведение объекта выглядит как изменение его класса

	Плюсы:
		+ Позволяет вынести логику связанную с каждым состоянием в отдельные классы, что упрощает понимание и поддержку кода
		+ Позволяет добавлять новые состояния и изменять поведение объекта без изменения его класса
		+ Каждое состояние инкапсулируется в отдельном классе, что позволяет избежать сложных условий и зависимостей между состояниями

	Минусы:
		- Большое количество состояний может привести к увеличению количества классов и усложнению структуры программы
		- Каждое новое состояние требует создания отдельного класса, что может привести к разрастанию кода

	Пример ниже.
*/

// State - Интерфейс состояния
type State interface {
	Handle() string
}

// ConcreteStateA - Конкретные реализации состояний
type ConcreteStateA struct{}

// Handle - Метод обработки
func (s *ConcreteStateA) Handle() string {
	return "State A handled"
}

// ConcreteStateB - Конкретные реализации состояний
type ConcreteStateB struct{}

// Handle - Метод обработки
func (s *ConcreteStateB) Handle() string {
	return "State B handled"
}

// StateContext - Контекст, который использует состояние
type StateContext struct {
	state State
}

// SetState - Метод установки состояния
func (c *StateContext) SetState(state State) {
	c.state = state
}

// Request - Метод запроса
func (c *StateContext) Request() string {
	return c.state.Handle()
}

// StateTest - Тест состояний
func StateTest() {
	// Создаем контекст
	StateContext := &StateContext{}

	// Устанавливаем состояние A и выполняем запрос
	StateContext.SetState(&ConcreteStateA{})
	fmt.Println("Request in state A:", StateContext.Request())

	// Устанавливаем состояние B и выполняем запрос
	StateContext.SetState(&ConcreteStateB{})
	fmt.Println("Request in state B:", StateContext.Request())
}
