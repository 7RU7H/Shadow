# Cobweb Readme

Cobweb is a Golang is https server with an ssh backdoor. The idea being a upgrade from a webshell in that it is a web server that pretends to be part of another webserver. In the hope using in THM KOTH and HTB Battlegrounds as an Easy way back in
- Redirect to legitimate site
- Upload
- Download
- With Provide Header will start negotiation open a backdoor and respond with fingerprint string to connect
- Pretend to be Proxy or VHOST and hopeful someday a parasite page that blends in. 

Objectives:
- Harsh Rate-Limiting 
- Simple Loadbalancing for multiple shell
- Anti-Fingerprinting 
- Internalise TLS information inside the binary not seperate files
- Improve stealth at this endpoint not at LKM Rootkit level
- Hopeful Process injection and migration one distant day

