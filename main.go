package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const commentComa = "#"

func main() {}

func Configure(path string) error {
	file, err := readFile(path)
	if err != nil {
		return errors.New("error while opening env file")
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

		index := strings.Index(line, commentComa)
		if index != -1 {
			lineParsed := strings.Split(line, "#")
			if len(lineParsed[0]) < 1 {
				continue
			}

			v := strings.Split(lineParsed[0], "=")
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
		return errors.New("error while reading env file")
	}

	return nil
}

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
