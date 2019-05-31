package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mholt/archiver"
)

func downloadHTTPFile(httpPath, localPath string) error {
	// download the file with the list of SP public agents

	r, err := http.Get(httpPath)
	dest, err := os.Create(localPath)
	defer dest.Close()
	_, err = io.Copy(dest, r.Body)
	if err != nil {
		return fmt.Errorf("could not download HTTPfile: %v", err)
	}
	return nil
}

func createFolderIfDoesntExist(expressions ...string) {
	for _, expr := range expressions {
		if _, err := os.Stat(expr); os.IsNotExist(err) {
			os.Mkdir(expr, 0777)
		}
	}
}

func downloadPublicAgentsFile(forceDownload bool) error {
	var errorsToReturn []error
	createFolderIfDoesntExist(downloadFolder, extractFolder)

	if _, err := os.Stat(downloadedFilePath); os.IsNotExist(err) {
		err = downloadHTTPFile(httpPath, downloadedFilePath)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
	} else if forceDownload {
		err := os.Remove(downloadedFilePath)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
		err = downloadHTTPFile(httpPath, downloadedFilePath)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
	}

	if _, err := os.Stat(extractedFilePath); os.IsNotExist(err) {
		err := archiver.Unarchive(downloadedFilePath, extractFolder)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
	} else if forceDownload {
		err := os.Remove(extractedFilePath)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
		err = archiver.Unarchive(downloadedFilePath, extractFolder)
		if err != nil {
			errorsToReturn = append(errorsToReturn, err)
		}
	}
	if errorsToReturn != nil {
		return fmt.Errorf("Could not download public agents file: %v", errorsToReturn)
	}
	return nil
}

func readStringInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	var choice string
	fmt.Print("\n", message)
	choice, _ = reader.ReadString('\n')
	return strings.Replace(choice, "\n", "", -1)
}
