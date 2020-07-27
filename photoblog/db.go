package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:pass123rds@tcp(database-2.c85adelgml2r.us-east-2.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)

	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func delete(w http.ResponseWriter, req *http.Request) bool {
	u := getUser(req)

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

func insert(w http.ResponseWriter, req *http.Request) bool {

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		t := req.FormValue("phone")
		r := req.FormValue("role")

		t1, err := strconv.Atoi(t)
		if err != nil {
			http.Error(w, "Error thrown while saving contact details", http.StatusForbidden)
			return false
		}
		//check username data
		stmt, err := db.Prepare(`select username from users where username=?;`)
		check(err)

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username Already Taken", http.StatusForbidden)
			return false
		}

		//create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// u := user{un, bs, f, l, t1, r} //ask how to pass directly
		// dbUsers[un] = u

		stmt, err := db.Prepare(`Insert into customer values(un,bs,f,l,t1,r);`)
		check(err)
		r, err = stmt.Exec()
		check(err)
		n, err := r.RowsAffected()
		check(err)
		fmt.Fprintln(w, "Rows Inserted ", n)

	}
}
