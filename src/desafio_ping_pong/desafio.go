package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(1 * time.Second)
	}

}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(1 * time.Second)
	}
}

func printer(c1 chan string, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping(c1)
	go pong(c2)
	go printer(c1, c2)

	var input string
	fmt.Scanln(&input)
}
