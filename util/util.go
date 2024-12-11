package util

import "os"

func ReadFile(filename string) (string, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return ``, err
	}
	return string(buffer), nil
}
