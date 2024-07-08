package main

import (
	"net/http"
	"fmt"
	"sync"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// x/X subsitute from template  
// xDESIGNPATTERN 


//Credits 
//rollercoast tutorial that the wheels rolling.
//kubucation //www.youtube.com/CHANNEL


//REST API QUESTION CHECKLIST

//Json but LESS IS MORE, beware for feature creeping fields. RELEVANCY
//Use Nouns Instead of Verbs in Endpoints= COMMAND, NOT CRUD -getpost,delete,update ARE USED IN HTTP
//Name Collections with Plural Nouns, Collection = WindowsVER[]COMMANDS  
//Use Status Codes in Error Handling, Does each error have status code 
//Use Nesting on Endpoints to Show Relationships:  avoid nesting that is more than 3 levels deep 
// Use Filtering, Sorting, and Pagination to Retrieve the Data Requested
//DONT EXPOSE ENDPOINTS
//Use SSL for Security -> https! NOT http
// Provide Accurate API Documentation

//Be Clear with Versioning




//Commands: OS, 
//Updates: OS related, Bot related, C2 related 
//Consider what is get request need, and what should not be
//what should a post request 
//Which parts of the database IDs that dynamic, what are static

//C2 server <-> API <-> Define functionality 

//Risk of exposed API or breaking
// hijacked uploads
// discovery
// endpoint counterattack
// logging loss for testing and redteaming

	//Recieve run command ON implanted system not on server system| only recieve the CORRECT commands for the OS - nest Windows with windowscommands, etc
	//
	//download ACCESSIBLE from server | filter get requests to stop api-like LFI or IDOR
	//upload 
	//transmit the success or output of that command


//implant/rat/bot/rootkit/trojan/backdoor/whatever



//FUNCTIONALITY OF the kits not the api
//check to response
//recon 
//exfiltrate
//privEsc
//hide
//propagate
//create shells
//compile software
//x will be served and postable as data! - 


//hide, directory, format, os
//(bootstrappable return after X bootFromBOARD,list of SAFE C2 desposable connectors to proxy through, (serverholder phasekey)+encryptedKey to avoid upline tracing or breaching)

//contexts:
//Setup a nice safe if testable or lastresort so you dont wipe your redteam/blueteam machines
//testable - upon errors transmit logs, backup logs, save states and such in onBoxhandler
//lastresort - upon errors or panic wipe the machine

//if err := exec.CommandContext(,,)
//active
//sleep
//hide & destroy
//wipe


type x struct {
	name string `json:"name"`
	ID string `json:"id"`
}

type xHandler struct {
	sync.Mutex
	store map[string]x
}

func (h * xHandlers) xs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w ,r)
		return
	case "POST":
		h.post(w,r)
		return
	//PUT
	//DELETE
	default:
		w.WriteHeader(http.StatusMethodNotAllowded)
		w.Write([]byte("Method not allowed"))
		return
	}
}

//POST BLOCK
func (h *xHandlers) post(w.httpResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	
	//TODO do common function for error handlings
	//kubucation is all raw Error for each
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())
		return
	}
	

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaTypee)
		w.Write([]byte(fmt.Sprintf("Need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var x X
	err = json.Unmarshal(bodyBytes, &x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())
		return
	}

	//SETTING ID HERE
	x.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[x.ID] = x
	defer h.Unlock()
}



//GET BLOCKopen 
func (h * xHandlers) get(w http.ResponseWriter, r *http.Request) {
	xs := make([]x, len(h.store))
	h.Lock()
	i := 0
	for _, x := range h.store {
		xs[i] = x
		i++
	}
	h.Unlock()
	jsonBytes, err := json.Marshal(xs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//GET X BLOCK
func (h * xHandlers) getX(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL, "/")
	if Len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.Lock()
	x , ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(StatusNotFound)
		return
	}


	jsonBytes, err := json.Marshal(xs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}



func newXHanders() *xHandlers {
	return &xHandlers {
		store: map[string]x{
		//"id": x{ // STATIC DEMO DATA INSERT HERE 
		//data	}
		},
	}
}
func xHandlers(w http.ResponseWriter, r "http.Request") {	
}

type adminPortal struct {
	password string
}

func newAdminPortal() *adminPortal {
	pasword := os.Genenv("ADMIN_PASSWORD") //pentesting WEBAPP 101 dont hardcode passwords!!
	if password == "" { 
		panic("REquired env var ADMIN_PASSWORD not set")
	}

	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := .BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHEader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
		return
	}

		w.Write([]byte("<html><h1>Admin Portal</h1></html>"))
}

func main() {
	admin := newAdminPortal()
	xHandlers := newxHandlers()
	http.HandleFunc("/xhandlers", xHandlers.xs) //switcted based on the method
	http.HandleFunc("/xhandlers/", xHandlers.getX) //getsX
	http.HandleFunc("/admin", admin.handler) //SET ADMIN_PASSWORD!
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
