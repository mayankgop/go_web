package main

import (
	"fmt"
	// "io/ioutil"
	"net"
	// "strings"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	fmt.Fprintln(con, "i dialed you")
	defer con.Close()
}
