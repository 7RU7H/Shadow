package ninjashell

import (
	""
)

//Handles TCP connections and performs synchroninization
//TCP -> Stdout	and Stdin -> TCP
func tcpConnectionHandler(connection net.Conn) {
	chanStdout := copyTCPStream(connection, os.Stdout)
	chanStdin := copyTCPStream(os.Stdin, connection)
	select {
		case <-chanStdout:
			log.Println("Stdout closed")
		case <-chanStdin:
			log.Println("Stdin closed")
	}		
}

//Copy streams between os and tpc stream
func copyTCPStream(src io.Reader, dst io.Writer) <-chan int {
	buf := make([]byte, 1024)
	sync_channel := make(chan int)
	go func() {
		defer func() {
			if con, ok := dst.(*net.TCPConn); ok {
				con.CloseWrite()
				log.Printfc("Closed connection from %s", con.RemoteAddr().String())
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