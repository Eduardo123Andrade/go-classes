package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func splitSemiColon(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil // Stop scanning
	}

	if i := bytes.IndexByte(data, ';'); i >= 0 {
		return i + 1, data[:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func main() {
	valueToScann := []byte("Hello;World")
	scanner := bufio.NewScanner(bytes.NewReader(valueToScann))
	scanner.Split(splitSemiColon)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

}
