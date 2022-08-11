package ninjashell

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
func checkEncoding(s string) {
        // Check if its Gzip
        if strings.Contains(s, '\\') {
                decodeGzip(s)


        // Check if its Base64
        if //
        result, err = decodeBase64(s)

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
        originalSize := len(s)-1
        wsCount := strings.Count(s,' ')
        resizeSize := originalSize - (wsCount/2)
	builder := sitnrgs.Builder{}
        //if checkEven(resizeSize) != true {

	result := initStrArray(resizeSize)
	for i := 0; i <=originalSize; i++ {
		if i > 0  && (string.Contains(slice[i], ' ' && string.Contains(slice[i-1], ' ') || !string.Contains(slice[1], ' ') { 
			builder.WriteString(slice[i])
               }
        }
        return builder.String() 
}

// Change case to lower case
func convLowerCase(s string) string {
	return strings.ToLower(s)
}

