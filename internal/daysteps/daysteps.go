package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("incorret number of values")
	}

	steps, err := strconv.Atoi(data[0])

	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("steps less then zero")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])

	if err != nil {
		return err
	}

	if duration <= 0 {
		return fmt.Errorf("time less then zero")
	}

	ds.Duration = duration
	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories), nil
}
