package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONNECTION TIMED OUT")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("Code got here")

}
