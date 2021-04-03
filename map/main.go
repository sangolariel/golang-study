package main

import "fmt"

func main() {

	colors := make(map[string]string) // make an empty map

	colors["white"] = "#ffffff"

	delete(colors, "white") // delete a key value pair

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "$4bf645",
	// }
	fmt.Println(colors)
}
