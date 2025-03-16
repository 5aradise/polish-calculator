package poca

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualResult(t *testing.T) {
	output := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  strings.NewReader("2 2 +"),
		Output: output,
	}

	err := handler.Compute()

	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 2 2)", output.String())
	}
}

func TestSyntaxError(t *testing.T) {
	handler := ComputeHandler{
		Input:  strings.NewReader("faang"),
		Output: new(bytes.Buffer),
	}

	err := handler.Compute()

	assert.NotNil(t, err)
}
