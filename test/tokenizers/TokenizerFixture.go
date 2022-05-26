package test_tokenizers

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/stretchr/testify/assert"
)

// Checks is expected tokens matches actual tokens.
// Parameters:
//   - expectedTokens: An array with expected tokens.
//   - actualTokens: An array with actual tokens.
func AssertAreEqualsTokenLists(t *testing.T,
	expectedTokens []*tokenizers.Token, actualTokens []*tokenizers.Token) {

	assert.Equal(t, len(expectedTokens), len(actualTokens))

	for i := 0; i < len(expectedTokens); i++ {
		assert.Equal(t, expectedTokens[i].Type(), actualTokens[i].Type())
		assert.Equal(t, expectedTokens[i].Value(), actualTokens[i].Value())
	}
}
