package main

import (
	"bufio"
	"encoding/base64"
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

var activeConn = make(chan net.Conn, 1)

func listen() {
	listener, err := net.Listen("tcp", ":9673")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println(bold + green + "Listening for peers on port 9673..." + reset)

	conn, _ := listener.Accept()
	activeConn <- conn

}

func main() {
	go listen()

	go func() {

		scanner := bufio.NewScanner(os.Stdin)
		targetIP := ""
		if targetIP == "" {
			fmt.Print(bold + yellow + "Enter peer IP to start chat: " + reset)
			scanner.Scan()
			targetIP = scanner.Text()
		}

		conn, err := net.Dial("tcp", targetIP+":9673")
		if err == nil {
			fmt.Println("Connection Successful")
			activeConn <- conn
		}
	}()

	conn := <-activeConn

	fmt.Println(bold + green + "Connected! You can now type messages." + reset)
	chat(conn)
}

func chat(conn net.Conn) {
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "FILE:") {
				encoded := strings.TrimPrefix(scanner.Text(), "FILE:")
				encoded, filepath, _ := strings.Cut(encoded, "!")
				data, err := base64.StdEncoding.DecodeString(encoded)
				if err != nil {
					fmt.Printf("Error decoding file: %v\n", err)
					continue
				}
				err = os.WriteFile(strings.Split(filepath, "/")[len(strings.Split(filepath, "/"))-1], data, 0644)
			} else {
				fmt.Printf(bold+cyan+"[%s]:"+reset+" %s\n", conn.RemoteAddr().String(), scanner.Text())
			}
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
		if strings.HasPrefix(msg, "FILE:") {
			filePath := strings.TrimPrefix(msg, "FILE:")
			err := sendFile(conn, filePath)
			if err != nil {
				fmt.Printf("Error sending file: %v\n", err)
			}
		} else {
			fmt.Fprintln(conn, msg)
		}
	}
}

func sendFile(conn net.Conn, filePath string) error {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error While Reading File: ", err)
		return err
	}
	enc := base64.StdEncoding.EncodeToString(dat)
	fmt.Fprintln(conn, "FILE:"+enc+"!"+filePath)
	return nil
}
