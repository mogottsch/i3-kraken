package osutils

import (
	"bufio"
	"os"
)

const SESSIONIZER_RES_FILE = "/tmp/sessionizer-res"

func readFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var fileContent string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fileContent += scanner.Text()
	}

	return fileContent, nil
}

func ReadSessionizerRes() (string, error) {
	return readFile(SESSIONIZER_RES_FILE)
}
