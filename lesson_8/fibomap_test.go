package fibomap

import (
	"fibomap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibomap(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  int
	}{
		{name: "1", input: 1, want: 0},
		{name: "2", input: 2, want: 1},
		{name: "10", input: 10, want: 34},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fibomap.FiboMap(tc.input))
		})
	}
}
