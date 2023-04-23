package ninja

import (
        "strings"
        "encoding/base64"
        "bytes"
        "compress/gzip"
        "fmt"
        "log"
)

// This module will be executed before argparsing of command line arguments to evade IDS system
// In conjunction with how ninja-gencmd transpires


//First Check any encoding by characteristics
func checkEncoding(s string) (result string) {
        // Check if its Gzip
        if strings.Contains(s, '\\') {
                decodeGzip(s)
	}



        // Check if its Base64
        if //
        result, err = decodeBase64(s)

	if strings.Contains(s, '  ') {
		result = removeExcessWS(s)
	}
}

func decodeGzip(s string) {
        var buf bytes.Buffer
        sAsBytes := []bytes(s)
        gz := qzip.NewReader(&buf)
        // TODO


        gz.Close()

}

// Decode Base64 String
func decodeBase64(s string) (result string, err error) {
        result, err = base64.StdEncoding.DecodeString(s)
        if err != nil {
                fmt.Printf("Error decoding from base64: %s ", err.Error())
                return s, err
        }
        return result, err
}


// Remove extra spaces
func removeExcessWS(s string) string {
	builder := strings.Builder{}
	splitS := strings.Split(s, "")
	sliceSize := len(splitS) - 1

	for i := 0; i <= sliceSize; i++ {
		if !strings.Contains(splitS[i], " ") {
			builder.WriteString(splitS[i])
			if i != sliceSize && strings.Contains(splitS[i+1], " ") {
				builder.WriteString(" ")
			}
		}

	}
	return builder.String()
}

// Change case to lower case
func convLowerCase(s string) string {
	return strings.ToLower(s)
}

