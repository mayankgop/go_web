package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type QA struct {
	Qus string
	Ans string
}

func myparse(x [][]string) []QA {
	res := make([]QA, len(x))
	for i, v := range x {
		res[i] = QA{
			v[0],
			v[1],
		}
	}
	return res
}

func main() {
	x := flag.String("f", "problems.csv", "this is my input file")
	// timeLimit := flag.Int("limit", 20, "the time limit for the quiz in seconds")
	flag.Parse()
	data, _ := os.Open(*x)
	record := csv.NewReader(data)
	finaldata, _ := record.ReadAll()
	a := myparse(finaldata)
	var c int
	var answ string
	timer := time.NewTimer(10*time.Second)
	answerCh := make(chan string)

	for i, v := range a {
		fmt.Printf("question no. %d is %s \n", i+1, v.Qus)
			go func() {
				fmt.Scanf("%s", &answ)
				answerCh <- answ

			}()
		select {
			case answer := <-answerCh:
				if answer == v.Ans {
					c++
				}

			case <-timer.C:
				fmt.Printf("the score is %d out of %d ", c, len(finaldata))
				return

		}

	}

}
