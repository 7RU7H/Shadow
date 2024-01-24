package omniServer

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "strconv"
        "strings"
        "time"

        "github.com/7ru7h/Shadow/omniServer/cli.go"
        "github.com/7ru7h/Shadow/omniServer/console.go"
        "github.com/7ru7h/Shadow/omniServer/tls.go"
	"github.com/7ru7h/Shadow/omniServer/util.go"
)

//
//
// (NAME OF CONCEPT THAT MANAGES) -> server1,server2,...
// Seperation of the methods as I am double server and IDdatabase
// X-server: web-server.go, proxy-server.go
// CURRENT IDEA Database needs to be:
// - part of larger struct that: map[string](pointer) points to Server structs, ID database etc 
// - initialisation of array to make ID database - ID need negative space for stopped servers
// 

// Are negative ID is a good way of managing this why not flags
//
       
// 
// Creation to termination 
// Memory Arenas

// 
// IDs
// Memory Arenas and how.. 
// 
		
type Server struct {
        ServerType int // Integer reference for each - decimalise as in 0 - 9 is debug; 10 is webserver, 20 proxy, 30 capture - 11 is then an option for feature extension of a webserver
        ServerID int // 0 ID is temporary ID till checks, negative digits are stopped server IDs
        ServerWithCtx *http.Server
        Ctx Context
        CancelCtx CancelFunc
        Mux *ServerMux
        ServerInfo ServerInfo
        TLSInfo TLSInfo
        NewProc bool
        ProcInfo ProcInfo
}

// For if Server is required to be run as a new process
type ProcInfo struct {
        PID int
        UID int
}

type ServerInfo struct {
        Status int
        Hostnames []string
        TotalHostnames int
        ListeningPort int
        ServerAddr string
}

type TLSInfo struct {
        ServerCertPath string
        ServerKeyPath string
        CertExpiryDays int
}

// 0 ID is set for all initialing servers till checks
func (s *Server) InitServerStruct(hasTLS, hasHosts, newProc bool, argsServerInfo, fromArgsTlsInfo []string) (error) {
	//  
	//
	//
	tls := TLSInfo{}
	if hasTLS {
		checkCertPath, err := util.CheckFileExists(fromArgsTlsInfo[1]) 
		if !checkCertPath {
			//
			return err
		} else {
			tls.ServerCertPath = fromArgsTlsInfo[1]
		}
		checkKeyPath, err := util.CheckFileExists(fromArgsTlsInfo[2])
		if !checkKeyPath {
			//
			return err
		} else {
			tls.ServerKeyPath = fromArgsTlsInfo[2]

		}
        	tls.CertExpiryDays = fromArgsTlsInfo[3]
		s.TLSInfo = tls
	} else { 
		tls.ServerCertPath = "none"
		tls.ServerKeyPath = "none"
        	tls.CertExpiryDays = -1
		s.TLSInfo = tls
	}
	
	s.ServerID = 0

        //EvaluateHostnames return len(arr)
	ServerInfo {
		Status = 0,
                Hostnames = ,
                TotalHostnames = len()-1
		// hostnames =  
		// func () if !hasHosts { hostname = "" } else { hostnameList := fromArgsServerInfo[INDEX] }
		// 
	}
	
}

func (s *Server) CreateServer() (error)  {
        if CheckAvaliableIDs(s.ServerID) || CheckAvaliableIDs() {
                // ID in use
        }
        // ServerType == Integer reference for each - decimalise as in 0 - 9 is debug; 10 is webserver, 20 proxy, 30 capture - 11 is then an option for feature extension of a webserver
        switch s.ServerType {
                case 10: // HTTP Server
                        s.CreateWebServer()
                case 11: // HTTPS Server
                        // Handle TLS certificate generation, custom usage
                        tls.manageTLSCertInit() // pass ??.TLSInfo ->
                        s.CreateWebServer()
                default:
                        if s.ServerType <= 9 {  // Debug ServerType value
		        }
                        // Incorrect s.ServerType

        }

}

func (s *Server) CreateWebServer() (error) {
        // Define Mux first to then pass it in Context Creation
        s.Mux = CreateDefaultWebServerMux()
        // Context creation
        // Assigned to a struct!
        s.ServerWithCtx, s.Ctx, s.CancelCtx = InitServerContext(s.ServerInfo.ListeningPort, s.ServerInfo.ServerAddr, s.Mux)
        return nil
}

//
func (s *Server) StartServer() (error)  {
        if !CheckAvaliableIDs(s.ServerID) {

                // Error no server ID to
        }

        if !s.NewProc {


        } else {
        // Create new process
        TestProcInfo := ProcInfo{}
        // Check errors or assign
        s.ProcInfo = TestProcInfo
        }

        if errors.Is(err, http.ErrServerClosed) {
                fmt.Printf("%s closed\n", ServerID, err)
                log.Fatal("%s closed\n", ServerID, err)
                return err
        } else if err != nil {
                fmt.Printf("Error listening for %s: %s\n", ServerID, err)
                log.Fatal("Error listening for %s - ID %d: %s\n", ServerID, err)
                return err
        } else {
                log.Printf("%s is listening...\n", ServerID)
                return err
        }

}

