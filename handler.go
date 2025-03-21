package lab2

import (
	"io"
	"io/ioutil"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.

// Обробляє вхідні дані та записує результат
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Читає вхідні дані, викликає функцію з implementation.go та записує результат
func (ch *ComputeHandler) Compute() error {
	data, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := strings.TrimSpace(string(data))
	result, err := PostfixToInfix(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result + "\n"))
	return err
}
