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

func CreateTLSCertKeyPair(days int) error {
	builder := strings.Builder{}
	partMkCertCmdStr := "openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days "
	daysStr := strconv.Itoa(days)
	build.WriteString(partMkCertCmdStr + daysStr)
	mkCertCmd := builder.String()

	keyGen := exec.Command("openssl req -new -sha256 -key server.key -out server.csr")
	if errors.Is(keyGen.Err, exec.ErrDot) {
		cmd.Err = nil
		return err
	}
	if err := keyGen.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully generated TLS key\n")
	fmt.Fprintf("Successfully generated TLS key\n")

	mkCert := exec.Command(mkCertCmd)
	if errors.Is(keyGen.Err, exec.ErrDot) {
		mkCert.Err = nil
	}
	if err := mkCert.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully created TLS certificate\n")
	fmt.Fprintf("Successfully created TLS certificate\n")

	return nil
}
