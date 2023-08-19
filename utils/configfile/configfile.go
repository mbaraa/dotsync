package configfile

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mbaraa/dotsync/config"
)

func GetValue(key string) (string, error) {
	values, err := readFile()
	if err != nil {
		return "", err
	}

	if value, ok := values[key]; ok {
		return value, nil
	}

	return "", errors.New("key was not found")
}

func SetValue(key, value string) error {
	values, err := readFile()
	if err != nil {
		return err
	}
	values[key] = value

	return writeToFile(values)
}

func readFile() (map[string]string, error) {
	file, err := openFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}
		keyValue := strings.Split(line, "=")
		values[keyValue[0]] = keyValue[1]
	}

	return values, nil
}

func writeToFile(values map[string]string) error {
	file, err := openFile()
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range values {
		if len(key) == 0 || len(value) == 0 {
			continue
		}
		fmt.Fprintf(file, "%s=%s\n", key, value)
	}

	return nil
}

func openFile() (file *os.File, err error) {
	file, err = os.OpenFile(config.ConfigFilePath, os.O_RDWR, 0755)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(config.ConfigFilePath)
		if err != nil {
			return
		}
	}

	return
}
