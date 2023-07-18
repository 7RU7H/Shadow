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
