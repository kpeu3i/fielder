package main

import (
	"fmt"
	"io/fs"
	"os"
)

func writeToFile(data []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("could not open file %q: %v", filename, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}
