package main

import "fmt"

// Определим структуру
type Rectangle struct {
    width  int
    height int
}

// Метод для структуры Rectangle, который вычисляет площадь
func (r Rectangle) Area() int {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 10, height: 57}
    fmt.Println("Area:", rect.Area()) // Вывод: Area: 570
}
