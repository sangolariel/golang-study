package main

import "fmt"

func main() {
	string := returnString()

	fmt.Println(string)
}

func returnString() string {
	return "Func return string type"
}
