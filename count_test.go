package count_test

import (
	"bytes"
	"testing"

	"github.com/Zipeer2/count"
)

func TestCountLines(t *testing.T) {
	t.Parallel()
	input := "line1\nline2\nline3\n"
	buf := bytes.NewBufferString(input)
	counter := count.NewCounter(buf)
	got := counter.CountLines()
	want := 3
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
