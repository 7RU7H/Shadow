package ninja

// The module is used by the ninja to encrypt traffic and decrypt traffic.
// Handling password and file encryption and decryption.
// Password once verified at both end is cached in memory for the duration of the session
// and is used for all subsequent traffic.

import (
	"encoding/base64"
	"crypt/sha3"
	"log"
	"net"
)

type Secret struct {
	//Salt string
	localKey string
}

// Encode the password with Base64 and the use the password encrypt to itself in to be sent to the remote host
func (s *Secret) CreateLocalKey(password string) string {
	b64Password := base64.StdEncoding.EncodeToString([]byte(password))
	return sha3.new512(b64Password)
}

func (s *Secret) InitialiseSecretStruct(password string) {
	s.localKey = s.CreateLocalKey(password)
}

func validateHandshake(buffer []bytes) error {
	if buffer != nil {
		if buffer[256] != 0 || buffer[255] == 0{
			//error invalid buffer content expecting a 256 bit hash
			return errors.New("Invalid buffer content")
		}
		// Additional checks go here
	return nil
}

// Use local password as decryption key for decrypting handshake passed by Listener
// Decrypt the handshake and verify the password, 
func (s *Secret) DecryptHandshake(handshake []bytes)  {
	err := validateHandshake(handshake)
	if err != nil {
		log.Printf("Error validating handshake: %s", err)
		return err
	}
	// Decrypt the handshake with local password

	return
}


func AwaitHandshakeInitiation(connection net.Conn, timeout int) (bool, error) {
	// Wait for a response from the remote host
	buf := make([]byte, 1024)
	nBytes, err := connection.Read(buf)
	//Timeout while loop

	if err != nil {
		log.Printf("Error reading from connection: %s", err)
		return false, err
	} else {
		//Send verification  and request message connection type to remote host
		connection.Write()
		return true, nil
	}
}

func AwaitHandsakeResponse(connection net.Conn, timeout int) (bool, error) {
}


// Initiate the handshake with the remote host from client side
// Wait for a response else timeout after 100 seconds
func (s *Secret) InitiateHandsake(connection net.Conn) error {
	connection, err := connection.Write(s.localKey)

	if err != nil {
		log.Printf("Error writing to connection: %s", err)
		return err
	}
	return nil
}


func (app *gc.appEnv)SelectHandshakeAgent(, connection net.Conn) (err error) {
	agent := Secret{}
	agent.InitialiseSecretStruct(app.Password)
	if !app.isListener {
		err = InitiateHandsake(connection)
		if err != nil {
			log.Printf("Error initiating handshake: %s", err)
			return err
		}
		serverSideVerification, err := AwaitHandshakeResponse(connection, 100)
		if err != nil {
			log.Printf("Error awaiting handshake: %s", err)
			return err
		}
		if app.isFileTransfer {
		}
	}

	} else {
		// Listener waits for handshake from remote host
		clientInitation := err = AwaitHandshakeInitiation()
		if err != nil {
			log.Printf("Error awaiting handshake: %s", err)
			return err
		}
		clientVerification, err := AwaitHandshakeResponse(connection, 100)
		if err != nil {
			log.Printf("Error awaiting handshake: %s", err)
			return err
		}
		if app.isFileTransfer {
			// Server side file recieving branch
		} else {
			if clientInitation && clientVerification {
			//Configure listener to continue listening for client after sending final acceptance message
			} else {
				log.Printf("Error failed handshake checks! first: %t, second: %t ", clientInitation, clientVerification)
				return err
			}
		}

	}
}




//stdin -> till enter key or end of buffer -> encrypt the buffer -> send to remote host ->
//read incoming -> decrypt the buffer -> stdout

//Buffer management for encrypting/decrypting messages

// Encrypt the buffer chunk
//return encrypted buffer chunk, pointer to start and size of buffer chunk
func (s *Secret) EncryptBufferChunk(buffer []bytes, chunkSize int) []bytes {

}

//Buffer management for encrypting/decrypting files would just be based on the (stub + nounce + chunk) ?heartbeat.
