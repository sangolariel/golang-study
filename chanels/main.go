package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for {
		go checkLinks(<-c, c)
	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
