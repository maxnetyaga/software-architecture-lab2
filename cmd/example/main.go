package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/maxnetyaga/software-architecture-lab2"
)

func main() {
	expressionFlag := flag.String("e", "", "Expression for calculation")
	fileFlag := flag.String("f", "", "File with expression for calculation")
	outputFlag := flag.String("o", "", "File for recording the result")

	flag.Parse()

	if *expressionFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: cannot use both -e and -f flags at the same time")
		os.Exit(1)
	}

	var input io.Reader
	if *expressionFlag != "" {
		input = strings.NewReader(*expressionFlag)
	} else if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input expression specified (use -e or -f)")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to create file for result: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
