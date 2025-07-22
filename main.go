package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Opts struct {
	GZIP bool
}

func main() {
	var opts Opts
	opts.GZIP = false // TODO: parse cli

	repoDir := ".save"
	if err := os.MkdirAll(repoDir, 0755); err != nil {
		errExit(1, err.Error())
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102T150405")
	targetPath := filepath.Join("Draft.md")
	targetExtension := filepath.Ext(targetPath)
	newFileBaseName := formattedTime + "--" + "draft" + targetExtension
	newFilePath := filepath.Join(repoDir, newFileBaseName)

	body, err := os.ReadFile(targetPath)
	if err != nil {
		errExit(1, err.Error())
	}

	if opts.GZIP {
		newFileBaseName += ".gz"
		newFilePath += ".gz"

		var compressedBuffer bytes.Buffer
		w := gzip.NewWriter(&compressedBuffer)
		w.Write(body)
		w.Close()

		f, err := os.Create(newFilePath)
		if err != nil {
			errExit(1, err.Error())
		}
		defer f.Close()

		n, err := f.Write(compressedBuffer.Bytes())
		if err != nil {
			errExit(1, err.Error())
		}
		_ = n
	} else {
		f, err := os.Create(newFilePath)
		if err != nil {
			errExit(1, err.Error())
		}
		defer f.Close()

		n, err := f.Write(body)
		if err != nil {
			errExit(1, err.Error())
		}
		_ = n
	}

	fmt.Printf("Saved %v at %v\n", targetPath, formattedTime)
}

func errExit(status int, format string, a ...any) {
	fmt.Fprintf(os.Stderr, "save error: "+format+"\n", a...)
	os.Exit(status)
}
