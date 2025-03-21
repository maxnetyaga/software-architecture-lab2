package lab2

import (
	"io"
	"io/ioutil"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

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
