# ninjashell
Can't think of better name just yet so working title is ninjashell
Golang Netcat/Socat/Ncat-like with file transfer, encrypted transfer and traffic without become a full blown C2, but the tool you would want see in a list of Tools in a C2.


Ok a month on, been reading and researching a way and a few things came up that meant that the current design is bad for the long term:

1. Found out about Quic
1. Researching more AV/IDS/IPS evasion techniques
1. https://github.com/cheetz/c2
 
First researched a better way from The Hacker Playbook 2; TL DR - Build the Protocol

Benefits being: TCP will be replace with QUIC at some point so might aswell jump big time on that bandwagon early 

## Requirements:
Be QUIC easily ready to QUIC compatibility with backward user - protocol compatible while being smart and using the same functions 
Bypass Firewalls
           - Onioning within a masquade

Bypass IDS - Command variation with smart string variation 
           - Extendable with tech and times
Make it fast and easy to fit with everyone elses automation - keep project alive
Generate Client Executables - Researching 
AV Evasion - Researching 
Make a protocol that itself is onioned so that it hide within legitmate like QUIC, TCP, UDP etc Traffic so that you could still ssl, tls cert host like Socat - OR WHATEVER Quic standards I definately dont know about.
Encrypted Traffic without Cert hadnling


## OnionMask Protocol
```
        Happy SOC team Traffic - Masquade
                        \
                         \
                        Onioning - to evade patternisation
                                \
                                Ninjashell Protocol Encoding
                                                \
                                                Traffic Encrypted
```



Masquading - Legitimate Protocols are alway required
Onioning Mimicry to embedded the Ninjashell protocol

- Idea being to be very stealthy would be to Generate Drip traffic that would end up at the listener 
Like every nth term bit is a part next key part to which seeding would produuce the same result 
The Two castles problem is just solvable by the proof that detirministic seeding of pseudo randomness will alway produce the same randomness.  

 
Jibberish Translation to find EXPECTED Encoding 
Decryption/Encryption


## Quic Research will probably take place after I finish OSCP,
**Dut** re-build everything so that framework is 'Masquade' (TCP, UDP,QUIC,etc) protocol-generic

## Modern Firewall are smart and will get smarter with AI
Solution: Onioning and Masqurading within Legitimate Traffic and Protocols
Long Term Solution: Aversarial AI - Firewall Rule sim VS Brutforcing SMARTLY ninjashell 

## Make User-Automation friendly
Solution be more like Nuclei than some of the older golang cli apps and actually be cli friendly
Keep it simple and clean - it should still be just a shell with alot of hardcore tech

## Command Evasion Client IDS

Solution:

Layed with accepting potentially accepting base64 versions
Command Variation Generator and inside the Ninjashell is a translator
Seperate Binary like Msfvenom
With logging to automate, help with reporting to make manual 



The Main Binary 
 
Checks if its base64/gzip -> convert that and/or convert malformed generated cmd string to regular argparser, then inits argparser stuff

## No Cert Handling

For file transfer chunking and other encrypting and streaming data inside the above. Key are generated on attacker machine. TO Evade hardware monitoring of cryptographic parts of CPU, simple copy and pasting the hash into the command will do for now...Other than creating a shell with a process to then enter the key - I cant think of anything, but that is on the future feature creep list, as this way out of my depth atm


Tried out github copilot to speed up and revise my golang a bit if that is a problem, I apologise to poor copilot being dragged through dealing with me learning. Quite happy to have help and help managing anything, accepting additions and very willing to take a backseat to anyone wants to replace/join the copilot on pedestal of cringe. Willing to end using copilot if a majority vote of anything greater than one decides against me, but for now I just want something that works.


Inspired and refactored from old versions of go netcat clones on github [go-netcat](https://github.com/vfedoroff/go-netcat/blob/master/main.go), [go-nc](https://github.com/opencoff/go-nc/blob/master/gonc.go) and [gocat](https://github.com/sumup-oss/gocat). Wanted to add more functionality like socat encryption and easy ncat file transfers, but in golang where the it can be ran on anything. 

Early Days, having way too much fun with Copilot.



Issues
Hacking Go to like a malware dev to make the binaries small or somehow invent some sort of (or copy more like)
drip feeding process that is similar to staged payload but more packet stealthy and bootstrapping.  


## Future Feature Creep to stop me getting way out of my depth


Process hollowing - for shell to enter hash to evade IDS string matching with hashes on its own Database being a possible simple blueteam mitigation
Payload injection
Process migration awareness
