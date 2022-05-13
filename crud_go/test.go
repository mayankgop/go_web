package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type user struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {

	r := mux.NewRouter()

	// r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/book", pbook).Methods("POST")  //inserting values into books table
	r.HandleFunc("/book", gbook).Methods("GET")   // reading all values from books table
	r.HandleFunc("/book/{mid}", getbyid).Methods("GET")  // reading values by id from books table
	r.HandleFunc("/book/{mid}", delet).Methods("DELETE")   // deleting values by id from books table
	r.HandleFunc("/book/{mid}", upd).Methods("PUT")   // updating values by id from books table


	http.ListenAndServe(":8080", r)
}

func pbook(w http.ResponseWriter, req *http.Request) {
	var u user
	err := json.NewDecoder(req.Body).Decode(&u)
	fmt.Println(u)
	db, err := sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	// _, b := db.Exec(`Insert into books values(?,?)`, u.Id, u.Name)
	query := fmt.Sprintf("insert into books values(%d,'%s')", u.Id, u.Name)
	_, b := db.Exec(query)

	if b != nil {
		fmt.Println("error found in pbook")
	}
	// fmt.Println(err, u)
	// fmt.Fprint(w, u)
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func gbook(w http.ResponseWriter, req *http.Request) {
	var u user
	err := json.NewDecoder(req.Body).Decode(&u)

	db, err := sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	x, b := db.Query("select * from books")

	for x.Next() {
		var u user
		if err := x.Scan(&u.Id, &u.Name); err != nil {
			log.Fatal(err)
		}
		// log.Printf("id %d name   %s\n", u.Id, u.Name)
		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}

	if b != nil {
		fmt.Println("error found in gook")
	}

}

func getbyid(w http.ResponseWriter, req *http.Request) {
	var u user
	m := mux.Vars(req)
	// fmt.Println(m["mid"])
	err := json.NewDecoder(req.Body).Decode(&u)
	// fmt.Println(u)
	db, err := sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	x, _ := strconv.Atoi(m["mid"])
	query := fmt.Sprintf("select * from books where id=%d", x)
	row, b := db.Query(query)

	for row.Next() {
		row.Scan(&u.Id, &u.Name)
	}

	if b != nil {
		fmt.Println("error found in get by id", b)
	}

	uj, _ := json.Marshal(&u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}


func delet(w http.ResponseWriter, req *http.Request) {
	var u user
	m := mux.Vars(req)
	fmt.Println(m["mid"])
	err := json.NewDecoder(req.Body).Decode(&u)
	fmt.Println(u)
	db, err := sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	x, _ := strconv.Atoi(m["mid"])
	query := fmt.Sprintf("DELETE FROM books WHERE id=%d",x)
	_, b := db.Exec(query)

	if b != nil {
		fmt.Println("error found in pbook")
	}


	uj, _ := json.Marshal(&u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func upd(w http.ResponseWriter, req *http.Request) {
	var u user
	m := mux.Vars(req)
	// fmt.Println(m["mid"])
	err := json.NewDecoder(req.Body).Decode(&u)
	// fmt.Println(u)
	db, err := sql.Open("mysql", "admin:qwerty123@tcp(localhost:3306)/bookstore?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	x, _ := strconv.Atoi(m["mid"])
	query := fmt.Sprintf("update books set bookname='%s' where id=%d",u.Name,x)
	_, b := db.Exec(query)

	if b != nil {
		fmt.Println("error found in pbook")
	}


	uj, _ := json.Marshal(&u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

