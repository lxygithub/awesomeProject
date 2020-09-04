package models

import "fmt"

type Country struct {
	Name string
}

func (c Country) GetName() {
	fmt.Print(c.Name)
}
