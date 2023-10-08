package omniServer

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// https://opensource.com/article/22/9/dynamically-update-tls-certificates-golang-server-no-downtime
// https://medium.com/@harsha.senarath/how-to-implement-tls-ft-golang-40b380aae288
// https://gist.github.com/denji/12b3a568f092ab951456

func CreateTLSCertKeyPair(days int) (err error) {
	builder := strings.Builder{}
	partMkCertCmdStr := "openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days "
	daysStr := strconv.Itoa(days)
	builder.WriteString(partMkCertCmdStr + daysStr)
	mkCertCmd := builder.String()

	keyGen := exec.Command("openssl req -new -sha256 -key server.key -out server.csr")
	if errors.Is(keyGen.Err, exec.ErrDot) {
		keyGen.Err = nil
		return err
	}
	if err := keyGen.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully generated TLS key\n")
	fmt.Printf("Successfully generated TLS key\n")

	mkCert := exec.Command(mkCertCmd)
	if errors.Is(keyGen.Err, exec.ErrDot) {
		mkCert.Err = nil
	}
	if err := mkCert.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully created TLS certificate\n")
	fmt.Printf("Successfully created TLS certificate\n")

	return nil
}

// Needs to pass tlsInfo Struct that is nested in Sever struct
func (t *TLSInfo) ManageTLSCertInit() error {
	if customCert != "" {
		t.ServerCertPath = cli.UserDefinedServerCertPath
		t.ServerKeyPath = cli.UserDefinedServerKeyPath
		log.Printf("Custom TLS Certificate used at: %s", ServerCertPath)
		log.Printf("Custom TLS Key used at: %s", t.ServerKeyPath)
	} else {
		switch t.CertDaysSetting {
		case 0: // Default 30
			err := CreateTLSCertKeyPair(30)
			t.CertExpiryDays = 30
			log.Printf("TLS Certificate created with expiry of days: 30")
		case 1: // randomised days
			// NOT IMPLEMENTED
			// ADD feature
			randomExpiryDays := 0
			err := CreateTLSCertKeyPair(randomExpiryDays)
			t.CertExpiryDays = randomExpiryDays
			log.Printf("TLS Certificate created with expiry of days: %d", randomExpiryDays)
		case 2: // customised days
			err := CreateTLSCertKeyPair(userDefinedCertExpiryDays)
			t.CertExpiryDays = userDefinedCertExpiryDays
			log.Printf("TLS Certificate created with expiry of days: %d", userDefinedCertExpiryDays)
		default:
			// error

		}
	}
	return nil
}
