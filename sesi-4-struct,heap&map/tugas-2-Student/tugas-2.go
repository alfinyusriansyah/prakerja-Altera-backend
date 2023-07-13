package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Scores []int
}

func calculateAverageScore(students []Student) int64 {
	totalScore := 0
	for _, student := range students {
		for _, score := range student.Scores {
			totalScore += score
		}
	}
	averageScore := int64(totalScore) / int64(len(students))
	return averageScore
}

func findMinMaxScores(students []Student) (string, int, string, int) {
	minScore := students[0].Scores[0]
	maxScore := students[0].Scores[0]
	minStudent := students[0].Name
	maxStudent := students[0].Name

	for _, student := range students {
		for _, score := range student.Scores {
			if score < minScore {
				minScore = score
				minStudent = student.Name
			}
			if score > maxScore {
				maxScore = score
				maxStudent = student.Name
			}
		}
	}

	return minStudent, minScore, maxStudent, maxScore
}

func main() {
	students := []Student{
		{Name: "rizky", Scores: []int{80}},
		{Name: "andri", Scores: []int{90}},
		{Name: "lola", Scores: []int{70}},
		{Name: "lili", Scores: []int{65}},
		{Name: "lulu", Scores: []int{60}},
	}

	averageScore := calculateAverageScore(students)
	minStudent, minScore, maxStudent, maxScore := findMinMaxScores(students)

	fmt.Println("Average score:", averageScore)
	fmt.Printf("Min score of Students: %s (%d)\n", minStudent, minScore)
	fmt.Printf("Max score of Students: %s (%d)\n", maxStudent, maxScore)
}
