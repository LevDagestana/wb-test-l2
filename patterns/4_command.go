package patterns

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Паттерн Комманда позволяет параметризовать клиентов с различными запросами, очередями или журналами запросов
		и поддерживать отмену операций

	Плюсы:
		+ Команда инкапсулирует запрос в виде объекта, что позволяет отделить отправителя запроса от его исполнителя
		+ Команды могут поддерживать операцию отмены, что позволяет отменять и повторять операции
		+ Благодаря использованию объектов команд, можно параметризовать и настраивать поведение программы динамически

	Минусы:
		- Усложнение структуры классов
		- Используется больше памяти, а также на создание объектов требуется лишнее время

	Пример ниже.
*/

// Command - Интерфейс Команды
type Command interface {
	Execute()
}

// ConcreteCommand - Конкретная реализация Команды
type ConcreteCommand struct {
	receiver *Receiver
}

// Execute - Выполнение команды
func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Receiver - Получатель команды
type Receiver struct{}

// Action - Выполнение действия
func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// Invoker - Отправитель команды
type Invoker struct {
	command Command
}

// SetCommand - Установка команды
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand - Выполнение команды
func (i *Invoker) ExecuteCommand() {
	fmt.Println("Invoker: выполнение команды")
	i.command.Execute()
}

// CommandTest - Тест команд
func CommandTest() {
	// Создаем получателя
	receiver := &Receiver{}

	// Создаем команду и связываем ее с получателем
	command := &ConcreteCommand{receiver}

	// Создаем отправителя и устанавливаем команду
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// Отправляем команду
	invoker.ExecuteCommand()
}
