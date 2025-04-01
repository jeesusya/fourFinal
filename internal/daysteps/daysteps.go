package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65  // длина шага в метрах
	Kilometer  = 1000. // количество метров в километре
)

func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, errors.New("invalid input")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	dist := (float64(steps) * StepLength) / Kilometer

	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	//Количество шагов: 792.
	//Дистанция составила 0.51 км.
	//	Вы сожгли 221.33 ккал.
	result := fmt.Sprintf("Количество шагов : %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, dist, calories)
	return result
}
