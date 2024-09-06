package controller

import "fmt"

func CalBMI(w float64, h float64) float64 {
	return w / (h * h)
}

func CompareBMI(bmi float64) {
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
