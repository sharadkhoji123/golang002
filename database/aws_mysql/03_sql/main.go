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
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/customers", customers)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`drop table customer;`)
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "Table Dropped")

}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`delete from test02.customer where name="John"`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "Record deleted", n)
}
func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`update test02.customer set name="John" where name="James"`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "Record updated", n)
}
func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`Insert into customer values("James");`)
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "Rows Inserted ", n)

}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`create table patients(name varchar(20));`)
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Patients Table Created with rows", n)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully Connected")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)

	//data to be used in query
	var s, name string
	s = "Retrieved Records:\n"

	//Query
	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func customers(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select name from customer;`)
	check(err)
	var s, aname string
	s = "Retrieved Customers: \n"
	for rows.Next() {
		err := rows.Scan(&aname)
		check(err)
		s += aname + "\n"
	}
	fmt.Fprintln(w, s)
}
