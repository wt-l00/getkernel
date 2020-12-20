package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeVersionID(t *testing.T) {
	versionIDTests := []struct {
		actual   string
		expected string
	}{
		{"5.8.7", "5.8.7-050807"},
		{"5.1-rc1", "5.1.0-050100rc1"},
		{"5.1", "5.1.0-050100"},
	}
	for _, versionID := range versionIDTests {
		assert.Equal(t, versionID.expected, makeVersionID(versionID.actual))
	}
}
