package actioninfo

import (
    "fmt"
    "log"
)

// DataParser интерфейс для парсинга данных и получения информации
type DataParser interface {
    Parse(datastring string) error
    ActionInfo() (string, error)
}

// Info обрабатывает набор данных и выводит информацию о активностях
func Info(dataset []string, dp DataParser) {
    for _, data := range dataset {
        // Парсим данные
        if err := dp.Parse(data); err != nil {
            log.Printf("Ошибка парсинга данных: %v", err)
            continue
        }
        
        // Получаем информацию
        info, err := dp.ActionInfo()
        if err != nil {
            log.Printf("Ошибка получения информации: %v", err)
            continue
        }
        
        // Выводим результат
        fmt.Println(info)
		fmt.Println("--------------------------------------------------")
    }
}