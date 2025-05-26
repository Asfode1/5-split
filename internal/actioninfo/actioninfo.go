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
        if err := dp.Parse(data); err != nil {
            log.Printf("Ошибка парсинга данных: %v", err)
            continue
        }

        info, err := dp.ActionInfo()
        if err != nil {
            log.Printf("Ошибка получения информации: %v", err)
            continue
        }

        // Если info не заканчивается на \n, добавим его
        if len(info) == 0 || info[len(info)-1] != '\n' {
            info += "\n"
        }

        fmt.Print(info)
        fmt.Printf("--------------------------------------------------")
    }
}
