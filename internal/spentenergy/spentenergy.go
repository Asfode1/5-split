package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance рассчитывает дистанцию в километрах по количеству шагов и росту пользователя.
func Distance(steps int, height float64) float64 {
	// Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient

	// Вычисляем расстояние в километрах
	return (float64(steps) * stepLength) / mInKm
}

// MeanSpeed рассчитывает среднюю скорость (км/ч) по шагам, росту и длительности.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка на отрицательные шаги или нулевую продолжительность
	if steps < 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	hours := duration.Hours()

	// Избегаем деления на ноль
	if hours == 0 {
		return 0
	}

	return distance / hours
}

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
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

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed == 0 {
		return 0, errors.New("средняя скорость равна нулю, проверьте входные данные")
	}

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
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

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed == 0 {
		return 0, errors.New("средняя скорость равна нулю, проверьте входные данные")
	}

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}
