package main

import (
	"fmt"

	"main.go/MVC/gpax/controller"
	"main.go/MVC/gpax/model"
)

func main() {
	stu := []model.Student{
		{
			Id:    "65395",
			Name:  "Mook",
			Grade: []string{"A", "B", "C", "A"},
		},
		{
			Id:    "65425",
			Name:  "Fin",
			Grade: []string{"B", "A", "C", "B"},
		},
		{
			Id:    "65397",
			Name:  "Khim",
			Grade: []string{"A", "C", "C", "B"},
		},
	}
	m := model.Major{
		Majorname: "Information Technology",
		Students:  stu,
	}

	fmt.Println(m.Majorname)
	fmt.Println("----------------------------")
	for _, s := range m.Students {
		fmt.Printf("%s %s\n", s.Id, s.Name)
		fmt.Printf("Have GPAX: %.2f\n", controller.CalGPAX(s.Grade))
	}
}
