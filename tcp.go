package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

	// Connect to the server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("Error connecting to server: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Use bufio to read from stdin and write to the connection
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(conn)
