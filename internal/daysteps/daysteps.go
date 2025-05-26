package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные о дневных прогулках
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку формата "678,0h50m"
func (ds *DaySteps) Parse(datastring string) error {
    parts := strings.Split(datastring, ",")
    if len(parts) != 2 {
        return fmt.Errorf("неверный формат данных: должно быть 2 части")
    }
    stepsStr := parts[0]
    if strings.TrimSpace(stepsStr) != stepsStr {
        return fmt.Errorf("шаги не должны содержать пробелы в начале или конце")
    }
    steps, err := strconv.Atoi(stepsStr)
    if err != nil {
        return fmt.Errorf("ошибка при парсинге шагов: %w", err)
    }
    if steps <= 0 {
        return fmt.Errorf("количество шагов должно быть положительным")
    }
    ds.Steps = steps

    duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
    if err != nil {
        return fmt.Errorf("ошибка при парсинге длительности: %w", err)
    }
    if duration <= 0 {
        return fmt.Errorf("длительность должна быть положительной")
    }
    ds.Duration = duration

    return nil
}


// ActionInfo формирует строку с информацией о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка при расчёте калорий: %w", err)
	}
	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)
	return info, nil
}
