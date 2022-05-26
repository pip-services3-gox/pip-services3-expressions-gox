package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericSymbolStateNextToken(t *testing.T) {
	state := generic.NewGenericSymbolState()
	state.Add("<", tokenizers.Symbol)
	state.Add("<<", tokenizers.Symbol)
	state.Add("<>", tokenizers.Symbol)

	scanner := io.NewStringScanner("<A<<<>")

	token := state.NextToken(scanner, nil)
	assert.Equal(t, "<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "A", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "<<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "<>", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())
}
