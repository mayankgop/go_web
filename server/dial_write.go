package main

import (
	"fmt"
	"io"
	"net"
)

func main(){
	li,_:=net.Listen("tcp",":8080")
	defer li.Close()
	for{
		con,_:=li.Accept()
		io.WriteString(con,"hello from tcp\n")
		fmt.Fprintf(con,"how is your day")
		con.Close()	
	}

}