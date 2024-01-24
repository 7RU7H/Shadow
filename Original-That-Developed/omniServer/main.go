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
	
	// "github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/omniServer/cli.go
	// "github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/omniServer/console.go
	"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/omniServer/tls.go
	"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/omniServer/util.go
	"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/"github.com/7ru7h/Shadow/omniServer/metahandler.go 
)





func main() {
	// Testing values
	//interface := "eth0"
	vhost := "testwebserver.nvm"
	isTLS := false
	customCert := ""
	
	var portRequested := 8443 // dummy CLI value
	var listeningPort string
	ipAddress := "127.0.0.1"
	tlsCert := "/path/to/cert"
	uploadPath := "/path/to/upload"
	serverCertPath := "/path/to/cert"
	serverKeyPath := "/path/to/cert"

	appStartTime := time.Now()
	// CLI - complete without console and cli implemented first dumby vars in main -> CLI
	//banner := cli.Banner()
	//fmt.Printf("\n%s\n", banner)
	var userDefinedServerKeyPath string
	var userDefinedServerCertPath string

	// 0: Default 30
	// 1: Randomised
	// 2: Customised
	var certDaysSettings int
	// If 2 requires != 0,
	var userDefinedCertExpiryDays int

	var certExpiryDaysSeed string
	var certExpiryDaysRangeLowerBound int
	var certExpiryDaysRangeUpperBound int

	var certExpiryDaysRand int

	// Post CLI command checks
	// tempDir := //set to linux /tmp/ or Windows\Temp 


	// Check Server Addr 
	// Research interfaces 
	serverAddr := util.CheckValidIP(ipAddress) // ServerAddr required for context creation 

	// check port in use
	listeningPort = util.ConvPortNumber(portRequested)

	// checks completion check for cli and general goodness
	// Call CLI At Some Point
	
	metahandler.DUMMY()
	// MetaHander IDEA  - to create, run, close servers - isTLS, vhost, interface, listeningPort, ipAddress
	// Type of server 
	// Create X server
		// mux for handling requests
	// ListeningAndServer
	// CloseServer 

	// CloseApplication
	appTerminateTime := time.Now()
	// totalRuntime
	log.Printf("Application started: %s - Terminated: %s - Runtime: %s", appStartTime, appTerminateTime, totalRuntime)
	fmt.Fprintf("Application started: %s - Terminated: %s - Runtime: %s", appStartTime, appTerminateTime, totalRuntime)
	os.exit(0)
}


