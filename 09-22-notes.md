# Notes

Keeping it super simple and extend functionality later...
main.go -> handle.go -> client.go | listener.go

## Listener

prep cert 
check cert
respond ok or drop
stdin/ou listener

## Client

Client is just https://raw.githubusercontent.com/lucas-clemente/quic-go/master/client.go
till I write the listener, test both then start modifying things

## TLS Cert
TLS certs are best generated else where:
-k flag to specific non-default 
