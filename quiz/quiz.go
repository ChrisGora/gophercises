package main

import (
	"encoding/csv"
	"os"
)

type question struct {
	q, a string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func questions() []question {
	f, err := os.Open("questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _,row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

func main() {
	//fmt.Println(questions())
}
