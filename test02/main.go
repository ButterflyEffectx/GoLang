package main

import "fmt"

type student struct {
	id    string
	name  string
	grade []string
}

type major struct {
	majorname string
	students  []student
}

func calGPAX(grade []string) float64 {
	gradePoint := 0
	sum := 0.0
	for i := 0; i < len(grade); i++ {
		if grade[i] == "A" {
			gradePoint = 4
		} else if grade[i] == "B" {
			gradePoint = 3
		} else if grade[i] == "C" {
			gradePoint = 2
		} else if grade[i] == "D" {
			gradePoint = 1
		} else if grade[i] == "F" {
			gradePoint = 0
		} else {
			return -1
		}
		sum += float64(gradePoint)
	}
	gpax := sum / float64(len(grade))
	return gpax
}

func main() {
	stu := []student{
		{
			id:    "65395",
			name:  "Mook",
			grade: []string{"A", "B", "C", "A"},
		},
		{
			id:    "65425",
			name:  "Fin",
			grade: []string{"B", "A", "C", "B"},
		},
		{
			id:    "65397",
			name:  "Khim",
			grade: []string{"A", "C", "C", "B"},
		},
	}
	m := major{
		majorname: "Information Technology",
		students:  stu,
	}

	fmt.Println(m.majorname)
	fmt.Println("----------------------------")
	for _, s := range m.students {
		fmt.Printf("%s %s\n", s.id, s.name)
		fmt.Printf("Have GPAX: %.2f\n", calGPAX(s.grade))
	}
}
