# Shadow

Can't think of better name just yet so working title was ninjashell - and now Shadow is only slightly better...

Original just wanting a tcp/udp shell with the hope to pursue interesting protocols. QUIC has its faults and I am sure they will be resolved.

I want this is too be a light-weight and entirely modular and generally a pet project to try different languages and components. This will be slow burn for me to keep my head in programming 


#### AB0m1N4Bl3

As an reason to learn Malware Dev, Analysis and C/C++ I decide to collect code snippet techniques to then patchwork my way to actually understanding the three. This is neither good, but will hopeful be educational for me and itch in the back of my mind to remind me to work on this.


#### Gateway 

C2 Teamserver with a simple golang client to tie all the modules together

- Drones
    - Python3 Drone customizer to automate building

#### Cobwell Module


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
- Hopeful Process injection and migration one distant dayi

#### Monolith -> Zigarhack 

C Server to implement FTPs, FTP, HTTP(s), DNS, LDAP, Kerberos, etc - as modularly as possible in the hope if C dies and Zig replaces it that it will be converted to Zig 

####  Ninja Module

Ninja a the tcp/udp  Golang Netcat/Socat/Ncat-like with file transfer, encrypted transfer and traffic that is a light-weight as possible. I may end up becoming a rust or zig shell given Golang Telemetry

- Note to self on wild ideas see [the remmnant ninja readme](ninja/NINJA-README.md)

#### omniServer

Attempt to be a the "best in slot" web server for:
- Uploading to
- Download from
- Sending data and saving to a file

To define BiS:
- It is Golang - therefore most platforms can run it 
- CLI for multiple servers using golang's `context`s
- Memory arenas for lightless weight performance

#### Perspectishells

A set of abmonation-like mash up of webshell from respective language used to understand and learn more about said languages. Beware I will use ChatGTP, because I not good at webshell languages - You have been warned. Each - to keep with the unimaginative naming convention of most of this entire repository with be:
- cmd.EXT

All webshell I was inspired and stole code from will be credited
