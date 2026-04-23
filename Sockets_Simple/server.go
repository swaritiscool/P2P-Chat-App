package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Encountered Error While Reading Message: %v", err)
			break
		}
		fmt.Printf("Reacieved Message: %s", msg)
		conn.Write([]byte("Server says " + msg))
	}
}

func main() {
	var listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error While Listening On Port 8080: %v", err)
	}
	defer listener.Close()
	fmt.Println("Server Running")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error While Accepting Connection: %v", err)
			break
		}
		go handleConn(conn)
	}

}
