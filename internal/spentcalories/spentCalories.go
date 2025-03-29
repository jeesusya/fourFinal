package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, errors.New("invalid input")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}
	activity := slice[1]
	return steps, activity, duration, nil

}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return float64(steps) * lenStep / minInH
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps)
	return float64(dist) / float64(duration.Hours())
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.

// Тип тренировки: Бег
// Длительность: 0.75 ч.
// Дистанция: 10.00 км.
// Скорость: 13.34 км/ч
// Сожгли калорий: 18621.75
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		fmt.Println(err)
	}
	dist := distance(steps)
	speed := meanSpeed(steps, duration)

	switch activity {
	case "Бег":
		calories := RunningSpentCalories(steps, weight, duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %2.f км.\nСкорость: %2.f км/ч\nСожгли калорий: %2.f",
			activity, duration.Hours(), dist, speed, calories)
	case "Ходьба":
		calories := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %2.f км.\nСкорость: %2.f км/ч\nСожгли калорий: %2.f",
			activity, duration.Hours(), dist, speed, calories)
	default:
		return fmt.Sprint("неизвестный тип тренировки")
	}

}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	speed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	speed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
