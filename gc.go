package ninjashell

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type appEnv struct {
	command         string
	isListener      bool
	isClient        bool
	isFileTransfer  bool
	sourcePort      int
	destinationPort int
	ipAddress       string
	srcFilepath     string
	dstFilepath     string
	shellSpecifier  string
	password        string
	isEncrypted     bool
	isUDP           bool
	a.progressBar
	supportedShells []string
}

func (a *appEnv) setDefaultConfig() {
	a.command = ""
	a.isListener = false
	a.isClient = false
	a.isFileTransfer = false
	a.sourcePort = 60000
	a.destinationPort = 60000
	a.ipAddress = ""
	a.srcFilepath = ""
	a.dstFilepath = ""
	a.shellSpecifier = ""
	a.password = ""
	a.isEncrypted = false
	a.supportedShells = []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}
	a.progressBar
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
	checkIP := strings.split(ip, ".")
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

func (app *appEnv) fromArgs(args []string) error {
	app.setDefaultConfig()
	//Flag variables
	var password, shellSpecifier, ipAddress, srcFilepath, dstFilepath string
	var sourcePort, destinationPort int
	var isEncrypted, isUDP, progressBar bool
	var validFile, validShell, validPassword, validShellSpecifier, validIP, validPort bool
	supportedShells := []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}

	//Subcommands
	listenerCommand := flag.NewFlagSet("-l", flag.ExitOnError)
	clientCommand := flag.NewFlagSet("-c", flag.ExitOnError)
	fileTransferCommand := flag.NewFlagSet("-f", flag.ExitOnError)

	//Listener subcommand flags
	listenerCommand.IntVar(&sourcePort, "-p", "", "Sourc port")
	listenerCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
	listenerCommand.StringVar(&password, "-k", "", "Password is require for encryption")
	listenerCommand.StringVar(&dstFilepath, "-f", "", "File path")
	listenerCommand.BoolVar(&isUDP, "-u", false, "UDP listener, Does not support file transfer use -f subcommand on client ")
	listenerCommand.BoolVar(&progressBar, "-b", false, "Optional progress bar to view file transfer progress")

	//Client subcommand flags
	clientCommand.IntVar(&destinationPort, "-d", "", "Destination port")
	clientCommand.StringVar(&ipAddress, "-i", "", "IP address")
	clientCommand.BoolVar(&isEncrypted, "-e", false, "Encryption")
	clientCommand.StringVar(&password, "-k", "", "Password is required for encryption")
	clientCommand.StringVar(&shellSpecifier, "-s", "", "Shell specifier")
	clientCommand.BoolVar(&isUDP, "-u", false, "UDP client, Does not support file transfer use -f subcommand on client ")

	//File transfer subcommand flags
	fileTransferCommand.IntVar(&destinationPort, "-d", "", "Destination port")
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
		app.isEncrypted = isEncrypted
		app.password = password
		app.dstFilepath = dstFilepath
		app.isUDP = isUDP
	case "-c":
		app.isClient = true
		clientCommand.Parse(args[2:])
		app.destinationPort = destinationPort
		app.ipAddress = ipAddress
		app.isEncrypted = isEncrypted
		app.password = password
		app.shellSpecifier = shellSpecifier
		app.isUDP = isUDP
	case "-f":
		app.isFileTransfer = true
		fileTransferCommand.Parse(args[2:])
		app.destinationPort = destinationPort
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
		}
		if app.shellSpecifier != "" {
			return fmt.Errorf("Shell specifier is not supported for listener")
		}
		if app.sourcePort > 0 || app.sourcePort < 65535 {
			validPort = true
		}
	}

	if app.isClient {
		if app.destinationPort > 0 || app.destinationPort < 65535 {
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
					validShellSpecifier = true
				}
			}
			if !validShell {
				return fmt.Errorf("Shell specifier is not supported for client")
			}
		}
	}

	if app.isFileTransfer {
		if app.destinationPort > 0 || app.destinationPort < 65535 {
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

func (app *appEnv) run() error {
	//Switch based on command passed by user
	switch app.command {
	case "-l":
	case "c":
	case "-f":
	default:
		return fmt.Errorf("Invalid %s parsed and handed to run()", app.command)
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
		return 1
	}
	return 0
}
