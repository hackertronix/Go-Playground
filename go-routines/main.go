package main

import (
	"fmt"
	"time"
)

func count(id int) {

	for i := 0; i <= 10; i++ {
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {

	fmt.Println("inside main")
	for i := 0; i < 10; i++ {

		go count(i)
	}
	time.Sleep(time.Millisecond * 11000)
}
