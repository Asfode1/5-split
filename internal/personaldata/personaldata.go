package main

import "fmt"

// Personal — экспортируемая структура с данными пользователя
type Personal struct {
    Name   string  // Имя пользователя
    Weight float64 // Вес пользователя
    Height float64 // Рост пользователя
}

// Print выводит данные структуры Personal на экран
func (p Personal) Print() {
    fmt.Printf("Имя: %s\n", p.Name)
    fmt.Printf("Вес: %.2f\n", p.Weight)
    fmt.Printf("Рост: %.2f\n", p.Height)
}

func main() {
    user := Personal{
        Name:   "Иван",
        Weight: 75.5,
        Height: 1.82,
    }
    user.Print()
}
