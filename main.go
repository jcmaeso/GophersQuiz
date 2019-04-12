package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Problems []Problem
}

//MakeQuestion function to ask a question to the user and evaluate response
func (p *Problem) MakeQuestion() bool {
	fmt.Printf("%s >", p.Question)
	var userAnswer string
	fmt.Scanln(&userAnswer)
	if strings.ToLower(p.Answer) == strings.ToLower(userAnswer) {
		return true
	}
	return false
}

//Play function to launch some quizzes
func (q *Quiz) play() {
	var points int = 0
	for _, p := range q.Problems {
		if p.MakeQuestion() {
			points++
		}
	}
	fmt.Println("Your score is ", points)
}

//LoadQuiz Function to read csv file and dump it into a structure
func LoadQuiz(filename string) (*Quiz, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error opening file")
	}
	defer f.Close()

	//ReadLines to variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, errors.New("Error reading CSV")
	}
	if len(lines) != 2 {
		return nil, errors.New("Invalid CSV format")
	}

	quiz := Quiz{Problems: make([]Problem, 0, len(lines))}
	for _, line := range lines {
		quiz.Problems = append(quiz.Problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return &quiz, nil
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "problems.csv", "csv file to be read")
	flag.Parse()
	quiz, err := LoadQuiz(filename)
	if err != nil {
		panic(err)
	}
	quiz.play()
}
