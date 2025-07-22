package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	conn, err := http.Get("localhost:8000/ws")
	if err != nil {
		fmt.Println("err con:", err)
		return
	}
	defer conn.Body.Close()

	// conn * websocket.Conn

	fmt.Println("conn to serv is done")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("input msg: ")
		message, _ := reader.ReadString('\n')
		// conn.Write([]byte(message))
		conn.Body.Read([]byte(message))

		// response, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Printf("response serv: %s", response)
	}
}
