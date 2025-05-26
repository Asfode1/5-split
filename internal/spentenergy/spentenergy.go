package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance рассчитывает дистанцию в километрах по количеству шагов и росту пользователя.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

// MeanSpeed рассчитывает среднюю скорость (км/ч) по шагам, росту и длительности.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distance / hours
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

