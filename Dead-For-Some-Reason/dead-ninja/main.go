package ninja

import (
	"os"

	"handle.go"
)

func main() {
	os.Exit(handle.CLI(os.Args[1:]))
}
