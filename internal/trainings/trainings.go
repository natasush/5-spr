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
		return errors.New("wrong data")
	}
	steps, err := strconv.Atoi(vals[0]) //получение количества шагов из строки
	if err != nil {                     //проверка корректности преобразования шагов из строки в int
		err = errors.New("parse: steps")
		return err
	}
	if steps <= 0 {
		return errors.New("steps<=0")
	}
	t.Steps = steps //сохр. полученное значение шагов в поле структуры Trainig

	t.TrainingType = vals[1] //сохр. значение типа тренировки в поле структуры Training

	duration, err := time.ParseDuration(vals[2]) //получение продолжительности из строки
	if err != nil {                              // проверка корректности преобразования времени из строки в time.Duration
		err = errors.New("parse: duration")
		return err
	}
	if duration <= 0 {
		return errors.New("duration<=0")
	}
	t.Duration = duration // сохр. значение продолжительности в поле структуры Training
	return nil            //если нет ошибок
}
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)                //дистанция
	meanSp := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) //средняя скорость
	duration := t.Duration.Hours()                                 //продолжительность тренировки в часах

	caloriesRun, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil { // в случае ошибки при расчете калорий
		err = errors.New("run calories calculation error")
		return "", err
	}
	caloriesWalk, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil { // в случае ошибки при расчете калорий
		err = errors.New("walc calories calculation error")
		return "", err
	}
	switch t.TrainingType { //вывод на экран в зависимости от типа тренировки
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, duration, dist, meanSp, caloriesRun), nil
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, duration, dist, meanSp, caloriesWalk), nil
	default: // если неизвестный тип тренировки
		return "", errors.New("неизвестный тип тренировки")
	}

}
