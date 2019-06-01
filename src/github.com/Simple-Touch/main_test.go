package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestWithoutFileName(t *testing.T) {
	const expectedResult string = "Please provide a file to create"
	out, err := exec.Command("go", "run", "main.go").Output()
	result := strings.Trim(string(out), "\n")
	if err != nil {
		log.Fatalf("Expected to get message 'Please provide a file to create', but got %s", err.Error())
	} else if result != expectedResult {
		log.Fatalf("Expected to get message 'Please provide a file to create', but got %s", result)
	}
}

func TestWithFile(t *testing.T) {
	_, err := exec.Command("go", "run", "main.go", "test.txt").Output()
	if err != nil {
		log.Fatalf("Expected to create file")
	}
	cwd, _ := os.Getwd()
	filePath := cwd + string(os.PathSeparator) + "test.txt"
	_, err = os.Stat(filePath)
	if err != nil {
		log.Fatal("Failed to create the file")
	}
	os.Remove(filePath)
}

func TestWithMultipleFiles(t *testing.T) {
	cmdArgs := []string{"run", "main.go", "test1.txt", "test2.log"}

	_, err := exec.Command("go", cmdArgs...).Output()
	if err != nil {
		log.Fatalf("Expected to create files")
	}
	cwd, _ := os.Getwd()
	for _, v := range cmdArgs[2:] {
		filePath := cwd + string(os.PathSeparator) + v
		_, err = os.Stat(filePath)
		if err != nil {
			log.Fatal("Failed to create the file ", v)
		}
		os.Remove(filePath)
	}
}
