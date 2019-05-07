package helper

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

func getAbsPath(path string) string {
	fullPath, err := filepath.Abs(path)

	if err != nil {
		fmt.Println("error:", err)
	}

	return fullPath
}

func ReadLocalFile(path string) []byte {
	path = getAbsPath(path)
	b, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("error:", err)
	}

	return b
}

func WriteLocalFile(path string, content []byte) {
	path = getAbsPath(path)
	err := ioutil.WriteFile(path, content, 0755)

	if err != nil {
		fmt.Println("error:", err)
	}
}

func ReadAsHtml(path string) string {
	fileStream := string(ReadLocalFile(path))

	return "data:text/html," + url.PathEscape(fileStream)
}