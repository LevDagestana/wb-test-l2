package patterns

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

	Паттерн Строитель - это паттерн проектирования, который позволяет создавать сложные объекты пошагово.

	Плюсы:
		+ позволяет использовать один и тот же код для получения разных представлений
		+ позволяет создавать стандартные значения полей объекта и избавляться от мусорных аргументов
			со стандартными/пустыми значениями
		+ позволяет инициализировать поля объекта в зависимости от входных параметров

	Минусы:
		- увеличивает сложность кода и количество введенных функций (строителей конкретного объекта внутри конструируемого)
		- при конкурентном программировании необходимо мьютить конструируемый объект, поскольку он конструируется пошагово
*/

// Product - Продукт, который мы строим
type Product struct {
	part1 string
	part2 int
}

func (p *Product) String() string {
	return fmt.Sprintf("Part1: %s, Part2: %d", p.part1, p.part2)
}

// Builder - Интерфейс Строителя
type Builder interface {
	buildPart1()
	buildPart2()
	getProduct() *Product
}

// ConcreteBuilderA - Конкретный строитель для создания ProductA
type ConcreteBuilderA struct {
	product *Product
}

// NewConcreteBuilderA - Конструктор для A
func NewConcreteBuilderA() *ConcreteBuilderA {
	return &ConcreteBuilderA{
		product: &Product{},
	}
}

func (b *ConcreteBuilderA) buildPart1() {
	b.product.part1 = "PartA"
}

func (b *ConcreteBuilderA) buildPart2() {
	b.product.part2 = 1
}

func (b *ConcreteBuilderA) getProduct() *Product {
	return b.product
}

// ConcreteBuilderB - Конкретный строитель для создания ProductB
type ConcreteBuilderB struct {
	product *Product
}

// NewConcreteBuilderB - Конструктор для B
func NewConcreteBuilderB() *ConcreteBuilderB {
	return &ConcreteBuilderB{
		product: &Product{},
	}
}

func (b *ConcreteBuilderB) buildPart1() {
	b.product.part1 = "PartB"
}

func (b *ConcreteBuilderB) buildPart2() {
	b.product.part2 = 2
}

func (b *ConcreteBuilderB) getProduct() *Product {
	return b.product
}

// Director - Директор, который управляет строителями
type Director struct {
	builder Builder
}

// NewDirector - Конструктор для директора
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct - Конструктор для продукта
func (d *Director) construct() *Product {
	d.builder.buildPart1()
	d.builder.buildPart2()
	return d.builder.getProduct()
}

// BuilderTest - Тест строителя
func BuilderTest() {
	// Создание продукта A
	builderA := NewConcreteBuilderA()
	directorA := NewDirector(builderA)
	productA := directorA.construct()
	fmt.Println("Product A:", productA)

	// Создание продукта B
	builderB := NewConcreteBuilderB()
	directorB := NewDirector(builderB)
	productB := directorB.construct()
	fmt.Println("Product B:", productB)
}
