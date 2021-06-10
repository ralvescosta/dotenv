package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const commentCaractere = "#"

func main() {}

func Configure(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New(fmt.Sprintf("error while opening env file: %v", err))
	}
	defer func() error {
		if err = file.Close(); err != nil {
			return err
		}
		return nil
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		index := strings.Index(line, commentCaractere)
		if index != -1 {
			parsed := strings.Split(line, "#")
			if len(parsed[0]) < 1 {
				continue
			}

			v := strings.Split(parsed[0], "=")
			if len(v[1]) > 1 {
				os.Setenv(v[0], v[1])
			}
			continue
		}
		v := strings.Split(line, "=")
		if len(v[1]) > 1 {
			os.Setenv(v[0], v[1])
		}
	}

	err = scanner.Err()
	if err != nil {
		return errors.New(fmt.Sprintf("error while reading env file: %v", err))
	}

	return nil
}
