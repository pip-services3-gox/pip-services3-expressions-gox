package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestSymbolRootNodeNextToken(t *testing.T) {
	node := generic.NewSymbolRootNode()
	node.Add("<", tokenizers.Symbol)
	node.Add("<<", tokenizers.Symbol)
	node.Add("<>", tokenizers.Symbol)

	scanner := io.NewStringScanner("<A<<<>")

	token := node.NextToken(scanner)
	assert.Equal(t, "<", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "A", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "<<", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "<>", token.Value())
}
