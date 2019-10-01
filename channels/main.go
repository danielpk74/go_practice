package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	st := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range st {
		go checkLink(link, c)
	}

	// for i := 0; i < len(st); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {

		// Function literals
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

		//fmt.Println(l)
		// time.Sleep(3 * time.Second)
		//go checkLink(l, c)
	}

	// Infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might b down!")
		c <- "Might be down I think!"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yeeah it's up!"
}
