package gzep

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		log.Fatal(err)
		return false
	}
	if os.IsNotExist(err) {
		log.Fatal("Not Exists")
		return false
	}
	return true
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// With a little help from Phind
func gzepFileExclude(file string, patterns []byte) (int, error) {
	patCount := int(0)
	f, err := os.Open(file)
	checkError(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		exclude := false
		for _, pattern := range patterns {
			if bytes.Contains(line, []byte{pattern}) {
				exclude = true
				patCount++
				break
			}
		}
		if !exclude {
			fmt.Println(string(line))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return patCount, nil
}

// peter's grep modified from https://stackoverflow.com/questions/26709971/could-this-be-more-efficient-in-go
func gzepFile(file string, patterns []byte) (map[int]string, error) {
	patCount := int(0)
	artifacts := make(map[int]string)
	builder := strings.Builder{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i := 0; i <= len(patterns)-1; i++ {
			if bytes.Contains(scanner.Bytes(), []byte{patterns[i]}) {
				patCount++
				builder.WriteString(string(patterns[i]))
				artifacts[patCount] = builder.String()
				builder.Reset()
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return artifacts, nil
}

// Change the map to be patternCount, lineNumber and the pattern
func main() {

	var countFlag, fileFlag, helpFlag, excludeFlag string
	// var regexFlag, recursiveFlag string

	//	flag.StringVar(&recursiveFlag, "-r", "", "Path to a Diretory")
	//	flag.StringVar(&regexFlag, "-e", "", "") At some point
	// 	caseInsenstiveFlag, version flag

	flag.StringVar(&helpFlag, "-h", "help", "")
	flag.StringVar(&fileFlag, "-f", "", "File path to a file")
	flag.StringVar(&excludeFlag, "-x", "", "")
	flag.StringVar(&countFlag, "-c", "", "")
	flag.Parse()

	arg := flag.Args()
	argLen := len(arg)

	if argLen < 1 || argLen > 1 {
		fmt.Printf("Incorrect number of patterns provided to stdin: %d", argLen)
		flag.Usage()
		os.Exit(1)
	}

	if flag.Parsed() != true {
		fmt.Printf("Incorrect number of flags provided to stdin")
		flag.Usage()
		os.Exit(1)
	}

	if fileFlag == "" {
		fmt.Printf("A file path to a file must be provided with -f")
		flag.Usage()
		os.Exit(1)
	}

	if checkFileExists(fileFlag) != true {
		fmt.Printf("File: %s does not exist!!!", fileFlag)
		os.Exit(1)
	}
	resultMap := make(map[int]string)
	resultCount := 0
	convArgStr := strings.Join(arg, "")
	convArgBytes := []byte(convArgStr)
	if excludeFlag != "" {
		resultCount, _ = gzepFileExclude(fileFlag, convArgBytes)
	} else {
		resultMap, _ = gzepFile(fileFlag, convArgBytes)
		resultCount = len(resultMap)
	}

	if countFlag != "" {
		fmt.Printf("The total pattern count for: %s was %d", convArgStr, resultCount)
	}

	os.Exit(1)
}
