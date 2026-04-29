package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, fmt.Errorf("steps less then zero")
	case weight <= 0:
		return 0, fmt.Errorf("weight less then zero")
	case height <= 0:
		return 0, fmt.Errorf("height less then zero")
	case duration <= 0:
		return 0, fmt.Errorf("duration less then zero")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	calories := weight * averageSpeed * duration.Minutes() / float64(minInH)

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, fmt.Errorf("steps less then zero")
	case weight <= 0:
		return 0, fmt.Errorf("weight less then zero")
	case height <= 0:
		return 0, fmt.Errorf("height less then zero")
	case duration <= 0:
		return 0, fmt.Errorf("duration less then zero")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	calories := weight * averageSpeed * duration.Minutes() / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	distance := float64(steps) * stepLength / mInKm
	return distance
}
