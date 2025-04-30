package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { //проверка значений на отрицательность
		err := errors.New("negative values are not allowed")
		return 0.0, err
	}
	maenSp := MeanSpeed(steps, height, duration)                              //средняя скорость
	durationInMinutes := duration.Minutes()                                   //продолжительность в минутах
	walkCalories := (weight * maenSp * durationInMinutes) / (float64(minInH)) //количество калорий потраченных при ходьбе
	walkCalories *= walkingCaloriesCoefficient                                //умножение на корректирующий коэффициент
	return walkCalories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { //проверка значений на отрицательность
		err := errors.New("negative values are not allowed")
		return 0.0, err
	}
	maenSp := MeanSpeed(steps, height, duration)                             //средняя скорость
	durationInMinutes := duration.Minutes()                                  //продолжительность в минутах
	runCalories := (weight * maenSp * durationInMinutes) / (float64(minInH)) //количество калрий потраченных при беге
	return runCalories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0.0
	}
	distance := Distance(steps, height)   //дистанция
	meanSp := distance / duration.Hours() //средняя скорость
	return meanSp

}

func Distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient                 //расчет длины шага
	distance := (stepLenght * (float64(steps))) / float64(mInKm) //расчет дистанции
	return distance
}
