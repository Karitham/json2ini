package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/go-json-experiment/json"
)

type Json2Ini map[string]map[string]string

func Transform(r io.Reader, w io.Writer) error {
	var ini Json2Ini
	err := json.UnmarshalRead(r, &ini)
	if err != nil {
		return err
	}

	buf := bufio.NewWriter(w)

	for section, values := range ini {
		// write section header
		buf.Write([]byte("["))
		buf.Write([]byte(section))
		buf.Write([]byte("]\n"))

		for key, value := range values {
			buf.Write([]byte(key))
			buf.Write([]byte("="))
			buf.Write([]byte(value))
			buf.Write([]byte("\n"))
		}
	}

	return buf.Flush()
}

func main() {
	err := Transform(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
