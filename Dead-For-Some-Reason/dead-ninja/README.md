# README


This project is dead - xc exists and I learnt alot about project design, golang and golang design, AI, time management, QUIC - much of the go-nc/netcat/xc liek or copied code will go to OmniServer and the ideas will die else where or here. 


# OLD Readme

The original starting point that i want to be a module in Shadow

Inspired and refactored from old versions of go netcat clones on github [go-netcat](https://github.com/vfedoroff/go-netcat/blob/master/main.go), [go-nc](https://github.com/opencoff/go-nc/blob/master/gonc.go) and [gocat](https://github.com/sumup-oss/gocat). Wanted to add more functionality like socat encryption and easy ncat file transfers, but in golang where the it can be ran on anything. I recently discover [xct](https://github.com/xct/xc) that actually works, which is awesome and has some seriously cool functionality. I highly recommend it, but I would want it in rust with the upcoming changes to telemetry. So this maybe become a rust shell.

## Ideas

Ok a month on, been reading and researching a way and a few things came up that meant that the current design are bad for the long term:

1. Found out about Quic
1. Researching more AV/IDS/IPS evasion techniques
1. https://github.com/cheetz/c2
1. Researching better Golang practice and project design like:
        - Anything Project Discovery
        - https://github.com/jpillora/chisel
        - Regular application design best practices

First researched a better way from The Hacker Playbook 2; TL DR - Build the Protocol

Benefits being: TCP will be replace with QUIC at some point so might aswell jump big time on that bandwagon early

#### Requirements:
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


#### OnionMask Protocol
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

Masquading - Legitimate Protocols are alway required, using them to piggyback a subprotocol
Onioning Mimicry to embedded the Ninjashell protocol

- Idea being to be very stealthy would be to Generate Drip traffic that would end up at the listener
Like every nth term bit is a part next key part to which seeding would produuce the same result
The Two castles problem is just solvable by the proof that detirministic seeding of pseudo randomness will alway produce the same randomness.


Jibberish Translation to find EXPECTED Encoding
Decryption/Encryption


#### Quic Research will probably take place after I finish OSCP,
**Dut** re-build everything so that framework is 'Masquade' (TCP, UDP,QUIC,etc) protocol-generic

#### Modern Firewall are smart and will get smarter with AI
Solution: Onioning and Masqurading within Legitimate Traffic and Protocols
Long Term Solution: Aversarial AI - Firewall Rule sim VS Brutforcing SMARTLY ninjashell

#### Make User-Automation friendly
Solution be more like Nuclei than some of the older golang cli apps and actually be cli friendly
Keep it simple and clean - it should still be just a shell with alot of hardcore tech

#### Command Evasion Client IDS

Solution:

Layed with accepting potentially accepting base64 versions
Command Variation Generator and inside the Ninjashell is a translator
Seperate Binary like Msfvenom
With logging to automate, help with reporting to make manual

The Main Binary - Checks if its base64/gzip -> convert that and/or convert malformed generated cmd string to regular argparser, then inits argparser stuff

## No Cert Handling

For file transfer chunking and other encrypting and streaming data inside the above. Key are generated on attacker machine. TO Evade hardware monitoring of cryptographic parts of CPU, simple copy and pasting the hash into the command will do for now...Other than creating a shell with a process to then enter the key - I cant think of anything, but that is on the future feature creep list, as this way out of my depth atm


## Issues
Hacking Go to like a malware dev to make the binaries small or somehow invent some sort of (or copy more like)
drip feeding process that is similar to staged payload but more packet stealthy and bootstrapping.


## Author Notes


Tried out github copilot to speed up and revise my golang a bit if that is a problem, I apologise to poor copilot being dragged through dealing with me learning. Quite happy to have help and help managing anything, accepting additions and very willing to take a backseat to anyone wants to replace/join the copilot on pedestal of cringe. Willing to end using copilot if a majority vote of anything greater than one decides against me, but for now I just want something that works. May get it if I can afford the time and money after OSCP first. Early Days, had way too much fun with Copilot, just for autocomplete around half of what wanted and not lot of what I did not want, but it would be nice to have Ghost in the shell cyborg typing hands for justt the speed of typing

## Future Feature Creep to stop me getting way out of my depth

Process hollowing - for shell to enter hash to evade IDS string matching with hashes on its own Database being a possible simple blueteam mitigation
Payload injection
Process migration awareness
Shell
