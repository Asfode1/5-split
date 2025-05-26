package daysteps

import (
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/Yandex-Practicum/tracker/internal/personaldata" // импортируем пакет с Personal
    "github.com/Yandex-Practicum/tracker/internal/spentenergy"  // импортируем пакет с расчетами
)

// DaySteps содержит данные о дневных прогулках
type DaySteps struct {
    Steps    int
    Duration time.Duration
    personaldata.Personal
}

// Parse парсит строку формата "678,0h50m" и записывает данные в структуру DaySteps
func (ds *DaySteps) Parse(datastring string) error {
    parts := strings.Split(datastring, ",")
    if len(parts) != 2 {
        return fmt.Errorf("неверный формат данных: должно быть 2 части")
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return fmt.Errorf("ошибка при парсинге шагов: %w", err)
    }
    ds.Steps = steps

    duration, err := time.ParseDuration(parts[1])
    if err != nil {
        return fmt.Errorf("ошибка при парсинге длительности: %w", err)
    }
    ds.Duration = duration

    return nil
}

// ActionInfo формирует строку с информацией о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
    distance := spentenergy.Distance(ds.Steps, ds.Height)

    // Для подсчёта калорий используем WalkingSpentCalories
    calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
    if err != nil {
        return "", fmt.Errorf("ошибка при расчёте калорий: %w", err)
    }

    info := fmt.Sprintf(
        "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
        ds.Steps,
        distance,
        calories,
    )

    return info, nil
}

