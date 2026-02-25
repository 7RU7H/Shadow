package internal

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReplicateSelf(dstFilePath string) error {
	path, err := os.Getwd()
	if err != nil {
		ErrorHandling(err, 2)
		return err
	}

	srcBin, err := os.Open(path + os.Args[0])
	if err != nil {
		log.Fatal("Failed to open source file:", err)
	}
	defer srcBin.Close()

	stat, err := srcBin.Stat()
	if err != nil {

	}

	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		ErrorHandling(err, 2)
		return err
	}
	defer dstFile.Close()

	buffer := make([]byte, stat.Size())
	totalBytes := int64(0)
	written := 0
	for {
		n, err := srcBin.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			ErrorHandling(err, 2)
			return err
		}

		written, err = dstFile.Write(buffer[:n])
		if err != nil {
			ErrorHandling(err, 2)
			return err
		}

		totalBytes += int64(written)
	}
	if totalBytes != stat.Size() {
		err := fmt.Errorf("Failed to write %v bytes to dstFilePath", totalBytes-int64(written))
		ErrorHandling(err, 2)
		return err
	}

	return nil
}
