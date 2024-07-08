package ninjashell

import (
        "errors"
        "flag"
        "fmt"
        "log"
        "ninjashell/tcp.go"
        "ninjashell/udp.go"
        "os"
        "strconv"
        "strings"
        "time"
)

type appEnv struct {
        isListener      bool
        isClient        bool
        isFileTransfer  bool
        sourcePort      string
        destinationPort string
        ipAddress       string
        remoteAddress   string
        srcFilepath     string
        dstFilepath     string
        shellSpecifier  string
        password        string
        isEncrypted     bool
        isUDP           bool
        progressBar     bool
        supportedShells []string
}

func (app *appEnv) setDefaultConfig() {
        app.isListener = false
        app.isClient = false
        app.isFileTransfer = false
        app.sourcePort = "0"
        app.destinationPort = "0"
        app.ipAddress = ""
        app.srcFilepath = ""
        app.remoteAddress = ""
        app.dstFilepath = ""
        app.shellSpecifier = ""
        app.password = ""
        app.isEncrypted = false
        app.supportedShells = []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}
        app.progressBar = false
}

func checkError(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func checkValidIP(ip string) bool {
        if ip == "" {
                return false
        }
        checkIP := strings.Split(ip, ".")
        if len(checkIP) != 4 {
                return false
        }
        for _, ip := range checkIP {
                if octet, err := strconv.Atoi(ip); err != nil {
                        return false
                } else if octet < 0 || octet > 255 {
                        return false
                }
        }
        return true
}

//Checks if the file for transfer exists
func checkFileExists(path string) bool {
        file, err := os.Stat(path)
        if err != nil {
                log.Fatal(err)
                return false
        }
        if os.IsNotExist(err) {
                log.Fatal("%s does not exist", path)
                return false
        }
        return true
}

func (app *appEnv) fromArgs(args []string) error {
        app.setDefaultConfig()
        //Flag variables
        var password, shellSpecifier, ipAddress, srcFilepath, dstFilepath string
        var testValidPortNum int
        var isEncrypted, isUDP, progressBar, validFile, validShell, validPassword, validIP, validPort bool
        supportedShells := []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}

        //Subcommands
        listenerCommand := flag.NewFlagSet("-l", flag.ExitOnError)
        clientCommand := flag.NewFlagSet("-c", flag.ExitOnError)
        fileTransferCommand := flag.NewFlagSet("-f", flag.ExitOnError)

        //Listener subcommand flags
        listenerCommand.StringVar(&sourcePort, "-p", "0", "Source port")
        listenerCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
        listenerCommand.StringVar(&password, "-k", "", "Password is require for encryption")
        listenerCommand.StringVar(&dstFilepath, "-f", "", "File path")
        listenerCommand.BoolVar(&isUDP, "-u", false, "UDP listener, Does not support file transfer use -f subcommand on client ")
        listenerCommand.BoolVar(&progressBar, "-b", false, "Optional progress bar to view file transfer progress")

        //Client subcommand flags
        clientCommand.StringVar(&destinationPort, "-d", "0", "Destination port")
        clientCommand.StringVar(&ipAddress, "-i", "", "IP address")
        clientCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
        clientCommand.StringVar(&password, "-k", "", "Password is required for encryption")
        clientCommand.StringVar(&shellSpecifier, "-s", "", "Shell specifier")
        clientCommand.BoolVar(&isUDP, "-u", false, "UDP client, Does not support file transfer use -f subcommand on client ")

        //File transfer subcommand flags
        fileTransferCommand.StringVar(&destinationPort, "-d", "0", "Destination port")
        fileTransferCommand.StringVar(&ipAddress, "-i", "", "IP address")
        fileTransferCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
        fileTransferCommand.StringVar(&password, "-k", "", "Password")
        fileTransferCommand.StringVar(&srcFilepath, "-f", "", "File path")
        fileTransferCommand.BoolVar(&progressBar, "-b", false, "Optional progress bar to view file transfer progress")

        //Help and version flags
        var helpFlag, versionFlag string
        flag.StringVar(&helpFlag, "-h", "Help", "Help")
        flag.StringVar(&versionFlag, "-v", "Version", "Version")

        //Parse the flags
        if err := flag.Parse(args); err != nil {
                return err
        }
        argsLen := len(args)

        if argsLen > 1 {
                flag.Usage()
                os.Exit(1)
        }

        if argsLen == 1 {
                if flag.Lookup(helpFlag) != nil {
                        flag.Usage()
                        os.Exit(1)
                }
                if flag.Lookup(versionFlag) != nil {
                        fmt.Println("Ninja Shell v0.1")
                        os.Exit(1)
                }
        }
        //Check user input and store in appEnv
        switch args[1] {
        case "-l":
                app.isListener = true
                listenerCommand.Parse(args[2:])
                app.sourcePort = sourcePort
                testValidPortNum = strconv.Atoi(sourcePort)
                app.isEncrypted = isEncrypted
                app.password = password
                app.dstFilepath = dstFilepath
                app.isUDP = isUDP
        case "-c":
                app.isClient = true
                clientCommand.Parse(args[2:])
                app.destinationPort = destinationPort
                testValidPortNum = strconv.Atoi(destinationPort)
                app.ipAddress = ipAddress
                app.isEncrypted = isEncrypted
                app.password = password
                app.shellSpecifier = shellSpecifier
                app.isUDP = isUDP
        case "-f":
                app.isFileTransfer = true
                fileTransferCommand.Parse(args[2:])
                app.destinationPort = destinationPort
                testValidPortNum = strconv.Atoi(destinationPort)
                app.ipAddress = ipAddress
                app.isEncrypted = isEncrypted
                app.password = password
                app.srcFilepath = srcFilepath
                app.isUDP = isUDP
        default:
                flag.Usage()
                os.Exit(1)
        }

        if app.isListener {
                if app.isUDP && app.dstFilepath != "" {
                        return fmt.Errorf("UDP listener cannot save file")
                }
                if app.isEncrypted && app.password == "" {
                        return fmt.Errorf("Password is required for encryption")
                }
                if app.dstFilepath == "" {
                        return fmt.Errorf("Destination file path is required")
                }
                if app.isEncrypted && app.password != "" {
                        validPassword = true
                }
                if app.dstFilepath != "" {
                        if checkFileExists(app.dstFilepath) {
                                return fmt.Errorf("Destination file path is required")
                }
                if app.shellSpecifier != "" {
                        return fmt.Errorf("Shell specifier is not supported for listener")
                }
                if testValidPortNum > 0 || testValidPortNum < 65535 {
                        validPort = true
                }
        }

        if app.isClient {
                if testValidPortNum > 0 || testValidPortNum < 65535 {
                        validPort = true
                }
                if checkValidIP(app.ipAddress) {
                        validIP = true
                }
                if app.isEncrypted && app.password == "" {
                        return fmt.Errorf("Password is required for encryption")
                }
                if app.shellSpecifier != "" {
                        for _, shell := range supportedShells {
                                if app.shellSpecifier == shell {
                                        validShell = true
                                }
                        }
                        if !validShell {
                                return fmt.Errorf("Shell specifier is not supported for client")
                        }
                }
        }

        if app.isFileTransfer {
                if testValidPortNum > 0 || testValidPortNum < 65535 {
                        validPort = true
                }
                if checkValidIP(app.ipAddress) {
                        validIP = true
                }
                if app.isEncrypted && app.password == "" {
                        return fmt.Errorf("Password is required for encryption")
                }
                if app.srcFilepath == "" {
                        return fmt.Errorf("Source file path is required")
                } else {
                        validFile = checkFileExists(app.srcFilepath)
                }
                if app.isUDP {
                        return fmt.Errorf("UDP file transfer is not supported")
                }
                if app.isEncrypted && app.password != "" {
                        validPassword = true
                }
                if app.shellSpecifier != "" {
                        return fmt.Errorf("Shell specifier is not supported for file transfer client")
                }
                if app.sourcePort > 0 || app.sourcePort < 65535 {
                        validPort = true
                }
                if !validFile {
                        return fmt.Errorf("Invalid source file path is required")
                }
        }
        if !validPort {
                return fmt.Errorf("Invalid port number")
        }
        if !validIP {
                return fmt.Errorf("Invalid IP address")
        }
        return nil
}

func (app *appEnv) createUDPListener() error {
        listener, err := net.Listen("udp", app.destinationPort)
        checkError(err)
        log.Println("Listening on", app.destinationPort)
        connection, err := listener.Accept()
        checkError(err)
        log.Println("Connection from ", connection.RemoteAddr())
        udp.UDPConnectionHandler(connection)
}

//func (app *appEnv) createListenerUDPEncrypted() error  {}

//func (app *appEnv) createListenerEncrypted() error  {}

func (app *appEnv) createTCPListener() error  {
        listener, err := net.Listen("tcp", app.destinationPort)
        checkError(err)
        log.Println("Listening on", app.destinationPort)
        connection, err := listener.Accept()
        checkError(err)
        log.Println("Connection from ", connection.RemoteAddr())
        tcp.TCPConnectionHandler(connection)
}

func (app *appEnv) createTCPClient() error  {
        connection, err, := net.Dial("tcp", )
        checkError(err)

}

func (app *appEnv) createTCPClientEncrypted() error {
        connection, err, := net.Dial("tcp", app.ipAddress+":"+app.destinationPort)
        checkError(err)
        log.Println("Connecting to", app.ipAddress+":"+app.destinationPort)
        //Prompt for password
        //Encrypt password 
        //Send Sha512 hash
        //Wait for response
        //If response is OK, continue
        //If response is NO, exit
        //

}

func (app *appEnv) createUDPClient() error  {

}

func (app *appEnv) createUDPClientEncrypted() error  {

}

func (app *appEnv) fileTransferEncrypted() error  {

}

func (app *appEnv) fileTransfer() error  {

}

func (app. *appEnv) buildAddressString() (string error) {
        builder := strings.Builder{}
        if isListener {
                builder.WRiteString(app.ipAddress,":",app.sourcePort)
        } else {
                builder.WriteString(app.ipAddress, ":",app.destinationPort)
        }
        return builder.String()
}

//Select listener based on parameters provided
func (app *appEnv) selectListener() error {
        if app.isUDP {
                if app.isEncrypted {
                        err = app.createUDPEncryptedListener()
                } else {
                err = app.createUDPListener()
                }
        } else if app.isEncrypted {
                err = app.createListenerEncrypted()
        } else {
                err = app.createTCPListener()
        }
        if err != nil {
                return fmt.Errorf("Error setting up Listener: %s", err)
        }
        return nil
}

func (app *appEnv) selectClient() error {
        if app.isUDP {
                if app.isEncrypted {
                        err = app.clientUDPEncrypted()
                } else {
                err = app.clientUDP()
                }
        } else if app.isEncrypted {
                err = app.clientTCPEncrypted()
        } else {
                err = app.createTCPClient()
        }
        if err != nil {
                return fmt.Errorf("Error connecting to server: %s", err)
        }
}

func (app *appEnv) selectFileTransfer() error {
        if app.isUDP {
                return fmt.Errorf("UDP file transfer is not supported")
        }
        if app.isEncrypted {
                err = app.fileTransferEncrypted()
        } else {
                err = app.fileTransfer()
        }
        if err != nil {
                return fmt.Errorf("Error connecting to server with file transfer: %s", err)
        }
        return nil
}


func (app *appEnv) run() error {
        //Switch based on command passed by user
        switch app.command {
                case "-l":
                        err = app.selectListener(); if err != nil {
                                return fmt.Errorf("Error running listener: %s", err)
                        }
                case "-c":
                        app.remoteAddress = buildAddressString()
                        err = app.selectClient(); if err != nil {
                                return fmt.Errorf("Error running client: %s", err)
                        }
                case "-f":
                        app.remoteAddress = buildAddressString()
                        err = app.selectFileTransfer(); if err != nil {
                                return fmt.Errorf("Error running file transfer: %s", err)
                        }
                default:
                        flag.Usage()
                        os.Exit(1)
        }
        return nil
}

func CLI(args []string) int {
        app := appEnv{}
        err := app.fromArgs(args)
        if err != nil {
                log.Fatal(err)
                return 2
        }
        if err = app.run(); err != nil {
                log.Fatal(err)
                fmt.Fprintf(os.Stderr, "Runtime Error: %s\n", err)
                return 2
        }
        return 0
}
  
