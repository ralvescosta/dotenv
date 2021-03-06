package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const comment = "#"
const separator = "="

func Configure(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("[DotEnv::Configure] - error while opening the file: %v", err)
	}
	defer func() error {
		if err = file.Close(); err != nil {
			return fmt.Errorf("[DotEnv::Configure] - error while close the file: %v", err)
		}
		return nil
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.ReplaceAll(line, "\"", "")
		if line == "" {
			continue
		}
		index := strings.Index(line, comment)
		if index != -1 {
			parsed := strings.Split(line, comment)
			if len(parsed[0]) < 1 {
				continue
			}

			v := strings.Split(parsed[0], separator)
			if len(v) == 2 && v[1] != "" && v[1] != " " {
				os.Setenv(strings.TrimSpace(v[0]), strings.TrimSpace(v[1]))
			}
			continue
		}
		v := strings.Split(line, separator)
		if v[1] != "" && v[1] != " " {
			os.Setenv(strings.TrimSpace(v[0]), strings.TrimSpace(v[1]))
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("[DotEnv::Configure] - error while reading env file: %v", err)
	}

	return nil
}
