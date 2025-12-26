package src

import (
	"errors"
	"os"
	fp "path/filepath"
	"strings"
)

func ValidInput() error {
	if len(os.Args) < 2 {
		return errors.New(`not enough arguments
Usage: ./scop foo.obj`)
	} else if len(os.Args) > 2 {
		return errors.New(`too many arguments
Usage: ./scop foo.obj`)
	}

	return validFile(os.Args[1])
}

func validFile(filename string) error {
	if strings.ToLower(fp.Ext(filename)) != ".obj" || len(fp.Base(filename)) <= 4 {
		return errors.New("file must be a '.obj'")
	}

	file, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if file.IsDir() {
		return errors.New("provided path is actually a directory")
	} else if (file.Mode().Perm() & 0400) != 0400 {
		return errors.New("no permission to read that file")
	}

	return nil
}
