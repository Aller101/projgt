package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("err creating serv: %v", err)

	}

	defer listener.Close()
	log.Println("server is starting")
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err connecting to serv...", err)
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("new conn:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err read data:", err)
			return
		}

		fmt.Printf("getting msg: %s", message)
		conn.Write([]byte("msg is got\n"))
	}
}
