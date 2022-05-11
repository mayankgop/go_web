package main

import (
	"bufio"
	"fmt"
	"net"
	// "strings"
	// "text/scanner"
)

func main(){
	li,_:=net.Listen("tcp",":8080")
	defer li.Close()
	for{
		con,_:=li.Accept()
		go h(con)
	}
}

func h(con net.Conn){
	a:=bufio.NewScanner(con)
	for a.Scan(){
		l:=a.Text()
		fmt.Println(l)
	}
	defer con.Close()
	fmt.Println("code not able to reach here")

}

