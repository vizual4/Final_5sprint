package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	prs := strings.Split(datastring, ",")
	if len(prs) != 3 {
		return fmt.Errorf("incorrect number of values")
	}

	steps, err := strconv.Atoi(prs[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("steps less then zero")
	}

	t.Steps = steps

	t.TrainingType = prs[1]

	duration, err := time.ParseDuration(prs[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("time of training less then zero")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)

	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		spentCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, averageSpeed, spentCalories), nil

	case "Ходьба":
		spentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, averageSpeed, spentCalories), nil
	}

	return "", fmt.Errorf("неизвестный тип тренировки")

}
