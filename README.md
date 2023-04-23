# Shadow

Can't think of better name just yet so working title was ninjashell - and now Shadow is only slightly better...

Original just wanting a tcp/udp shell with the hope to pursue interesting protocols. QUIC has its faults and I am sure they will be resolved.

I want this is too be a light-weight and entirely modular


#### Gateway 

C2 Teamserver with a simple golang client to tie all the modules together

- Drones
    - Drone customizer to automate building

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

####  Ninja Module

Ninja a the tcp/udp  Golang Netcat/Socat/Ncat-like with file transfer, encrypted transfer and traffic that is a light-weight as possible. I may end up becoming a rust or zig shell given Golang Telemetry

- Note to self on wild ideas see [the remmnant ninja readme](ninja/NINJA-README.md)

#### Monolith -> Zigarhack

C Server to implement HTTP(s), DNS, LDAP, Kerberos, etc - as modularly as possible in the hope if C dies and Zig replaces it that it will be converted to Zig

