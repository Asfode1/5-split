package personaldata

import "fmt"

// Personal — структура с данными пользователя
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные структуры Personal на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n\n", p.Height)
}

