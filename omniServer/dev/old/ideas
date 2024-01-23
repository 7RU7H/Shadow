package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

type  struct {
	
}

// a struct containing all addresses to all http.servers, 1.20 memory arenas

// interface that handles the interfaces to createNewServer, makeServerListen, and closeServer
// main interface passes *http.servers till closing

//  Functionality
// 	Host pages
// 	Download - FileServer
// 	Upload -
// 	Save reponse to file

// Manage servers
	// Query logs
	// Process memory etc
	// 

//  WebServer 

// go routines
// Multiple servers -

// SSL http.ServeTLS - Server struct has TLSConfig; http.ListenAndServeTLS instead ListenAndServe
// Multithreading
// Memory arenas

// CLI
// cmds: manager, server, help
// - cli to manage closing or if exit or ctrl then all gracefully close

//  ProxyServer - Receive from an address and port and send to a address and port
	// -C capture traffic, log and save to file

	// https://gobyexample.com/channels - channels to pass data
// Login form to interactive CLI 

type Server struct {
	serverType string
	serverAddr string
	serverPort string
	serverInfo struct
}

func createNewServer(serverType, serverAddr, serverPort) {

	// serverInfo set base on type: Proxy, Web
}

func (s *Server) createWebServer(portNumber int, serverAddr string) (*http.Server, context.Context, context.CancelFunc) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	mux := http.NewServeMux()
	httpServer := &http.Server{
		Addr:    s.listeningPort,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverAddr, l.Addr().String())
			return ctx
		},
	}
	return httpServer, ctx, cancelCtx
}

func listenAndServe(server *http.Server, serverName string, ctx context.Context, cancelCtx context.CancelFunc) error {

	err := server.ListenAndServe()
	// Better error handling - account for contexts, go routines, when it should return exit out of this block
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("%s closed\n", serverName)
	} else if err != nil {
		fmt.Printf("Error listening for %s: %s\n", serverName, err)
	} else {
		log.Printf("%s is listening...\n", serverName)
	}
	cancelCtx()
	<-ctx.Done()

	return nil
}

func main() {
	// cmds: manager, server, help

	portNumber := 80           // uer defined
	serverAddr := "serverAddr" // user defined
	// - cli to manage closing or if exit or ctrl then all gracefully close
	// a struct containing all addresses to all http.servers, 1.20 memory arenas

	// interface that handles the interfaces to createNewServer, makeServerListen, and closeServer

	// 	fileServerPath := "/var/www/html"
	//	fileServer := http.FileServer(http.Dir(fileServerPath))
	//    http.Handle("/", fileServer)

	createNewServer(portNumber, serverAddr)
	// pass *http.servers
	//go listenAndServer()
}

func convPortNumber(portNumber int) string {
	builder := strings.Builder{}
	portStr := strconv.Itoa(portNumber)
	builder.WriteString(":" + portStr)
	listeningPort := builder.String()
	builder.Reset()
	return listeningPort
}

// This function is template per page; replace template with pageName and call as http.HandlerFunc("/", pageName)
func templateHandler(w http.ResponseWriter, _ *http.Request) {
	// If to set a status code - refactor out! beware of error reflection
	w.WriteHeader(404)
	w.Write([]byte("404\n"))
	return
}
