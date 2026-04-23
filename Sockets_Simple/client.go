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
		fmt.Printf("Error While Attempting Connection: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connection made to port 8080")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter Message: ")
		msg, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error Encountered While Sending Message: ", err)
			break
		}
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error While Recieving Message: ", err)
			break
		}
		fmt.Println("Response: ", response)

	}
}
