package omniServer

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI
// cmds: console, server, help
// - cli to manage closing or if exit or ctrl then all gracefully close
// Aim for Simpler version of:
// https://github.com/BishopFox/sliver/blob/master/client/console/console.go

// Improvement ideas
// Autocomplete? at some points

type AppEnv struct {
	appPID  int
	appEUID int
	command string
}

func (app *AppEnv) setDefaultAppEnv() {
	app.appPID = os.Getpid()
	app.appEUID = os.Geteuid()
	app.command = ""
}

func (app *AppEnv) setStealthAppEnv() {

}

func (app *AppEnv) selectServer() {

}

func (app *AppEnv) selectConsole() error {
	err := Console.InitialiseConsole()
	if err != nil {

		return err
	}
	return nil
}

func (app *AppEnv) run() error {
	//Switch based on command passed by user
	switch app.command {
	case "server":
		err := app.selectServer()
		if err != nil {
			return fmt.Errorf("Error running server: %s", err)
		}
	case "console":
		err := app.selectConsole()
		if err != nil {
			return fmt.Errorf("Error running cli console: %s", err)
		}
	default:
		flag.Usage()
		os.Exit()
	}
	return nil
}

func (app *AppEnv) fromArgs(args []string) error {
	//Flag variables
	var sessionName string
	var quietFlag bool
	var stealthFlag bool

	//Commands
	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)
	consoleCommand := flag.NewFlagSet("console", flag.ExitOnError)

	//Server command flags
	serverCommand.StringVar()
	serverCommand.BoolVar(&stealthFlag, "-S", true, "Applies all possible reductions to reduce detectable footprint - development in progress...")
	serverCommand.BoolVar(&quietFlag, "-q", true, "No Banner and no verbosity")

	//Console command flags
	consoleCommand.StringVar(&sessionName, "-n", "defaultSessionName", "Provide a custom session name")
	consoleCommand.BoolVar(&quietFlag, "-q", true, "No Banner and no verbosity")

	//Help, version, verbose, quiet flags
	var helpFlag, versionFlag string
	flag.StringVar(&helpFlag, "-h", "Help", "Help")
	flag.StringVar(&versionFlag, "-V", "Version", "Version")

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
			fmt.Println("omniServer 0.0")
			os.Exit(1)
		}
	}

	PrintBanner()
	// Set all arguments

	// Check Console arguments

	// Check Server arguements

	// Set AppEnv
	if flag.Lookup(stealthFlag) != nil {
		app.setStealthAppEnv()
	} else {
		app.setDefaultAppEnv()
	}
	return nil
}

func CLI(args []string) int {
	app := AppEnv{}
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

func PrintBanner(banner string) error {
	GetBanner()
	fmt.Printf("%", banner)
}

// Select a banner and return it
func GetBanner() (string, error) {
	var selectedBanner string
	return selectedBanner, nil
}
