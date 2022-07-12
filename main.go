package ninjashell

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

	"ninjashell/gc.go"
	"ninjashell/tcp.go"
	"ninjashell/udp.go"
	"ninjashell/filetransfer.go"
	"ninjashell/encrypt.go"
)

func main() {
	os.Exit(gc.CLI(os.Args[1:]))
}


