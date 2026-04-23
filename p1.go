package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	reset   = "\033[0m"
	cyan    = "\033[36m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	bold    = "\033[1m"
)

func savePeerIP(ip string) {
	file, err := os.Create("peer_ip.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(ip)
}

func loadPeerIP() string {
	data, err := os.ReadFile("peer_ip.txt")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func listen() {
	listener, err := net.Listen("tcp", ":9673")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println(bold + green + "Listening for peers on port 9673..." + reset)

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
	peerIP := remoteAddr[:len(remoteAddr)-6]
	savePeerIP(peerIP)

	fmt.Printf(bold+cyan+"\n--- New Chat Started with %s ---\n"+reset, remoteAddr)
	chat(conn)
}

func send(target string, msg string) {
	conn, err := net.Dial("tcp", target+":9673")
	if err != nil {
		fmt.Printf(bold+magenta+"\nCould not connect to %s: %v\n"+reset, target, err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, msg)
}

func main() {
	go listen()

	scanner := bufio.NewScanner(os.Stdin)

	targetIP := loadPeerIP()
	if targetIP == "" {
		fmt.Print(bold + yellow + "Enter peer IP to start chat: " + reset)
		scanner.Scan()
		targetIP = scanner.Text()
	} else {
		fmt.Printf(bold+green+"Using saved peer IP: %s\n"+reset, targetIP)
	}

	conn, err := net.Dial("tcp", targetIP+":9673")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println(bold + green + "Connected! You can now type messages." + reset)
	chat(conn)
}

func chat(conn net.Conn) {
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf(bold+cyan+"[%s]:"+reset+" %s\n", conn.RemoteAddr().String(), scanner.Text())
			fmt.Print(bold + yellow + ">>> " + reset)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(bold + yellow + ">>> " + reset)
		if !scanner.Scan() {
			break
		}
		msg := scanner.Text()
		fmt.Fprintln(conn, msg)
	}
}
