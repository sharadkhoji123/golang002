package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

// type person struct {
// 	Name string
// 	Age  int
// 	City string
// }

func main() {

	rcvd := []byte(`{"Fname":"James","Lname":"Bond","Items":["Don't","Panic","When","Scared"]}`)
	//rcvd := `{"Name":"John", "Age":31, "City":"New York"}`
	var p2 person
	err := json.Unmarshal([]byte(rcvd), &p2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("%v", p2)
	fmt.Println()
	for i, v := range p2.Items {
		fmt.Println(i, v)
	}
}
