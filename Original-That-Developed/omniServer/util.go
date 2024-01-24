package omniServer

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvPortNumber(portNumber int) string {
	builder := strings.Builder{}
	portStr := strconv.Itoa(portNumber)
	builder.WriteString(":" + portStr)
	listeningPort := builder.String()
	builder.Reset()
	return listeningPort
}

// Build and improve
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		log.Fatal(err)
		return false, err
	}
	if os.IsNotExist(err) {
		log.Fatal("File path does not exist")
		return false, err
	}
	return true, nil
}

func CheckValidIP(ip string) bool {
	if ip == "" {
		return false
	}
	checkIP := strings.Split(ip, ".")
	if len(checkIP) != 4 {
		return false
	}
	for _, ip := range checkIP {
		if octet, err := strconv.Atoi(ip); err != nil {
			return false
		} else if octet < 0 || octet > 255 {
			return false
		}
	}
	return true
}
