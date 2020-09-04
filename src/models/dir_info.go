package models

import "fmt"

type Car struct {
	Name string
}

func (car *Car) Run() {
	fmt.Println(car.Name + " running........")
}
