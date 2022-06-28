package ninjashell

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"ninjashell/cli"
)




//Create a key for decryption 
func generateKey(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return string(h.Sum(nil))
}

func encryptBufferContent(buffer []byte) []byte {
	//read buffer and encrypt it
	h := sha256.New()
	h.Write(buffer)
	return h.Sum(nil)
}

func encodeBuffer(buf []byte) []byte {
	return bytes.Replace(buf, []byte("\n"), []byte("\\n"), -1)
}

//Reads a file and splits it into chunks
func splitFileToChunks(file string, chunkSize int) [][]byte {
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var chunks [][]byte
	for i := 0; i < len(fileContent); i += chunkSize {
		chunks = append(chunks, fileContent[i:i+chunkSize])
	}
	return chunks
}

func countChunks(chunks [][]byte) int {
	return len(chunks)
}

//Append a nounce of a chunk to the slice of nounces
func appendToNounceStub(nounceStub []byte, nounce []byte) []byte {
	return append(nounceStub, nounce...)
}

func createnounceStub(nounceSize int) []byte {
	//create nounce slices to be prepended before encrypted chunks
	var nounce []byte
	for i := 0; i < nounceSize; i++ {
		nounce = append(nounce, 0)
	}
	return nounceStub
}

//Create Nounce for chunk
func createChunkNounce(chunk []byte) string {	
	h := sha256.New()
	h.Write(chunk)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Marshalls the slice of nounces and chunks into byte stream
func marshallNouncesAndChunks(nounceStub []byte, chunks [][]byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(nounceStub)
	for _, chunk := range chunks {
		buffer.Write(chunk)
	}
	return buffer.Bytes()
}
	

func transformFile(file string) []bytes {
	chunks := splitFileToChunks(file, 1024)
	totalChunks := countChunks(chunks)
	switch encryptFile {
		case 0:
			nouceSlice := createNouceslice(totalChunks)
			for i := 0; i < totalChunks; i++ {
				nounce := createChunkNounce(chunks[i])
				chunks[i] = encryptBufferContent(chunks[i])
				appendToNounceStub(nouceSlice, nounce)
			}
		case 1:
			for i := 0; i < totalChunks; i++ {
			}

		default:
			log.Fatal("Invalid option")
			os.Exit(1)
			}	
	return chunks
}





//Create stub message to warn receiver that the file is being transferred and whether it is encrypted or not
func prepareFileTransferMessage(nounceStubSize, fileSize int, isEncrypted bool) []byte {
	if !isEncrypted && nounceStubSize == 0 {
		fileTransferMessage := []byte("FILETRANSFER" + strconv.Itoa(fileSize))
		return fileTransferMessage 
	} else if isEncrypted && nounceStubSize != 0 {
		fileTransferEncyptedMessage := []byte("FILETRANSFERENCRYPTEDNOUCESTUBIS" + strconv.Itoa(nounceStubSize) + "FILESIZEIS" + strconv.Itoa(fileSize))
		return fileTransferEncyptedMessage
}

//Decodes incoming message and returns the file size and whether it is encrypted or not and if so with nounce stub size
func decodeFileTransferPreparationMessage(message []byte) (bool, stubsize int, filesize int) {
	if strings.Contains(string(message), "ENCRYPTEDNOUCESTUBIS") {
		stubsize = strings.Index(string(message), "ENCRYPTEDNOUCESTUBIS") + len("ENCRYPTEDNOUCESTUBIS")
		filesize = strings.Index(string(message), "FILESIZEIS") + len("FILESIZEIS")
		if stubsize != 0 {		
			return true, stubsize, filesize
		} else {
			log.Fatal("Invalid message nounce size of %v is not valid", stubsize)
			return false, 0, 0
		}
	} else if strings.Contains(string(message), "FILETRANSFER") {
		filesize = strings.Index(string(message), "FILETRANSFER") + len("FILETRANSFER")
		return false, 0, filesize
	} else {
		log.Fatal("Invalid message")
		os.Exit(1)
	}
}

//Receives a sha256 encrypted file and decrypts it with nounceStub and key then the chunks that are written to the file
func receiveEncryptedFile(conn net.Conn, nounceStubSize, fileSize int, ) {
	//read nounceStub
	nounceStub := make([]byte, nounceStubSize)
	_, err := conn.Read(nounceStub)
	if err != nil {
		log.Fatal(err)
	}
	//read file size
	fileSizeBytes := make([]byte, fileSize)
	_, err = conn.Read(fileSizeBytes)
	if err != nil {
		log.Fatal(err)
	}
	//decrypt buffer with nounceStub

//Recieve a non-encrypted file from a sender
func recieveRegularFile(conn net.Conn, fileSize int) {
	//Create a buffer to hold the file
	buffer := make([]byte, fileSize)
	//Read the file into the buffer
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	//Write the buffer to a file
	err = ioutil.WriteFile(file, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Send a rejection message to the file transfer to connection
func rejectFileTransfer() {
	rejectMessage := []byte("REJECT")
	conn.Write(rejectMessage)
}

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

//Accept data from UDP connection and copy it to the stream
func fromUDPToStream(connection *net.UDPConn, dst io.Writer) <-chan net.Addr {
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
func fromStreamToUDP(src io.Reader, dst net.Conn, remoteADdr net.Addr) <-chan net.Addr {
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
func udpConnectionHandler(connection *net.UDPConn) {
	chanStdout := fromUDPToStream(connection, os.Stdout)
	log.Printf("Awaiting connection from %s", connection.RemoteAddr().String())
	remoteAddr := <-chanStdout
	log.Printf("Connected from %s", remoteAddr.String())
	chanStdin := fromStreamToUDP(os.Stdin, connection, addr)
	select {
		case <-chanStdout:
			log.Println("Remote connection closed")
		case <-chanStdin:
			log.Println("Local program terminated")
	}		
}





//Create a progress bar for the file transfer in Stdout for listener
func createProgressBar(fileSize int) {
	bar := pb.New(fileSize)
	bar.SetMaxWidth(80)
	bar.SetRefreshRate(time.Millisecond * 10)
	bar.Start()
	return bar
}

//Update the progress bar for the file transfer in Stdout for listener
func updateProgressBar(bar *pb.ProgressBar, nBytes int) {
	bar.Increment()
}




func main() {
	os.Exit(cli.Run())

	

	if !isUdp {
		log.Println("TCP protocol selected")
		if isListener {
			listener, err := net.Listen("tcp", destinationPort)
			checkError(err)
			log.Printf("Listening on %s", listener.Addr().String())
			con, err := listener.Accept()
			checkError(err)
			log.Printf("Accepted connection from %s", con.RemoteAddr().String())
			tcpConnectionHandler(con)
		} else if ipAddress != "" {
			//Regular shell
			//Specifier shell
			//File transfer
			if validFile != false {
				log.Println("File transfer selected")
				fileTransfer(ipAddress, sourcePort, destinationPort, filepath)
			}
			//Encrypted shell
			//Encrypted file transfer 

			
			con, err := net.Dial("tcp", ipAddress)
			checkError(err)
			log.Printf("Connected to %s", con.RemoteAddr().String())
			tcpConnectionHandler(con)


		} else {
			flag.Usage()
			log.Fatal("Invalid option")
			os.Exit(1)
		}
	} else {
		log.Printf("UDP protocol selected")
		if isListener {
			addr, err := net.ResolveUDPAddr("udp", destinationPort)
			checkError(err)
			con, err := net.ListenUDP("udp", addr)
			checkError(err)
			log.Printf("Listening on %s", destinationPort)
			udpConnectionHandler(con)
		} else if ipAddress != "" {
			addr, err := net.ResolveUDPAddr("udp", ipAddress+":"+destinationPort)
			checkError(err)
			con, err := net.DialUDP("udp", nil, addr)
			checkError(err)
			handleUPDConnection(con)
		}
	}
}
}
