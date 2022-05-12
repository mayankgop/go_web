package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		panic(err)
	}
}
func main() {


	var i int
	
	for{
	fmt.Println("(1)  create table  ")
	fmt.Println("(2)  create row  ")
	fmt.Println("(3)  delete   ")
	fmt.Println("(4)  read   ")
	fmt.Println("(5)  update")
	fmt.Println("(6)  exit")
	
		fmt.Println("enter the number to perform operation")
		fmt.Scanf("%d",&i)
		switch i {
		case 1:
			createt()
		case 2:
			create()
		case 3:
			del()
		case 4:
			read()
		case 5:
			update()
		case 6:
			return
		
		}
	}

}


func create() {
	var id int
	var n string
	fmt.Println("enter id and name to create")
	fmt.Scanf("%d %s",&id,&n)
	query := fmt.Sprintf("insert into books values(%d,'%s')", id, n)
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println("id already exists")
		fmt.Println()
		fmt.Println()
		return
	}

	r, err := stmt.Exec()
	if err != nil {
		fmt.Println("id already exists")
		return
	}

	num,err:=r.RowsAffected()
	if err != nil {
		fmt.Println("id already exists")
		return
	}

	fmt.Println("number of rows affected are",num)
	fmt.Println()
	fmt.Println()
}

func del(){
	var id int
	fmt.Println("enter id to delete")
	fmt.Scanf("%d",&id)
    del,_:=db.Query("delete from books where id=?",id)
	defer del.Close()
	fmt.Println()
	fmt.Println()

}

func read(){
	x, _ := db.Query("select * from books")
	for x.Next() {
		var (
			id   int64
			name string
		)
		if err := x.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name   %s\n", id, name)
	}
	fmt.Println()
	fmt.Println()

}

func update(){
	var id int
	var n string
	fmt.Println("enter id to update its name")
	fmt.Scanf("%d",&id)
	fmt.Println("enter new name")
	fmt.Scanf("%s",&n)
	query := fmt.Sprintf("update books set bookname='%s' where id=%d ",n,id)
	// stmt, _ := db.Prepare(query)   no need to prepare
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}


	
	fmt.Println()
	fmt.Println()

// 	UPDATE Customers
// SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
// WHERE CustomerID = 1;

}
func createt(){
	query := `CREATE table books(
		id int ,
		bookname varchar(20),
		primary key(id)
		)`
	_,err:=db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}

}
