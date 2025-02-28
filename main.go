package main

import (
	"github.com/jessesimpson36/zip-test/internal/archive"
	"fmt"
	"runtime/debug"
)

func createZipArchive(filesToArchive []string, outputPath string) error {
	if err := archive.ZipSource(filesToArchive, outputPath); err != nil {
		return fmt.Errorf("failed to create c8run package: %w\n%s", err, debug.Stack())
	}
	return nil
}

func main() {

	filesToArchive := []string{
		"test",
	}
	outputPath := "test.zip"
	if err := createZipArchive(filesToArchive, outputPath); err != nil {
		fmt.Println("createZip : %w", err)
	}
}
