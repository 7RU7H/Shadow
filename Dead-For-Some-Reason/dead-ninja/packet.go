package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)


// TCP


func createClientTCP() {

}

func createServerTCP() {

}

func createConnectionTCP() {

}

func closeConnectionTCP() {

}


// UDP

func createClientUDP() {
}
func createServerUDP() {
}
func createConnectionUDP() {
}
func closeConnectionUDP() {
}


// QUIC

// CUSTOM



	// Connect to the server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("Error connecting to server: %s\n", err)
		os.Exit(1)
	}
	defer conn.CloseTCP()

	// Use bufio to read from stdin and write to the connection
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(conn)
