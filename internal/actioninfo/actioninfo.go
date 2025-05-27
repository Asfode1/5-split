package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

// DataParser интерфейс для парсинга данных и получения информации
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных и выводит информацию о активностях
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// 1. Очищаем входные данные
		cleanData := strings.TrimSpace(data)

		if err := dp.Parse(cleanData); err != nil {
			log.Printf("Ошибка парсинга данных: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации: %v", err)
			continue
		}

		// 2. Очищаем вывод от лишних переводов строк в конце
		info = strings.TrimSpace(info)

		// 3. Добавляем перевод строки и разделитель
		fmt.Printf("%s\n", info)
	}
}



