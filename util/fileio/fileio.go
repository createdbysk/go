// Package fileio implements some File I/O utilities.
package fileio

import (
	"bufio"
	"os"
)

// ReadLines returns the lines in a file as an array.
func ReadLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines, nil
}
