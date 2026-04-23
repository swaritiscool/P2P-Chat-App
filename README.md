# P2P Chat Application

A simple peer-to-peer chat application built with Go for learning purposes. This project focuses on understanding P2P connections and networking fundamentals.

## Overview

This learning project implements a basic P2P chat system where two peers can communicate directly without a central server. The application uses TCP sockets for communication and includes simple file handling capabilities.

## Features

- Direct peer-to-peer communication using TCP
- Automatic IP persistence (saves peer IP to file)
- Color-coded terminal output for better readability
- Concurrent handling of incoming/outgoing messages
- Simple file transfer mechanism (planned enhancement)

## How It Works

1. Each peer runs an instance of the application
2. One peer waits for incoming connections (listener)
3. The other peer initiates connection to the listener
4. Once connected, both peers can send and receive messages concurrently
5. Peer IP addresses are saved to `peer_ip.txt` for future sessions

## File Structure

- `p1.go` - Main application code
- `peer_ip.txt` - Stores the last known peer IP address
- `send.sh` - Script to transfer the compiled binary to another machine
- `Sockets_Simple/` - Directory for simple socket experiments

## Known Issues

1. **Multiple Entries**: Even after established connection on one end, the other user needs to type in the ip manually.
2. **Peer Saving**: When a peer is saved once, it gets called instantly the next time, causing the users to start the app at the exact same time or to completely remove the file(`peer_ip.txt`) itself.

## Planned Enhancements

As part of my learning goals, I plan to implement file handling directly within the messaging system:

1. **File Transfer via Messages**: Instead of separate file transfer protocols, I want to implement a way to send files by embedding them in the message stream
2. **Automatic File Handling**: When a file is "sent" in a message, it gets automatically saved to the recipient's designated directory
3. **File Metadata**: Include filename, size, and type information in the message protocol

## Usage

1. Build the application:
   ```bash
   go build p1.go
   ```

2. On the first machine (listener):
   ```bash
   ./p1
   ```

3. On the second machine (caller):
   ```bash
   ./p1
   ```
   Enter the IP address of the first machine when prompted.

4. To transfer the compiled binary to another machine:
   ```bash
   ./send.sh
   ```
   (Modify the IP address in send.sh to match your target machine)

## Learning Objectives

- Understanding TCP socket programming in Go
- Learning about concurrent programming with goroutines
- Exploring P2P architecture concepts
- Implementing persistent storage for connection information
- Planning and implementing file transfer mechanisms within chat protocols

## Note

This is a learning project focused on understanding the fundamentals of P2P networking. The implementation is intentionally kept simple to facilitate learning and experimentation.
