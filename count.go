package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input  io.Reader
	Output io.Writer
}

type option func(*counter) error

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}

func (c *counter) PrintLines() {
	fmt.Fprintln(c.Output, c.Lines())
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.Input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.Output = output
		return nil
	}
}

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Lines())
}
