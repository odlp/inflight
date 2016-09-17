package runner

import "os"

type Writer struct {
}

func (w Writer) WriteToFile(filePath string, text string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	gracefulExitIfError(err)

	defer f.Close()

	_, err = f.WriteString(text)
	gracefulExitIfError(err)
}
