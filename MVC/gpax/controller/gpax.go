package controller

func CalGPAX(Grade []string) float64 {
	GradePoint := 0
	sum := 0.0
	for i := 0; i < len(Grade); i++ {
		if Grade[i] == "A" {
			GradePoint = 4
		} else if Grade[i] == "B" {
			GradePoint = 3
		} else if Grade[i] == "C" {
			GradePoint = 2
		} else if Grade[i] == "D" {
			GradePoint = 1
		} else if Grade[i] == "F" {
			GradePoint = 0
		} else {
			return -1
		}
		sum += float64(GradePoint)
	}
	gpax := sum / float64(len(Grade))
	return gpax
}
