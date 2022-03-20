package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindBinary(command string) (string, error) {
	if command == "" {
		return "", errors.New("Please specify command")
	}
	if strings.Contains(command, "/") {
		return command, nil
	}

	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return "", errors.New("PATH is empty")
	}

	paths := filepath.SplitList(pathEnv)
	for _, p := range paths {
		t := filepath.Join(p, command)
		if _, err := os.Stat(t); err != nil {
			continue
		}
		return t, nil
	}
	return "", fmt.Errorf("Command %s not found", command)
}

func Main() error {
	if len(os.Args) < 2 {
		return errors.New("Please specify command")
	}
	bin := os.Args[1]
	path, err := FindBinary(bin)
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err.Error())
	}
}
