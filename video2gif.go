package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Set input and output file paths
	inputFile := "rat1.mp4"
	outputFile := "rat1.gif"

	// Set imaginary server URL
	serverURL := "http://localhost:9000"

	// Set imaginary conversion options
	options := "?format=gif"

	// Construct imaginary conversion URL
	conversionURL := fmt.Sprintf("%s/convert%s", serverURL, options)

	// Use cURL to send a POST request to imaginary server
	cmd := exec.Command("curl", "-X", "POST", "-F", fmt.Sprintf("image=@%s", inputFile), conversionURL)
	response, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// Save the converted file to disk
	err = saveFile(outputFile, response)
	if err != nil {
		log.Fatal(err)
	}
}

func saveFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
