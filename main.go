package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	saveDir := ".save"

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		errExit(1, err.Error())
	}

	t := time.Now()
	formattedDate := t.Format("20060102T150405")

	targetPath := filepath.Join("Draft.md")
	targetExtension := filepath.Ext(targetPath)
	newFileBaseName := formattedDate + "--" + "draft" + targetExtension
	newFilePath := filepath.Join(saveDir, newFileBaseName)

	if err := os.Link(targetPath, newFilePath); err  != nil {
		errExit(1, err.Error())
	}

	fmt.Println(newFileBaseName)
}

func errExit(status int, format string, a ...any) {
	fmt.Fprintf(os.Stderr, "save error: " + format + "\n", a...)
	os.Exit(status)
}
