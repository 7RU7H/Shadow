package ninja

import (
        "context"
        "crypto/rand"
        "crypto/rsa"
        "crypto/tls"
        "crypto/x509"
        "encoding/pem"
        "fmt"
        "io"
        "log"
        "math/big"

        "github.com/lucas-clemente/quic-go"
)

func Server() error {
        listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
        if err != nil {
                return err
        }
        conn, err := listener.Accept(context.Background())
        if err != nil {
                return err
        }
        stream, err := conn.AcceptStream(context.Background())
        if err != nil {
                panic(err)
        }
        // Echo through the loggingWriter
        _, err = io.Copy(loggingWriter{stream}, stream)
        return err
}

// Logging functionality here for test, depending on build variation 
// Would like to have logging
// A wrapper for io.Writer that also logs the message.
type loggingWriter struct{ io.Writer }

func (w loggingWriter) Write(b []byte) (int, error) {
        fmt.Printf("Server: Got '%s'\n", string(b))
        return w.Writer.Write(b)
}

