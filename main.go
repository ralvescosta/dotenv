package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	Configure(".env")
}

func Configure(path string) {
	scanner := bufio.NewScanner(readFile(path))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return file
}
