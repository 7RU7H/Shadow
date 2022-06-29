package ninjashell

import (
	"io"
	"log"
	"net"
	"os"
)

//Accept data from UDP connection and copy it to the stream
func FromUDPToStream(connection *net.UDPConn, dst io.Writer) <-chan net.Addr {
	buf := make([]byte, 1024)
	sync_channel := make(chan net.Addr)
	con, err := src.(*net.UDPConn)
	if !ok {
		log.Fatal("Error casting source to UDP connection")
		return sync_channel
	}
	go func() {
		var remoteAddr net.Addr
		for {
			nBytes, err := con.ReadFromUDP(buf)
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading from source: %s", err)
				}
				break
			}
			if remoteAddr == nil && remoteAddr != addr {
				remoteAddr = addr
				sync_channel <- remoteAddr
			}
			_, err = dst.Write(buf[:nBytes])
			if err != nil {
				log.Printf("Error writing to destination: %s", err)
			}
		}
	}()
	log.Printf("Closed connection from %s", con.RemoteAddr().String())
	return sync_channel
}

//Input data from stream to UDP connection
func FromStreamToUDP(src io.Reader, dst net.Conn, remoteAddr net.Addr) <-chan net.Addr {
	buf := make([]byte, 1024)
	sync_channel := make(chan net.Addr)
	if !ok {
		log.Fatal("Error casting source to UDP connection")
		return sync_channel
	}
	go func() {
		var remoteAddr net.Addr
		for {
			nBytes, err := src.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading from source: %s", err)
				}
				break
			}
			if remoteAddr == nil && remoteAddr != addr {
				remoteAddr = addr
				sync_channel <- remoteAddr
			}
			_, err = con.WriteToUDP(buf[:nBytes], remoteAddr)
			if err != nil {
				log.Printf("Error writing to destination: %s", err)
			}
		}
	}()
	log.Printf("Closed connection from %s", con.RemoteAddr().String())
	return sync_channel
}

//Handle UDP connections and perform synchroninization
func UDPConnectionHandler(connection *net.UDPConn) {
	chanStdout := FromUDPToStream(connection, os.Stdout)
	log.Printf("Awaiting connection from %s", connection.RemoteAddr().String())
	remoteAddr := <-chanStdout
	log.Printf("Connected from %s", remoteAddr.String())
	chanStdin := FromStreamToUDP(os.Stdin, connection, addr)
	select {
	case <-chanStdout:
		log.Println("Remote connection closed")
	case <-chanStdin:
		log.Println("Local program terminated")
	}
}
