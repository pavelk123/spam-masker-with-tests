package maskerspam

import (
	"bufio"
	"fmt"
	"os"
)

const defaultPath string = "./test/output.txt"

// FilePresenter - for data presenter unit.
type FilePresenter struct {
	outputFile string
}

// NewFilePresenter is constructor of FilePresenter
// If path for output file is empty, then output file will be default
func NewFilePresenter(outputFile string) *FilePresenter {
	if outputFile == "" {
		outputFile = defaultPath
	}

	return &FilePresenter{outputFile: outputFile}
}

func (fp *FilePresenter) present(data []string) error {
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	writer := bufio.NewWriter(file)

	for _, str := range data {
		if _, err = writer.WriteString(str); err != nil {
			return fmt.Errorf("write string: %w", err)
		}
	}

	if err = writer.Flush(); err != nil {
		return fmt.Errorf("writer flush: %w", err)
	}

	if err = file.Close(); err != nil {
		return fmt.Errorf("file close: %w", err)
	}

	return nil
}
