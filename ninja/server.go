package main

// ChatGTP to the rescue to streamline futher efforts and templating



import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	// Create a TLS certificate for the server
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %s", err)
	}

	// Create a QUIC server listener
	listener, err := quic.ListenAddr("localhost:4242", &tls.Config{
		Certificates: []tls.Certificate{cert},
	}, nil)
	if err != nil {
		log.Fatalf("Failed to create QUIC listener: %s", err)
	}
	defer listener.Close()

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn quic.Session) {
	// Read data from the connection
	stream, err := conn.AcceptStream()
	if err != nil {
		log.Printf("Failed to accept stream: %s", err)
		return
	}
	defer stream.Close()

	data := make([]byte, 1024)
	for {
		n, err := stream.Read(data)
		if err != nil {
			log.Printf("Failed to read from stream: %s", err)
			break
		}
		fmt.Printf("Received %d bytes: %s\n", n, data[:n])
	}
}

// generateTLSConfig generates a TLS config for the server.
// In a real application, you would use a proper certificate and key.
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos
	}
