package model

type Student struct {
	Id    string
	Name  string
	Grade []string
}

type Major struct {
	Majorname string
	Students  []Student
}
