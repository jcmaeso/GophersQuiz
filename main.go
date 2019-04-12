package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Problems []Problem
}

func (p *Problem) MakeQuestion() bool {
	fmt.Printf("%s >", p.Question)
	var userAnswer string
	fmt.Scanln(&userAnswer)
	if p.Answer == userAnswer {
		return true
	}
	return false
}

func (q *Quiz) play() {
	var points int = 0
	for _, p := range q.Problems {
		if p.MakeQuestion() {
			points++
		}
	}
	fmt.Println("Your score is ", points)
}

func LoadQuiz(filename string) *Quiz {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//ReadLines to variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	quiz := Quiz{Problems: make([]Problem, 0, len(lines))}
	for _, line := range lines {
		quiz.Problems = append(quiz.Problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return &quiz
}

func main() {
	filename := "problems.csv"
	quiz := LoadQuiz(filename)
	quiz.play()
}
