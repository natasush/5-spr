package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
	//Info
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v) // Используем метод Parse интерфейса
		if err != nil {
			log.Println("error: ", err)
			continue
		}
		info, err := dp.ActionInfo() // Используем метод ActionInfo интерфейса
		if err != nil {
			log.Println("error ", err)
			continue
		}
		fmt.Println(info)
	}
}
