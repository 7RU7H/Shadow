package ninjashell

// The module is used by the ninjashell to encrypt traffic and decrypt traffic.
// Handling password and file encryption and decryption.
// Password once verified at both end is cached in memory for the duration of the session
// and is used for all subsequent traffic.

import (	
	"crypto/sha3"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Secret struct {
	Password string // From the gc.go module refactor out
	Salt string
	localKey string
	remoteKey string
}

// Initiate the handshake with the remote host from client side
// Wait for a response else timeout after 100 seconds
func InitiateEncryptedHandsake(connection net.Conn, password string) (bool, error) {
	encryptedGreeting := EncryptPassword(password)
	// Initiate the handshake with the remote host
	connection, err := connection.Write(encryptedGreeting)
	verification, err := AwaitHandshakeResponse(connection)
	if err != nil {
		log.Printf("Error writing to connection: %s", err)
		return false, err
	}


// Encode the password with Base64 and the use the password encrypt to itself in to be sent to the remote host
func (s *Secret) EncryptPassword(password string) []bytes {
	b64Password := base64.StdEncoding.EncodeToString(password)
	return sha3.new512(b64Password)
}

func AwaitHandshakeInitiation(connection net.Conn, timeout int) (bool, error) {
	// Wait for a response from the remote host
	buf := make([]byte, 1024)
	nBytes, err := connection.Read(buf)

	if err != nil {
		log.Printf("Error reading from connection: %s", err)
		return false, err
	} else {
		//Send localkey to remote host
		connection.Write(sha3.new512(localKey))
		return true, nil
		}
	}

func AwaitHandsakeResponse(connection net.Conn, timeout int) (bool, error) {
}

// Use local password as decryption key for decrypting handshake passed by Listener
// Decrypt the handshake and verify the password
func (s *Secret) DecryptHandshake(password string, handshake []bytes)  {}