// ConfigServer?
// OPcode = modify Info,  
func ( *) ConfigServer(s *Server) (error)  {
	switch s.ServerType {
	case 10: // DefaultWebServer
		// By s.ServerID, OPcode?
	default:
		if s.ServerType < 10 {
		// Debug ServerType value
		}

		// Invalid server type
	}
}

// BIG REWRITE:
        // Seperate TLS and nonTLS as functions that then are the below go routine,
        // For ServerID, serverCertPath, serverKeyPath
        go func(server *http.Server, ctx Context, cancelCtx CancelFunc) error {

                // ListenAndServer either HTTP or HTTPS server
                // HTTP
                // WILL NOT REQUIRE ListenAndServer() as Contexts will be used was just an idiot
                // go ListenAndServerWebServer()
                if !isTLS {
                // goroutine this function
                        // Better error handling - account for contexts, go routines, when it should return exit out of this block
                        err := server.ListenAndServe()
                        if errors.Is(err, http.ErrServerClosed) {
                                        fmt.Printf("%s closed\n", ServerID, err)
                                        log.Fatal("%s closed\n", ServerID, err)
                                        return err
                        } else if err != nil {
                                        fmt.Printf("Error listening for %s: %s\n", ServerID, err)
                                        log.Fatal("Error listening for %s: %s\n", ServerID, err)
                                        return err
                        } else {
                                        log.Printf("%s is listening...\n", ServerID)
                                        return err
                        }
                        cancelCtx()
                } else {
                        // If HTTPS server
                        //serverStartTime := time.Now()
                        err := http.ListenAndServeTLS(listeningPort, serverCertPath, serverKeyPath, nil)
                        if errors.Is(err, http.ErrServerClosed) {
                                        fmt.Printf("%s closed\n", ServerID, err)
                                        log.Fatal("%s closed\n", ServerID, err)
                                        return err
                        } else if err != nil {
                                        fmt.Printf("Error listening for %s: %s\n", ServerID, err)
                                        log.Fatal("Error listening for %s - ID %d: %s\n", ServerID, err)
                                        return err
                        } else {
                                        log.Printf("%s is listening...\n", ServerID)
                                        return err
                        }
                        cancelCtx()
                }

                return nil
        }()



// Pause server, retain memory and does not deallocate
func (s *Server) StopServer() (error)  {
        if !CheckAvaliableIDs(s.ServerID) {
                // Error no server ID to
        }
}

// What does restart mean and why? - Recreate Context and reassign memory etc
func (s *Server) RestartServer() (error)  {
        if !CheckAvaliableIDs(s.ServerID) {

                // Error no server ID to
        }
}

// CloseServer
func (s *Server) CloseServer() (error)  {
        if !CheckAvaliableIDs(s.ServerID) {
                // Error no server ID to

        }

        // Context termination
        s.CancelCtx()
        <-s.Ctx.Done()
        ServerTerminationTime := time.Now()
        // Checks on termination

        return ServerTerminationTime, time.Now()
}


// manager/handler
// server 



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

        if !util.checkFileExists(requestedFileToDownload) {
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

func saveReqBodyFileHandler(w http.ResponseWriter, r *http.Request) error {
        builder := strings.Builder()
        startTime := time.Now()
        builder.WriteString(os.TempDir() + "/" + strings.ReplaceAll(r.RemoteAddr, ".", "-") + "-T-" + strconv.Itoa(int(time.Now().Unix())))
        filepath :=     builder.Write()
        err := os.Create(filepath,  0644)
        if err != nil {
                log.Fatal(err)
                return err
        }
        err := io.Copy(filepath, r.Body)
        if err != nil {
                log.Fatal(err)
                return err
        }
        defer f.Close()
        endTime := time.Now()
        // Log file and time
        builder.Flush()
        return nil
}

// Mux is a multiplexer to handle routes for Webserver
func CreateDefaultWebServerMux() *ServerMux {
        mux := http.NewServeMux()
        // Setup routes
        mux.HandleFunc("/upload", uploadFileHandler())
        mux.HandleFunc("/download", downloadFileHandler())
        mux.HandleFunc("/saveReqBody", saveReqBodyFileHandler())
        return mux
}

func InitServerContext(lportString, keyServerAddr string, srvMux *ServerMux)  (*http.Server, Context, CancelFunc, error) {
        ctx, cancelCtx := context.WithCancel(context.Background())
        server := &http.Server{
                Addr:    lportString,
                Handler: srvMux,
                BaseContext: func(l net.Listener) context.Context {
                        ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
                        return ctx
                },
        }
        return server, ctx, cancelCtx, nil
}
