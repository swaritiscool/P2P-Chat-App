package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func listen() {
	listener, err := net.Listen("tcp", ":9673")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening for peers on port 9673...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handlePeer(conn)
	}
}

func handlePeer(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("\n--- New Chat Started with %s ---\n", remoteAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("\n[%s]: %s\n", remoteAddr, scanner.Text())
		fmt.Print("Me: ")
	}
}

func send(target string, msg string) {
	conn, err := net.Dial("tcp", target+":9673")
	if err != nil {
		fmt.Printf("\nCould not connect to %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, msg)
}

func main() {
	go listen()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter peer IP to start chat: ")
	scanner.Scan()
	targetIP := scanner.Text()

	conn, err := net.Dial("tcp", targetIP+":9673")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("Connected! You can now type messages.")
	chat(conn)
}

func chat(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Me: ")
		if !scanner.Scan() {
			break
		}
		msg := scanner.Text()

		fmt.Fprintln(conn, msg)
	}
}
