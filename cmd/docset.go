package cmd
import (
  "path/filepath"
  "io/ioutil"
  "github.com/russross/blackfriday"
)

// Compile files into markdown and combine
func Compile(fileList []string, notice string) []byte {
  combinedInput := []byte{}
  for _, file := range fileList {
    input, _ := ioutil.ReadFile(file)
    combinedInput = append(combinedInput, input...)
  }

  html := blackfriday.Run(combinedInput)
  return append([]byte(notice), html...)
}

// Resolve interates through all files in all docs expanding globs and preserving uniqueness
// of the slice. This allows you to declare `DocConfiguration.Files` in the order you want and glob directories
// to fall back to alphabetical order.
func Resolve(fileList []string) []string {
	resolvedFiles := []string{}

	for _, fileName := range fileList {
		abs, _ := filepath.Abs(fileName)
		files, _ := filepath.Glob(abs)
		for _, file := range files {
			resolvedFiles = appendIfMissing(resolvedFiles, file)
		}
	}

	return resolvedFiles
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}
