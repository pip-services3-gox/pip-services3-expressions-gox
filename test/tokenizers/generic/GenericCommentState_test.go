package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericCommentStateNextToken(t *testing.T) {
	state := generic.NewGenericCommentState()

	scanner := io.NewStringScanner("# Comment \r# Comment ")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, "# Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())

	scanner = io.NewStringScanner("# Comment \n# Comment ")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "# Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())
}
