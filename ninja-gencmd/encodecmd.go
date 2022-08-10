package ninja-gencmd

import (
        "bytes"
        "compress/gzip"
        "fmt"
        "log"
)

// After malformed payload cmd to evade IDS is generated encoding via various will be declared in this module

func encodeGzip(s string) {
        var buf bytes.Buffer
        gz := gzip.NewWriter(&buf)
        if _, err := gz.Write([]byte(s)); err != nil {
                log.Fatal(err)
        }
        if err := gz.Close(); err != nil {
                log.Fatal(err)
        }
   return b.Bytes()
}
