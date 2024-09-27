package main

import "fmt"

// Функция, которая принимает два целых числа и возвращает их сумму
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 3)
    fmt.Println("Sum:", sum) // Вывод: Sum: 8
}
