package omniServer 

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//cli.go
)

// https://drstearns.github.io/tutorials/goweb/
// https://tutorialedge.net/golang/go-file-upload-tutorial/U
// https://medium.com/@harisshafiq08/file-upload-server-in-golang-1db6f888fb47
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go


// Upload file - filename
func uploadFileHandler(w http.ResponseWriter, r *http.Request) error {

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    //r.ParseMultipartForm(10 << 20)

	// Get filename from body of r.Body
	
	// FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file

    log.Printf("/upload/%s - Upload requested by ...", handler.Filename)
    file, handler, err := r.FormFile()
    if err != nil {
		// Error retrieving file of filename
        return err
    }
	startTime := time.Now()
    defer file.Close()
	//log.Print("",  ) File upload request success 
	//log.Print("",  ) File upload INFO: 
    log.Printf("Uploaded File: %+v\n", handler.Filename)
    lof.Printf("File Size: %+v\n", handler.Size)
    log.Printf("MIME Header: %+v\n", handler.Header)
	//log.Print("",  ) File upload request success 

    // Create a temporary file within our temp-images directory that conforms to a naming scheme
    tempFile, err := ioutil.TempFile(tmpUploadDir, "tmp-")
    if err != nil {
        // Error creating temporary file
		return err
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        // Failed to read file being uploaded to byte array
		return err
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
	endTime := time.Now()
    //Return that we have successfully uploaded our file!
	log.Printf("Successfully Uploaded File - %s \n", handler.Filename)
    fmt.Fprintf(w, "Successfully Uploaded File - %s \n", handler.Filename)
	return nil
}

// Download file - filename
func downloadFileHandler(w http.ResponseWriter, r *http.Request) error {
	
	// client := Headers - IP User-Agent
	// requestedFileToDownload := 
	
	if !checkFileExists(requestedFileToDownload) {
        w.WriteHeader(404)
        w.Write([]byte("404\n"))
		log.Fatal("Failed to Download file: %s - requested by: %s, %s", requestedFiletoDownload, clientIP, clientUA)
		// Fail to download file error
		return err
    } else {  
		startTime := time.Now()
		log.Printf("Downloading file at:  %s - requested by: %s, %s", requestedFiletoDownload, clientIP, clientUA)
	}
	endTime := time.Now()
	log.Printf("Successfully Download File - %s by %s\n", handler.Filename, clientIP, clientUA)
	return nil
}

// Save Body - filename, data
func saveReqBodyFileHandler(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// ioutil unable to read requeset body
		return err
	} else {
		startTime := time.Now()
		// filename  
		// create file
		// write to file
		// close file
	}
	endTime := time.Now()
	return nil
}

func main() {
	// Testing values
	interface := "eth0"
	vhost := "testwebserver.nvm"
	isTLS := false
	customCert := ""
	listeningPort := 443 
	ipAddress := "127.0.0.1"
	tlsCert := "/path/to/cert"
	uploadPath := "/path/to/upload"
	serverCertPath := "/path/to/cert"
	serverKeyPath :=  "/path/to/cert"

	appStartTime := time.Now()
	// CLI 
	//banner := cli.Banner()
	//fmt.Printf("\n%s\n", banner)
	var userDefinedServerKeyPath string
	var userDefinedServerCertPath string

	// 0: Default 30
	// 1: Randomised
	// 2: Customised
	var certDaysSettings int 
	// If 2 requires != 0, 
	var	userDefinedCertExpiryDays int
	
	var certExpiryDaysSeed string
	var certExpiryDaysRangeLowerBound int 
	var certExpiryDaysRangeUpperBound int
	
	var certExpiryDaysRand int


	// MetaHander - to create, run, close servers - isTLS, vhost, interface, listeningPort, ipAddress
	// Mux is a multiplexer to handle routes
	mux := http.NewServeMux()

	// Make TLS or HTTP server
		// 
		// TLS

	// Default - host a page
	// Host main page

	// Setup routes
	
	
	mux.HandleFunc("/upload", uploadFileHandler())
	mux.HandleFunc("/download", downloadFileHandler()))
	mux.HandleFunc("/saveReqBody", saveReqBodyFileHandler())

	
	// 
	// ListenAndServer either TLS or HTTP server
		// HTTP	

		// TLS
		if customCert != "" {
			serverCertPath = cli.UserDefinedServerCertPath
			serverKeyPath = cli.UserDefinedServerKeyPath
		} else {
			switch certDaysSetting {
			case: 0 // Default 30
				tls.CreateTLSCertKeyPair(30)
			case: 1 // randomised days
				tls.CreateTLSCertKeyPair()
			case: 2 // customised days
				tls.CreateTLSCertKeyPair(userDefinedCertExpiryDays)
			default:
				// error

			}
		}

		// If TLS server TLS certificate
		err := http.ListenAndServeTLS(":443", serverCertPath, serverKeyPath, nil)
    	if err != nil {
      	  log.Fatal("ListenAndServe: ", err)
    	}



	// CloseServer

	// CloseApplication
	appTerminateTime := time.Now()
	// totalRuntime
	fmt.Printf("Application started: %s - Terminated: %s - Runtime: %s", appStartTime, appTerminateTime, totalRuntime)

}

func convPortNumber(portNumber int) string {
	builder := strings.Builder{}
	portStr := strconv.Itoa(portNumber)
	builder.WriteString(":" + portStr)
	listeningPort := builder.String()
	builder.Reset()
	return listeningPort
}

func checkFileExists(path string) (bool error) {
        _, err := os.Stat(path)
        if err == nil {
                log.Fatal(err)
                return false, err
        }
        if os.IsNotExist(err) {
                log.Fatal("File path does not exist")
                return false, err
        }
        return true, nil
}

