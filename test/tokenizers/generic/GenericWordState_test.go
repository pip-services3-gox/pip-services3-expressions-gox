package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericWordStateNextToken(t *testing.T) {
	state := generic.NewGenericWordState()

	scanner := io.NewStringScanner("AB_CD=")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, "AB_CD", token.Value())
	assert.Equal(t, tokenizers.Word, token.Type())
}
