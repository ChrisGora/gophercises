package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type questionT struct {
	q, a string
}

type score int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func questions() []questionT {
	f, err := os.Open("questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []questionT
	for _, row := range table {
		questions = append(questions, questionT{q: row[0], a: row[1]})
	}
	return questions
}

func ask(s score, question questionT) score {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		s++
	} else {
		fmt.Println("Incorrect :-(")
	}
	return s
}

func main() {
	s := score(0)
	for _, q := range questions() {
		s = ask(s, q)
	}
	fmt.Println("Final score", s)
}
