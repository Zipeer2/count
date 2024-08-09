package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input  io.Reader
	Output io.Writer
}

type Option func(*Counter)

func NewCounter(opts ...Option) *Counter {
	c := &Counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}

func WithInput(input io.Reader) Option {
	return func(c *Counter) {
		c.Input = input
	}
}

func WithOutput(output io.Writer) Option {
	return func(c *Counter) {
		c.Output = output
	}
}

func Main() {

}
