package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err con:", err)
		return
	}
	defer conn.Close()

	fmt.Println("conn to serv is done")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("input msg: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))

		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("response serv: %s", response)
	}
}
