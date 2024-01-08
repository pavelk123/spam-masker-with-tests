package main

import (
	"log/slog"
	"os"

	"github.com/pavelk123/spam-masker-with-tests/maskerspam"
)

const (
	zeroArgs = iota
	oneArg
	twoArgs
)

func main() {
	// "./test/e.txt"
	// "./test/output.txt"
	var inputFile, outputFile string

	paths := os.Args[1:]

	switch len(paths) {
	case zeroArgs:
		slog.Error("Params not provided")

		return

	case oneArg:
		inputFile = paths[0]
		outputFile = ""

	case twoArgs:
		inputFile = paths[0]
		outputFile = paths[1]

	default:
		slog.Error("Toooooo much params")

		return
	}

	fileProducer := maskerspam.NewFileProducer(inputFile)
	filePresenter := maskerspam.NewFilePresenter(outputFile)
	service := maskerspam.NewService(fileProducer, filePresenter)

	if err := service.Run(); err != nil {
		slog.Error(err.Error())

		return
	}

	slog.Info("Complete")
}
