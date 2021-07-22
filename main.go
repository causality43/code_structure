package main

import (
	"code_structure/person"
	"fmt"
)

func main() {
	showHello()
	person.ShowHi() //nama function harus huruf besar jika dari luar package
}

func showHello() {
	fmt.Println("hello world")
}
