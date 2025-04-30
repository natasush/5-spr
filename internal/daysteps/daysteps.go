package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 2 { //проверка количества данных
		return errors.New("wrong data")
	}
	steps, err := strconv.Atoi(vals[0]) //извлечение данных о шагах из строки
	if err != nil {                     //проверяем корректность преобразования данных о шагах:из строки в int
		err = errors.New("parse: steps")
		return err
	}
	if steps <= 0 {
		return errors.New("steps<=0")
	}
	dur, err := time.ParseDuration(vals[1]) //извлечение данных о продолжительности из строки
	if err != nil {                         // проверка корректности преобразования данных о времени: из строки в time.Duration
		err = errors.New("parse: duration")
		return err
	}
	if dur <= 0 {
		return errors.New("duration<=0")
	}
	ds.Steps = steps  //сохр. шагов в поле структуры DaySteps
	ds.Duration = dur //сохр. продолжительности в поле структуры DaySteps
	return nil        //если не было ошибок

}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)                                          //дистанция
	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration) //количество сожженных калорий
	if err != nil {
		err = errors.New("calories calculation error") //в случае ошибки при расчете калорий
		return "", err
	}
	s := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, ccal)
	return s, nil
}
