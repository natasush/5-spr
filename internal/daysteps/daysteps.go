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
		return errors.New("Parse: wrong data")
	}
	ds.Steps, err = strconv.Atoi(vals[0]) //сохр. полученное значение шагов в поле структуры DaySteps
	if err != nil || ds.Steps <= 0 {      //проверяем корректность преобразования данных о шагах:из строки в int (и положительность знач.)
		return errors.New("Parse: steps")
	}

	dur, err := time.ParseDuration(vals[1]) //сохр. продолжительности в поле структуры DaySteps
	if err != nil || dur <= 0 {             // проверка корректности преобразования данных о времени: из строки в time.Duration (и положительность знач.)
		return errors.New("Parse: duration")
	}
	ds.Duration = dur
	return nil //если не было ошибок

}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)                                          //дистанция
	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration) //количество сожженных калорий
	if err != nil {                                                                            //в случае ошибки при расчете калорий
		return "", errors.New("ActionInfo: WalkingSpentCalories")
	}
	s := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, ccal)
	return s, nil
}
