package string

import (
	"strings"
	"testing"
)

func TestEqualFold(t *testing.T) {

	fold := strings.EqualFold("hello", "hello")
	print(fold)
}
