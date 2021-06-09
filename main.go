package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	Configure(".env")
}

func Configure(path string) error {
	file, err := readFile(path)
	if err != nil {
		return errors.New("error while opening env file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return errors.New("error while reading env file")
	}

	return nil
}

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() error {
		if err = file.Close(); err != nil {
			return err
		}
		return nil
	}()

	return file, nil
}
