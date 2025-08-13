package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		errExit(2, "usage: %s <file...>", os.Args[0])
	}
	repoDir := ".save"

	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102T150405")

	targetPath := os.Args[1]
	targetExtension := filepath.Ext(targetPath)
	targetNoExtension := strings.TrimSuffix(targetPath, targetExtension)

	newFileBaseName := formattedTime + "--" + strings.ToLower(targetNoExtension) + targetExtension
	newFilePath := filepath.Join(repoDir, newFileBaseName)

	body, err := os.ReadFile(targetPath)
	if err != nil {
		errExit(1, err.Error())
	}

	// directory to save it to
	if err := os.MkdirAll(repoDir, 0755); err != nil {
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

	fmt.Println("saved", newFileBaseName, "to", repoDir)
}

func errExit(status int, format string, a ...any) {
	fmt.Fprintf(os.Stderr, "save error: "+format+"\n", a...)
	os.Exit(status)
}
