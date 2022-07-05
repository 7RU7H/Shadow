package ninjashell

import ( 
	"io"
	"log"
	"net"
	"os"
	"net"
)

type FileTransfer struct {
	conn net.Conn
	file string
	fileSize int
	encryptFile int
	password string
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


//Recieve a non-encrypted file from a sender
func recieveRegularFile(connection net.Conn, chunkTotal, fileSize int) {
	//Create a buffer to hold the file
	buffer := make([]byte, fileSize)
	//Read the file into the buffer
	_, err := connection.Read(buffer)
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
func rejectFileTransfer(connection net.Conn) {
	rejectMessage := []byte("REJECT")
	connection.Write(rejectMessage)
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
