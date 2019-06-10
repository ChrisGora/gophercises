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

func ask(question questionT, correct chan bool) {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		fmt.Println("Correct!")
		correct <- true
	} else {
		fmt.Println("Incorrect :-(")
		correct <- false
	}
}

//func timer() {
//
//}

func quiz(exit chan bool) {
	s := score(0)
	qs := questions()
	correct := make(chan bool)
	i := 0

	go ask(qs[i], correct)

	for {
		select {
		case c := <-correct:
			if c {
				s++
			}
			i++
			if i < len(qs) {
				go ask(qs[i], correct)
			} else {
				fmt.Println("End of questions, Final score", s)
				exit <- true
			}
		}
	}
}

func main() {
	exit := make(chan bool)
	go quiz(exit)
	<-exit
}
