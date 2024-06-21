package count

import (
	"bufio"
	"io"
)

type Counter struct {
	input io.Reader
}

func NewCounter(input io.Reader) *Counter {
	return &Counter{input: input}
}

func (c *Counter) CountLines() int {
	scanner := bufio.NewScanner(c.input)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines
}
