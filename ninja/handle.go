package ninja

import (
        "errors"
        "flag"
        "fmt"
        "log"
        "os"
        "strconv"
        "strings"
        "time"

	"server"
	"client"
	"packet"
)

type appEnv struct {
        isServer      bool
        isClient        bool
        sourcePort      string
        destinationPort string
        ipAddress       string
        remoteAddress   string
        shellSpecifier  string
        supportedShells []string
}

func (app *appEnv) setDefaultConfig() {
        app.isServer = false
        app.isClient = false
        app.sourcePort = "0"
        app.destinationPort = "0"
        app.ipAddress = ""
        app.remoteAddress = ""
        app.shellSpecifier = ""
        app.supportedShells = []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}
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


func (app *appEnv) fromArgs(args []string) error {
        app.setDefaultConfig()
        //Flag variables
        var password, shellSpecifier, ipAddress string
        var testValidPortNum int
        var validIP, validPort bool
        supportedShells := []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}

        //Subcommands
        serverCommand := flag.NewFlagSet("-l", flag.ExitOnError)
        clientCommand := flag.NewFlagSet("-c", flag.ExitOnError)
       
	//Server subcommand flags
        serverCommand.StringVar(&sourcePort, "-p", "0", "Source port")
        serverCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
        serverCommand.StringVar(&password, "-k", "", "Password is require for encryption")
        serverCommand.StringVar(&dstFilepath, "-f", "", "File path")
        serverCommand.BoolVar(&isUDP, "-u", false, "UDP server, Does not support file transfer use -f subcommand on client ")
        serverCommand.BoolVar(&progressBar, "-b", false, "Optional progress bar to view file transfer progress")

        //Client subcommand flags
        clientCommand.StringVar(&destinationPort, "-d", "0", "Destination port")
        clientCommand.StringVar(&ipAddress, "-i", "", "IP address")
        clientCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
        clientCommand.StringVar(&password, "-k", "", "Password is required for encryption")
        clientCommand.StringVar(&shellSpecifier, "-s", "", "Shell specifier")
        clientCommand.BoolVar(&isUDP, "-u", false, "UDP client, Does not support file transfer use -f subcommand on client ")

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
                        fmt.Println("Ninja Shell v0.2")
                        os.Exit(1)
                }
        }
        //Check user input and store in appEnv
        switch args[1] {
        case "-l":
                app.isServer = true
                serverCommand.Parse(args[2:])
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
        default:
                flag.Usage()
                os.Exit(1)
        }

        if app.isServer {
                if app.shellSpecifier != "" {
                        return fmt.Errorf("Shell specifier is not supported for server")
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

        if !validPort {
                return fmt.Errorf("Invalid port number")
        }
        if !validIP {
                return fmt.Errorf("Invalid IP address")
        }
        return nil
}

func (app. *appEnv) buildAddressString() (string error) {
        builder := strings.Builder{}
        if isServer {
                builder.WRiteString(app.ipAddress,":",app.sourcePort)
        } else {
                builder.WriteString(app.ipAddress, ":",app.destinationPort)
        }
        return builder.String()
}

//Select server based on parameters provided
func (app *appEnv) selectServer() error {
	server.CreateServer(app)
}



//Select client based on parameters provided
func (app *appEnv) selectClient() error {
	client.CreateClient(app)
}



func (app *appEnv) run() error {
        //Switch based on command passed by user
        switch app.command {
                case "-l":
                        err = app.selectServer(); if err != nil {
                                return fmt.Errorf("Error running server: %s", err)
                        }
                case "-c":
                        app.remoteAddress = buildAddressString()
                        err = app.selectClient(); if err != nil {
                                return fmt.Errorf("Error running client: %s", err)
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

