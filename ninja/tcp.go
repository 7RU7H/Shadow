package ninjashell

import (
	"io"
	"log"
	"net"
	"os"
)

//Handles Encrypted TCP connection and performs synchronization
//TCP -> Decrypted Stdout and Encrypted Stdin -> TCP
func EncryptedTCPConnectionHandler(connection net.Conn) {
	chanStdout := DecryptTCPStream(connection, os.Stdout)
	chanStdin := EncryptTCPStream(os.Stdin, connection)
	select {
	case <-chanStdout:
		log.Println("Stdout closed")
	case <-chanStdin:
		log.Println("Stdin closed")
	}
}

//Handles TCP connections and performs synchroninization
//TCP -> Stdout	and Stdin -> TCP
func TCPConnectionHandler(connection net.Conn) {
	chanStdout := CopyTCPStream(connection, os.Stdout)
	chanStdin := CopyTCPStream(os.Stdin, connection)
	select {
	case <-chanStdout:
		log.Println("Stdout closed")
	case <-chanStdin:
		log.Println("Stdin closed")
	}
}

//crypto.EncryptBuffer
//crypto.DecryptBuffer

//Copy streams between os and tpc stream
func CopyTCPStream(src io.Reader, dst io.Writer) <-chan int {
	buf := make([]byte, 1024)
	sync_channel := make(chan int)
	go func() {
		defer func() {
			if con, ok := dst.(*net.TCPConn); ok {
				con.CloseWrite()
				log.Printf("Closed connection from %s", con.RemoteAddr().String())
			}
			sync_channel <- 0 //Notify finished processing
		}()
		for {
			nBytes, err := src.Read(buf)
			if err != nil {
				log.Printf("Error reading from source: %s", err)
				return
			}
			break
		}
		_, err = dst.Write(buf[:nBytes])
		if err != nil {
			log.Printf("Error writing to destination: %s", err)
			return
		}
	}()
	return sync_channel
}
