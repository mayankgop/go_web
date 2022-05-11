package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type QA struct {
	Qus string
	Ans string
}

func myparse(x [][]string) []QA {
	res := make([]QA, len(x))  //created a slice of type QA here
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
	flag.Parse()
	data, _ := os.Open(*x)
	record := csv.NewReader(data)
	finaldata, _ := record.ReadAll()
	a := myparse(finaldata)
	var c int
	for i, v := range a {
		fmt.Printf("question no. %d is %s \n", i+1, v.Qus)
		var answ string
		fmt.Scanf("%s", &answ)
		// fmt.Printf("%T \n %T", v.Ans, answ)10

		if answ == v.Ans {
			c++
		}
	}
	fmt.Println("score is  ", c)

}
