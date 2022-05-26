package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericWhitespaceStateNextToken(t *testing.T) {
	state := generic.NewGenericWhitespaceState()

	scanner := io.NewStringScanner(" \t\n\r ")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, " \t\n\r ", token.Value())
	assert.Equal(t, tokenizers.Whitespace, token.Type())
}
