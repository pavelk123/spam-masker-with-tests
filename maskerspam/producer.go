package maskerspam

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// FileProducer - for data produce unit.
type FileProducer struct {
	inputFile string
}

// NewFileProducer is constructor of FileProducer
func NewFileProducer(inputFile string) *FileProducer {
	return &FileProducer{inputFile: inputFile}
}

func (f *FileProducer) produce() ([]string, error) {
	const fileMode uint = 0o666

	f.inputFile = strings.TrimSuffix(f.inputFile, "\n")

	f.inputFile = strings.TrimSuffix(f.inputFile, "\r")

	file, err := os.OpenFile(f.inputFile, os.O_RDONLY, fs.FileMode(fileMode))
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	var writer bytes.Buffer

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		if _, err = writer.WriteString(sc.Text()); err != nil {
			return nil, fmt.Errorf("write string: %w", err)
		}

		writer.WriteString("\n")
	}

	if err = file.Close(); err != nil {
		return nil, fmt.Errorf("file close: %w", err)
	}

	return []string{strings.TrimSpace(writer.String())}, nil
}
