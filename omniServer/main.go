package omniServer 

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//cli.go
)

// https://tutorialedge.net/golang/go-file-upload-tutorial/U
// https://medium.com/@harisshafiq08/file-upload-server-in-golang-1db6f888fb47

// Upload file - filename
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
    log.Print(" /upload - requested by ...")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)

	// Get filename from body of r.Body
	
	// FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file


    file, handler, err := r.FormFile()
    if err != nil {
		// Error retrieving file of filename
        return
    }
    defer file.Close()
	//log.Print("",  ) File upload request success 
	//log.Print("",  ) File upload INFO: 
    log.Print("Uploaded File: %+v\n", handler.Filename)
    lof.Print("File Size: %+v\n", handler.Size)
    log.Print("MIME Header: %+v\n", handler.Header)
	//log.Print("",  ) File upload request success 

    // Create a temporary file within our temp-images directory that conforms to a naming scheme
    tempFile, err := ioutil.TempFile(tmpUploadDir, "tmp-")
    if err != nil {
        // Error creating temporary file
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        // Failed to read file being uploaded to byte array
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
	log.Print("Successfully Uploaded File - %s \n", handler.Filename)
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// Download file - filename
func downloadFileHandler() {
	
	// client := Headers - IP User-Agent
	//requestedFileToDownload := 
	
	if !fileExists(requestedFileToDownload) {
        w.WriteHeader(404)
        w.Write([]byte("404\n"))
		log.Fatal("Failed to Download file: %s - requested by: %s, %s", requestedFiletoDownload, clientIP, clientUA)
    } else {  

		log.Print("Downloaded file: %s - requested by: %s, %s", requestedFiletoDownload, clientIP, clientUA)
	}
}

// Save Body - filename, data
func saveBodyFileHandler() {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	} else {
		// filename  
		// create file
		// write to file
		// close file
	}
}

func main() {
	// Testing values
	interface := "eth0"
	vhost := "testwebserver.nvm"
	isTLS := false
	listeningPort := 443 
	ipAddress := "127.0.0.1"
	tlsCert := "/path/to/cert"
	uploadPath := "/path/to/upload"

	// CLI 
	//banner := cli.Banner()
	//fmt.Printf("\n%s\n", banner)

	// MetaHander - to create, run, close servers - isTLS, vhost, interface, listeningPort, ipAddress


	// Make TLS or HTTP server
		// 
		// TLS

	// Default - host a page
	// Host main page

	// Setup routes

	http.HandleFunc("/upload", uploadFileHandler())
	http.HandleFunc("/download", downloadFileHandler()))
	http.HandleFunc("/savebody", saveBodyFileHandler())


	// 
	// ListenAndServer either TLS or HTTP server
		// HTTP	

		// TLS
		// If TLS server TLS certificate



	// CloseServer

	// CloseApplication
}


