package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	poca "github.com/5aradise/polish-calculator"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression")
	outputFile      = flag.String("o", "", "File to write results")

	ErrFlagMissing  = errors.New("error: -e or -f must be specified")
	ErrFlagConflict = errors.New("error: cannot use -e and -f together")
)

func checkFlags(e, f string) error {
	if e == "" && f == "" {
		return ErrFlagMissing
	}

	if e != "" && f != "" {
		return ErrFlagConflict
	}

	return nil
}

func GetInput(inputExpression, inputFile string) (io.Reader, *os.File, error) {
	var input io.Reader
	var file *os.File
	var err error

	if inputExpression != "" {
		input = strings.NewReader(inputExpression)
	} else if inputFile != "" {
		file, err = os.Open(inputFile)

		if err != nil {
			return nil, nil, err
		}

		input = file
	}

	return input, file, nil
}

func GetOutput(destination string) (io.Writer, error) {
	var output io.Writer

	if destination == "" {
		output = os.Stdout
	} else {
		file, err := os.Create(destination)

		if err != nil {
			return nil, err
		}

		output = file
	}

	return output, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	err := checkFlags(*inputExpression, *inputFile)
	checkError(err)

	input, file, err := GetInput(*inputExpression, *inputFile)
	checkError(err)

	output, err := GetOutput(*outputFile)
	checkError(err)

	handler := &poca.ComputeHandler{Input: input, Output: output}
	checkError(handler.Compute())

	if file != nil {
		file.Close()
	}
}
