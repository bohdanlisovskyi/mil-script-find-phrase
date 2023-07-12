package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findPhraseInFile(filePath string, phrase string, newFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error occurred while closing the file: %+v\n", err)
		}
	}()

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := newFile.Close(); err != nil {
			fmt.Printf("Error occurred while closing the new file: %+v\n", err)
		}
	}()

	writer := bufio.NewWriter(newFile)
	scanner := bufio.NewScanner(file)
	block := ""

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 { // empty line
			if strings.Contains(block, phrase) {
				_, err := writer.WriteString(block + "\n")
				if err != nil {
					return err
				}
				err = writer.Flush()
				if err != nil {
					return err
				}
			}
			block = ""
		} else {
			block += line + "\n"
		}
	}

	// check last block
	if strings.Contains(block, phrase) {
		_, err := writer.WriteString(block + "\n")
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Please provide the source file path, the phrase to search for, and the destination file path. Example: script.exe input.txt \"my phrase\" output.txt")
		return
	}

	filePath := os.Args[1]
	phrase := os.Args[2]
	newFilePath := os.Args[3]

	err := findPhraseInFile(filePath, phrase, newFilePath)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
	}
}
