package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type questionT struct {
	q, a string
}

type scoreT int

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

func quiz(exit chan bool) {
	qs := questions()
	score := scoreT(0)
	correct := make(chan bool)
	current := 0

	go ask(qs[current], correct)
	timeOver := time.After(5 * time.Second)

	for {
		select {

		case c := <-correct:
			if c {
				score++
			}
			current++
			if current < len(qs) {
				go ask(qs[current], correct)
			} else {
				fmt.Println("End of questions, final score", score)
				exit <- true
			}

		case <-timeOver:
			fmt.Println("Timed out, final score", score)
			exit <- true
		}
	}
}

func main() {
	exit := make(chan bool)
	go quiz(exit)
	<-exit
}
