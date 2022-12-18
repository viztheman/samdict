package cmu

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

// Downloads a given url to the specified path.
func Download(addr string, path string) error {
	_, err := url.Parse(addr)
	if err != nil {
		return err
	}

	resp, err := http.Get(addr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
