package main

import (
	"fmt"
	"os"

	"github.com/Zipeer2/count"
)

func main() {
	counter := count.NewCounter(os.Stdin)
	lines := counter.CountLines()
	fmt.Println(lines)
}
