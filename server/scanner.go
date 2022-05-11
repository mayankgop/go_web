package main

import (
	"bufio"
	"fmt"
	"strings"
	// "text/scanner"
)

func main(){
	s:="njncw wjcjnw xnw cewni nwkjwk knew neoiw\n ue ei hdwk iuwid wdiu\n jed eni dwendi i2indio\n denio doi ejd\nien\n jndi2 2d"
	a:=bufio.NewScanner(strings.NewReader(s))
	for a.Scan(){
		l:=a.Text()
		fmt.Println(l)
	}

}

