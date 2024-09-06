package main

import (
	"fmt"

	"main.go/MVC/bmi/controller"
	"main.go/MVC/bmi/model"
)

func main() {

	p1 := model.Person{
		Name:   "Fook",
		Age:    20,
		Weight: 40.25,
		Height: 1.75,
	}

	bmi := controller.CalBMI(p1.Weight, p1.Height)
	fmt.Println("Name:", p1.Name, "\nAge:", p1.Age)
	fmt.Println("Weight:", p1.Weight, "\nHeight:", p1.Height)
	fmt.Printf("BMI: %.2f\n", bmi)
	controller.CompareBMI(bmi)

}
