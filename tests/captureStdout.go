package tests

import (
	"bytes"
	"io"
	"os"
)

func CaptureStdout(fn func() error) (string, error) {
	old := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	os.Stdout = w
	defer func() {
		os.Stdout = old
	}()

	err = fn() //can be ommited, moslty I don't react on errors in tests

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String(), nil
}
