package main

import "fmt"

// Интерфейс Shape с одним методом Area
type Shape interface {
    Area() int
}

// Структура Rectangle
type Rectangle struct {
    width  int
    height int
}

// Метод Area для Rectangle
func (r Rectangle) Area() int {
    return r.width * r.height
}

// Структура Circle
type Circle struct {
    radius int
}

// Метод Area для Circle
func (c Circle) Area() int {
    return 3 * c.radius * c.radius // Упрощенная формула для примера
}

func main() {
    // Используем интерфейс для различных типов
    var s Shape

    rect := Rectangle{width: 10, height: 5}
    s = rect
    fmt.Println("Rectangle Area:", s.Area()) // Вывод: Rectangle Area: 50

    circle := Circle{radius: 5}
    s = circle
    fmt.Println("Circle Area:", s.Area()) // Вывод: Circle Area: 75
}
