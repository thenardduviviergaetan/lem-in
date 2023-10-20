package utils

import (
	"os"
)

func Readfile(path string) []byte {
	file, err := os.ReadFile(path)
	Check_err(err)
	return file
}
