package poca

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)

	str, _ := reader.ReadString('\n')

	res, err := PostfixToLisp(str)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(res))
	if err != nil {
		return err
	}

	return nil
}
