package helper

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

/* Private Methods */
func getAbsPath(path string) string {
	fullPath, err := filepath.Abs(path)

	if err != nil {
		fmt.Println("error:", err)
	}

	return fullPath
}

func checkFileExist(path string) bool {
	path = getAbsPath(path)
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

/* Public methods */
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

func GetHomePath() string {
	path, _ := os.UserHomeDir()
	return path
}
