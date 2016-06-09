package api

import (
	"bytes"
	"io"
	"io/ioutil"
)

type bodyReader struct {
	*bytes.Buffer
}

func (b bodyReader) Close() error {
	return nil
}

func readAndReturn(b io.ReadCloser) (io.ReadCloser, string) {
	bts, _ := ioutil.ReadAll(b)
	return bodyReader{bytes.NewBuffer(bts)}, string(bts)
}
