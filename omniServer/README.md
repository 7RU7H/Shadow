# omniServer

This will hopeful be the basis of all the HTTP(S) Web, Proxy, Etc Servers that are written in Golang.

The objectives:
- A simple, powerful base web server that would replace the `python3 -m http.server $LPORT` in CTFs for something that handles uploading files and saving HTTP Body data to files in a elegant way. 
- HTTPS for the CTF players with minimal fuss of cert management by default
- Has modern golang concurrency, threading and Memory Arena features
- Smallest binary possible prior to Packing
- Have a process that can then have multiple servers because it:
    - a simple CLI 
    - golangs Contexts, interfaces and go routines 



## CLI

FOR now use as a reference to then build out cli.go 
```bash

cmds: manager, server, help

args: 
```

## Ideas

#### CLI 

Manage servers
- Query logs
- Process memory etc
- cli to manage closing or if exit or ctrl then all gracefully close


#### Proxy Server
  ProxyServer - Receive from an address and port and send to a address and port
         -C capture traffic, log and save to file
        https:gobyexample.com/channels - channels to pass data
#### Capture Server

Perform packet capture of all traffic and dump to .pcap or stream back to another omniServer!


#### Web Server Improvements

- Login form to interactive CLI
- Modularise for the sake of modularity (and to make modularise go source to copy pasta to other projects within this repository):
	- Testing Cobwell 

## Help I used, thanks for all the go 

https:www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
https:www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
https:drstearns.github.io/tutorials/goweb/
https:tutorialedge.net/golang/go-file-upload-tutorial/U
https:medium.com/@harisshafiq08/file-upload-server-in-golang-1db6f888fb47
https:www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
https:www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go 
