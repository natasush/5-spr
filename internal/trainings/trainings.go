package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 3 { //проверка количества данных
		return errors.New("Parse: wrong data")
	}
	steps, err := strconv.Atoi(vals[0]) //получение количества шагов из строки
	if err != nil || steps <= 0 {       //проверка корректности преобразования шагов из строки в int (и положит. знач.)
		return errors.New("Parse: steps")
	}
	t.Steps = steps //сохр. полученное значение шагов в поле структуры Trainig

	t.TrainingType = vals[1] //сохр. значение типа тренировки в поле структуры Training

	duration, err := time.ParseDuration(vals[2]) //получение продолжительности из строки
	if err != nil || duration <= 0 {             // проверка корректности преобразования времени из строки в time.Duration  (и положит. знач.)
		return errors.New("Parse: duration")
	}
	t.Duration = duration // сохр. значение продолжительности в поле структуры Training
	return nil            //если нет ошибок
}
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)                //дистанция
	meanSp := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) //средняя скорость
	duration := t.Duration.Hours()                                 //продолжительность тренировки в часах

	switch t.TrainingType { //вычисление количества сожженных калорий в зависимости от типа тренировки
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil { // в случае ошибки при расчете калорий
			return "", errors.New("ActionInfo:RunningSpentCalories")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, duration, dist, meanSp, calories), nil
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil { // в случае ошибки при расчете калорий
			return "", errors.New("ActionInfo:WalkingSpentCalories")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, duration, dist, meanSp, calories), nil
	default: // если неихвестный тип тренировки
		return "", errors.New("неизвестный тип тренировки")
	}

}
