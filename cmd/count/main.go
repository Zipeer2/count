package main

import (
	"bytes"
	"fmt"

	"github.com/Zipeer2/count"
)

func main() {
	input := bytes.NewBufferString("line1\nline2\nline3\n")
	output := new(bytes.Buffer)

	c := count.NewCounter(count.WithInput(input), count.WithOutput(output))
	lines := c.Lines()
	fmt.Printf("Lines: %d\n", lines)
}
