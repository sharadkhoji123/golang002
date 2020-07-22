package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:pass123rds@tcp(database-2.c85adelgml2r.us-east-2.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)

	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully Connected")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aid,aName FROM amigos`)
	check(err)

	//data to be used in query
	//var s, name string
	type rec struct {
		id   int
		name string
	}
	var r rec
	arecs := []rec{}
	//var id int
	//s = "Retrieved Records:\n"

	//Query
	for rows.Next() {
		err := rows.Scan(&r.id, &r.name)
		//		err := rows.Scan(&name)
		check(err)
		arecs = append(arecs, r)
		//		s += name + "\n"
	}
	fmt.Fprintln(w, "Retrieved Records: ")
	for _, v := range arecs {
		fmt.Fprintln(w, v.id, v.name)
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
