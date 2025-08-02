package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// directory to save it to
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

	fmt.Println(newFileBaseName)
}

func errExit(status int, format string, a ...any) {
	fmt.Fprintf(os.Stderr, "save error: "+format+"\n", a...)
	os.Exit(status)
}
