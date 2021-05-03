package main

import (
	"bytes"
	"io"

	"github.com/vincent-petithory/dataurl"
)

func dataURLReader(dataURLSrc string) (io.Reader, error) {
	dataURL, err := dataurl.DecodeString(dataURLSrc)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(dataURL.Data), nil
}

