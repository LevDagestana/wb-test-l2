package patterns

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Паттерн Цепочка вызовов используется используется для организации потока обработки запросов,
		когда каждый обработчик в цепочке может принять решение о его обработке или
		передать запрос следующему обработчику в цепочке.

	Плюсы:
		+ Новые обработчики могут быть добавлены или удалены без изменения остальной части кода
		+ Обработчики не зависят друг от друга напрямую, что уменьшает связанность между компонентами системы

	Минусы:
		- Если ни один обработчик не обрабатывает запрос, он может остаться неподдержанным
		- Если цепочка становится слишком длинной или неуправляемой, это может привести к сложностям в понимании
			и поддержке кода, а также к падению скорости работы системы

	Пример ниже.
*/

// Handler - Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request int)
}

// BaseHandler - Базовая структура обработчика
type BaseHandler struct {
	nextHandler Handler
}

// SetNext - Метод установки следующего обработчика
func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// ConcreteHandler - Конкретная реализация обработчика
type ConcreteHandler struct {
	BaseHandler
}

// Handle - Метод обработки запроса
func (h *ConcreteHandler) Handle(request int) {
	if request >= 0 && request < 10 {
		fmt.Printf("%d обработан %T\n", request, h)
	} else if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	} else {
		fmt.Println("Необрабатываемый запрос")
	}
}

// ChainOfRespTest - Тест цепочки вызовов
func ChainOfRespTest() {
	// Создаем цепочку обработчиков
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	// Устанавливаем следующие обработчики
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Отправляем запросы
	requests := []int{5, 15, 25}
	for _, req := range requests {
		handler1.Handle(req)
	}
}
