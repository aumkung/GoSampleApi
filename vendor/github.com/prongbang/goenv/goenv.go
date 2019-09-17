package goenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Env map[string]string

func init() {
	if err := LoadEnv(); err != nil {
		fmt.Errorf("%v", err)
	}
}

func LoadEnv(filename ...string) error {
	if len(filename) == 0 {
		filename = []string{".env"}
	}

	for _, file := range filename {
		f, e := os.Open(file)
		if e != nil {
			return e
		}
		defer f.Close()

		env := parse(f)
		for key, val := range env {
			Setenv(key, val)
		}
	}
	return nil
}

func Setenv(key string, val string) {
	os.Setenv(key, val)
}

func parse(r io.Reader) Env {
	env := make(Env)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "=") {
			keyVal := strings.Split(text, "=")
			if len(keyVal) >= 2 {
				key := keyVal[0]
				val := keyVal[1]
				env[key] = val
			}
		}
	}

	return env
}
