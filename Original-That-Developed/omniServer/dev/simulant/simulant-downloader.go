package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"strings"
)

// RequestAction struct
type RequestAction struct {
	AuthRequest     string
	CheckRequest    string
	DownloadRequest string
}

// Properties of pretend human to download files in a human-like way to avoid rate limiting, detection, etc  
type Simulant struct {
	// ultimatejitter - seed a random value between the average time for user to download based on the (age, tool, filetype, service values) more age, using browser, files that larger are probably going to use web services that require more clicks so more time 

}


// open a file containing urls


// Check the file is downloadable else append to file title failedURLs.txt
// use a template checkRequest
// - Video name
// - Size of video
// Beware the neeed for some sort of config for browser and human-like mouse clicking and typing


// queue to manage downloads 

// replacer to replace URL with the url 

// wrapper for curl
func DownloadCMDFile(cmd,request, filename string) error {
	
	return nil
}

// Make a type for Check,Auth,Download
// type RequestAction struct {}
//const DownloadRequest string // 
// request factory - check,download,auth


// Modify default requests of auth,check and download with url 
func ReplaceURLRequest(request, url string) (string,error){
		return result, nil
}



func main() {
	defaultRequests := RequestAction{}
//
	// QUEUE IS ABSTRACTION - REMEMBER Simulate action not automation 	
	// create queues: TODO consider file -> map{},  
	checkingQueue := initStrArray()
	downloadingQueue := initStrArray()


	os.exit() 
}

func initStrArray(n int) []string {
        result := make([]string, n)
        for i := 0; i <= n-1; i++ {
                result[i] = ""
        }
        return result
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


func isDirectory(path string) (bool, error) {
        fileInfo, err := os.Stat(path)
        if err != nil {
                return false, err
        }

        return fileInfo.IsDir(), err
}
