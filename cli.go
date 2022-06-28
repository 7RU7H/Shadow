package ninjashell

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


type appEnv struct {
	isListener bool
	isClient bool
	fileTranferClient bool
	sourcePort int
	destinationPort int
	ipAddress string
	srcFilepath string
	dstFilepath string
	shellSpecifier string
	password string
	isEncrypted bool
	validFile bool
	validShell bool
	validPassword bool
	supportedShells []string
}

func (a *appEnv) initConfig() {
	a.isListener = false
	a.isClient = false
	a.fileTranferClient = false
	a.sourcePort = 60000 
	a.destinationPort = 60000
	a.ipAddress = ""
	a.srcFilepath = ""
	a.dstFilepath = ""
	a.shellSpecifier = ""
	a.password = ""
	a.isEncrypted = false
	a.validFile = false
	a.validShell = false
	a.validPassword = false
	a.supportedShells = []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}
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

func main {
	currAppEnv:= appEnv{}
	currAppEnv.initConfig()
	var validFile bool
	var validShell bool
	var validPassword bool
	var validShellSpecifier bool
	var validIP bool
	var validPort bool
	supportedShells := []string{"cmd.exe", "powershell.exe", "/bin/bash", "/bin/sh"}


	//Subcommands
	listenerCommand := flag.NewFlagSet("-l", flag.ExitOnError)
	clientCommand := flag.NewFlagSet("-c", flag.ExitOnError)
	fileTransferCommand := flag.NewFlagSet("-f", flag.ExitOnError)

	//Listener subcommand flags
	listenerCommand.IntVar(&sourcePort, "-p", "", "Sourc port")
	listenerCommand.StringVar(&isEncrypted, "-e", "", "Encryption")
	listenerCommand.StringVar(&password, "-k", "", "Password")
	listenerCommand.StringVar(&dstFilepath, "-f", "", "File path")
	listenerCommand.StringVar(&isUDP, "-u", "", "UDP listener, Does not support file transfer use -f subcommand on client ")

	//Client subcommand flags
	clientCommand.IntVar(&destinationPort, "-d", "", "Destination port")
	clientCommand.StringVar(&ipAddress, "-i", "", "IP address")
	clientCommand.StringVar(&isEncrypted, "-e", "", "Encryption")
	clientCommand.StringVar(&password, "-k", "", "Password")
	clientCommand.StringVar(&shellSpecifier, "-s", "", "Shell specifier")
	clientCommand.StringVar(&isUDP, "-u", "", "UDP client")

	//File transfer subcommand flags
	fileTransferCommand.IntVar(&destinationPort, "-d", "", "Destination port")
	fileTransferCommand.StringVar(&ipAddress, "-i", "", "IP address")
	fileTransferCommand.StringVar(&isEncrypted, "-e", "", "Encryption")
	fileTransferCommand.StringVar(&password, "-k", "", "Password")
	fileTransferCommand.StringVar(&srcFilepath, "-f", "", "File path")

	flag.Usage = func() {
		fmt.Printf(flag.Output(), "Usage of %s [options] host:port\nOptions: defaults in '[ ]", os.Args[0])
	}
	flag.StringVar(&helpFlag, "-h", "-h", "Show help")
	flag.StringVar(&versionFlag, "-v", "version", "Show version")
	flag.Parse()

	args := flag.Args()
	argsLen	:= len(args)
	
	if argsLen > 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if argsLen == 1 {
		if flag.Lookup(helpFlag) !=	nil {
		flag.PrintDefaults()
		os.Exit(1)
		}
		if flag.Lookup(versionFlag) != nil {
		fmt.Println("Ninja Shell v0.1")
		os.Exit(1)
		}
	}
		
	switch args[1] {
	case "-l":
		currAppEnv.isListener = true
		listenerCommand.Parse(args[2:])
		currAppEnv.sourcePort = sorcePort
		currAppEnv.isEncrypted = isEncrypted
		currAppEnv.password = password
		currAppEnv.dstFilepath = dstFilepath
		currAppEnv.isUDP = isUDP
	case "-c":
		currAppEnv.isClient = true
		clientCommand.Parse(args[2:])
		currAppEnv.destinationPort = destinationPort
		currAppEnv.ipAddress = ipAddress
		currAppEnv.isEncrypted = isEncrypted
		currAppEnv.password = password
		currAppEnv.shellSpecifier = shellSpecifier
		currAppEnv.isUDP = isUDP
	case "-f":
		currAppEnv.fileTranferClient = true
		fileTransferCommand.Parse(args[2:])
		currAppEnv.destinationPort = destinationPort
		currAppEnv.ipAddress = ipAddress
		currAppEnv.isEncrypted = isEncrypted
		currAppEnv.password = password
		currAppEnv.srcFilepath = srcFilepath
		currAppEnv.isUDP = isUDP
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if currAppEnv.isListener {
		if currAppEnv.isUDP && currAppEnv.dstFilepath != "" {
			fmt.Println("UDP listener for file transfer not supported")
			os.Exit(1)
		}
		if currAppEnv.isEncrypted && currAppEnv.password == "" {
			fmt.Println("Password required for encryption")
			os.Exit(1)
		}
		if currAppEnv.dstFilepath == "" {
			fmt.Println("File path required")
			os.Exit(1)
		}
		if currAppEnv.isEncrypted && currAppEnv.password != "" {
			currAppEnv.validPassword = true
		}
		if currAppEnv.dstFilepath != "" {
			if checkFileExists(currAppEnv.dstFilepath) {
				fmt.Println("File already exists")
				os.Exit(1)
			}
		}
		if currAppEnv.shellSpecifier != "" {
			fmt.Println("Shell specifier invalid and not supported for listener")
			os.Exit(1)
		}
		if currAppEnv.sourcePort > 0 || currAppEnv.sourcePort < 65535 {
			validPort = true
		}
	}

	if currAppEnv.isClient {
		if currAppEnv.destinationPort > 0 || currAppEnv.destinationPort < 65535 {
			validPort = true
		}
		if checkValidIP(currAppEnv.ipAddress) {
			validIP = true
		}
		if currAppEnv.isEncrypted && currAppEnv.password == "" {
			fmt.Println("Password required for encryption")
			os.Exit(1)
		}
		if currAppEnv.shellSpecifier != "" {
			for _, shell := range supportedShells {
				if currAppEnv.shellSpecifier == shell {
					currAppEnv.validShell = true
				}
			}
			if !currAppEnv.validShell {
				fmt.Println("Shell specifier invalid and not supported for client")
				os.Exit(1)
			}
		}
	}

	if currAppEnv.fileTranferClient {
		if currAppEnv.destinationPort > 0 || currAppEnv.destinationPort < 65535 {
			validPort = true
		}
		if checkValidIP(currAppEnv.ipAddress) {
			validIP = true
		}
		if currAppEnv.isEncrypted && currAppEnv.password == "" {
			fmt.Println("Password required for encryption")
			os.Exit(1)
		} 
		}		
		if currAppEnv.srcFilepath == "" {
			fmt.Println("File path required")
			os.Exit(1)
		} else { 
			validFile = checkFileExists(currAppEnv.srcFilepath)
		}
		if currAppEnv.isUDP {
			fmt.Println("UDP file transfer not supported")
			os.Exit(1)
		}
		if currAppEnv.isEncrypted && currAppEnv.password != "" {
			validPassword = true
		}
		if currAppEnv.shellSpecifier != "" {
			fmt.Println("Shell specifier invalid and not supported for file transfer client")
			os.Exit(1)
		}
		if currAppEnv.sourcePort > 0 || currAppEnv.sourcePort < 65535 {
			validPort = true
		}
	}
