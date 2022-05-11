package main

import (
	"fmt"
	"io"
	"net"
)

func main(){
	li,_:=net.Listen("tcp",":8080")
// listen asks for two strings:
// 1) on which network you want to listen on
// 2) on which port you want to listen on

// Listen() returns us a listner and an error

// Listener is an interface which has three methods:
// 1) Accept() : it waits and returns the next connection to the listner
// 2) Close()  : it closes the listner
// 3) Addr()   : it returns the listners network address and unblocks the blocked Accept operations

	defer li.Close()


	for{
		con,_:=li.Accept()  // con is a conn interface that has read and write methods

	io.WriteString(con,"\nHello from tcp sever\n")
	fmt.Fprintln(con,"how s it")
	con.Close()
	}
}