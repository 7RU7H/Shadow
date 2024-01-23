package omniServer

import (
	"fmt"
)

exitCmd := "exit"
helpCmd := "help"
infoCmd := "info"
newCmd := "new"
sessionsCmd := "sessions"
terminateCmd := "kill"

validSubCmdHelp
validSubCmdInfo
validSubCmdNew
validSubCmdSession list info interactive

func TerminateConsole() {
	// Closing messages
	os.exit(1)
}

func PrintStdOut() error {

}

func ListAllHelp() {
	//
	fmt.Printf("Listing and help...")
}

func ListCmdHelp(cmd string) {

	fmt.Printf("Help for %s...", cmd)
}

func CheckValidSession(id string) error {

}

func CheckHelpArgs(cmd []string) error {
	if len(cmd) != 2 {
		//
		return err
	}
	err := ValidCmd(cmd[1]) 
	// args error
	if err != nil {
		return err
	}
	return nil
}

func KillSession(id string) error {

}

func ValidKill() error {

}


func RunCmd(cmd []string, cmdlen int) {

	switch cmd[0] {
	case "exit":
		TerminateConsole()
	case "kill":
		err := ValidKill()
		if err != nil {
			break
 		}
		// Consider class a funcs PrintStdOut(KillSession()
	case "help":
		if cmdlen == 1 {
			ListAllHelp()
		} else {
			err := CheckHelpArgs()
			if err != nil {
				break
			}

		}
	}
	return nil	
}



func CheckArgs(cmd string, args []string) error {

	return nil
} 

func CheckCmd(cmd string) error {

	return nil
} 

func RecieveStdIn() error {
	input // till \n bufio.NewReader(os.Stdin).ReadBytes('\n')
	err := CheckCmd(input[0])
	if err != nil {
		PrintStdOut(err)
	}
	err := CheckArgs(input[0], input[1:])
	if err != nil {
		PrintStdOut(err)
	}
	RunCmd(input[0], input[1:])
	return nil
}

func InitialiseConsole() error {

}

