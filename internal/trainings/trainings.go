package trainings

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/Yandex-Practicum/tracker/internal/personaldata" // Импортируем пакет с Personal
    "github.com/Yandex-Practicum/tracker/internal/spentenergy"  // Импортируем пакет с расчетами
)

// Training содержит данные о тренировке
type Training struct {
    Steps        int
    TrainingType string
    Duration     time.Duration
    personaldata.Personal
}

// Parse парсит строку формата "3456,Ходьба,3h00m" и записывает данные в структуру Training
func (t *Training) Parse(datastring string) error {
    parts := strings.Split(datastring, ",")
    if len(parts) != 3 {
        return errors.New("неверный формат данных: должно быть 3 части")
    }

    // Парсим количество шагов
    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return fmt.Errorf("ошибка при парсинге шагов: %w", err)
    }
    t.Steps = steps

    // Сохраняем тип тренировки
    t.TrainingType = parts[1]

    // Парсим длительность
    duration, err := time.ParseDuration(parts[2])
    if err != nil {
        return fmt.Errorf("ошибка при парсинге длительности: %w", err)
    }
    t.Duration = duration

    return nil
}

// ActionInfo формирует строку с информацией о тренировке
func (t Training) ActionInfo() (string, error) {
    // Вычисляем дистанцию
    distance := spentenergy.Distance(t.Steps, t.Height)

    // Вычисляем среднюю скорость
    meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

    var calories float64
    switch strings.ToLower(t.TrainingType) {
    case "бег":
        c, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
        if err != nil {
            return "", fmt.Errorf("ошибка при расчёте калорий для бега: %w", err)
        }
        calories = c
    case "ходьба":
        c, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
        if err != nil {
            return "", fmt.Errorf("ошибка при расчёте калорий для ходьбы: %w", err)
        }
        calories = c
    default:
        return "", errors.New("неизвестный тип тренировки")
    }

    durationHours := t.Duration.Hours()

    info := fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
        t.TrainingType,
        durationHours,
        distance,
        meanSpeed,
        calories,
    )

    return info, nil
}
