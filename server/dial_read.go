package main

import (
	"fmt"
	"io/ioutil"
	"net"
	// "strings"
)
func main(){
	con,_:=net.Dial("tcp","localhost:8080")
	defer con.Close()
	bs,_:=ioutil.ReadAll(con)
	fmt.Println(string(bs))
}