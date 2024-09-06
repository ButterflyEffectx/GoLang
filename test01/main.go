package main

import "fmt"

type person struct {
	name   string
	age    int
	weight float64
	height float64
}

func calBMI(w float64, h float64) float64 {
	return w / (h * h)
}

func compareBMI(bmi float64) {
	//เปรียบเทียบเกณฑ์bmi
	if bmi > 30 {
		fmt.Println("Obesity")
	} else if bmi > 25 {
		fmt.Println("Overweight")
	} else if bmi > 18.5 {
		fmt.Println("Normal weight")
	} else {
		fmt.Println("Underweight")
	}
}

func main() {
	p1 := person{
		name:   "Mook",
		age:    20,
		weight: 40.25,
		height: 1.75, //รับเป็นเมตร
	}
	bmi := calBMI(p1.weight, p1.height)
	fmt.Println("Name:", p1.name, "age:", p1.age)
	fmt.Println("Weight:", p1.weight, "Height:", p1.height)
	fmt.Printf("BMI: %.2f\n", calBMI(p1.weight, p1.height))
	compareBMI(bmi)

}
