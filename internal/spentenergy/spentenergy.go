package spentenergy

import (
	"time"
    "errors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    // Проверка входных параметров
    if steps <= 0 {
        return 0, errors.New("количество шагов должно быть положительным")
    }
    if weight <= 0 {
        return 0, errors.New("вес должен быть положительным")
    }
    if height <= 0 {
        return 0, errors.New("рост должен быть положительным")
    }
    if duration <= 0 {
        return 0, errors.New("продолжительность должна быть положительной")
    }

    // Рассчитываем среднюю скорость (км/ч)
    meanSpeed := MeanSpeed(steps, height, duration)
    if meanSpeed == 0 {
        return 0, errors.New("средняя скорость равна нулю, проверьте входные данные")
    }

    // Продолжительность в минутах
    durationInMinutes := duration.Minutes()

    // Расчёт калорий с коэффициентом для ходьбы
    calories := (weight * meanSpeed * durationInMinutes) / minInH
    calories *= walkingCaloriesCoefficient

    return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    // Проверка входных параметров
    if steps <= 0 {
        return 0, errors.New("количество шагов должно быть положительным")
    }
    if weight <= 0 {
        return 0, errors.New("вес должен быть положительным")
    }
    if height <= 0 {
        return 0, errors.New("рост должен быть положительным")
    }
    if duration <= 0 {
        return 0, errors.New("продолжительность должна быть положительной")
    }

    // Рассчитываем среднюю скорость (км/ч)
    meanSpeed := MeanSpeed(steps, height, duration)
    if meanSpeed == 0 {
        return 0, errors.New("средняя скорость равна нулю, проверьте входные данные")
    }

    // Продолжительность в минутах
    durationInMinutes := duration.Minutes()

    // Расчёт калорий
    calories := (weight * meanSpeed * durationInMinutes) / minInH

    return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
    // Проверка на отрицательные шаги или нулевую продолжительность
    if steps < 0 || duration <= 0 {
        return 0
    }
    
    // Рассчитываем дистанцию
    distance := Distance(steps, height)
    
    // Переводим продолжительность в часы
    hours := duration.Hours()
    
    // Рассчитываем среднюю скорость
    return distance / hours
}

func Distance(steps int, height float64) float64 {
		// Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient
		
		// Вычисляем общее расстояние в метрах и конвертируем в километры
	return (float64(steps) * stepLength) / mInKm
	}

