package main

import "fmt"

// Интерфейс для геометрических фигур
type Shape interface {
    Area() int
}

// Структура Rectangle
type Rectangle struct {
    width, height int
}

// Метод для Rectangle, который вычисляет площадь
func (r Rectangle) Area() int {
    return r.width * r.height
}

// Структура Circle
type Circle struct {
    radius int
}

// Метод для Circle, который вычисляет площадь
func (c Circle) Area() int {
    return 3 * c.radius * c.radius // Упрощённая формула
}

// Функция для вывода площади фигуры
func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    circle := Circle{radius: 5}

    // Используем интерфейс Shape и функцию printArea
    printArea(rect)   // Вызов метода Area для Rectangle через интерфейс
    printArea(circle) // Вызов метода Area для Circle через интерфейс
}
