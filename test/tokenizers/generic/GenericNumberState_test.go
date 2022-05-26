package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericNumberStateNextToken(t *testing.T) {
	state := generic.NewGenericNumberState()

	scanner := io.NewStringScanner("ABC")
	failed := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				failed = true
			}
		}()
		state.NextToken(scanner, nil)
	}()
	assert.True(t, failed)

	scanner = io.NewStringScanner("123#")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, "123", token.Value())
	assert.Equal(t, tokenizers.Integer, token.Type())

	scanner = io.NewStringScanner("-123#")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "-123", token.Value())
	assert.Equal(t, tokenizers.Integer, token.Type())

	scanner = io.NewStringScanner("123.#")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "123.", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())

	scanner = io.NewStringScanner("123.456#")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "123.456", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())

	scanner = io.NewStringScanner("-123.456#")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "-123.456", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())
}
