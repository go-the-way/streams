package streams

import (
	"fmt"
	"testing"
)

func TestSkip(t *testing.T) {
	t.Log(Map(func(in int) string {
		return fmt.Sprintf("%d", in)
	}, Skip(1, 1, 2, 3, 4)...))
}
