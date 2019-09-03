package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(id int, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		output := fmt.Sprintf("you sent: %s\n", text)
		log.Printf("[%d] received line: %s", id, text)
		_, errWrite := conn.Write([]byte(output))
		if errWrite != nil {
			log.Printf("[%d] error on write: %s", id, errWrite.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[%d] error on scan: %s", err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}
	address := fmt.Sprintf(":%s", port)

	listener, errListen := net.Listen("tcp", address)
	if errListen != nil {
		log.Fatal(errListen)
	}

	var connID int
	for {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			log.Printf("error on accept: %s", errAccept.Error())
			time.Sleep(time.Second)
		} else {
			go handleConn(connID, conn)
		}

		connID++
	}
}
