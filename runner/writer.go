package runner

import "os"

type Writer struct {
}

func (w Writer) WriteToFile(filePath string, text string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
