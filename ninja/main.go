package ninja

import (
	"bufio"
	"bytes"
	"crypto/sha3"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"client.go"
	"listener.go"
	"tls.go"
	"handle.go"
)

func main() {
	os.Exit(handle.CLI(os.Args[1:]))
}


