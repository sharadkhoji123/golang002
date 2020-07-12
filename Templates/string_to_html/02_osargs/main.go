package main

import (
	"fmt"
)

func main(){
	name:= os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str :=fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
 	 <head>
    <meta charset="utf-8">
    <title>Hello World</title>
    </head>
  	<body>
	<h1> ` + name + `</h1>
  	</body>
	</html>
	)
}