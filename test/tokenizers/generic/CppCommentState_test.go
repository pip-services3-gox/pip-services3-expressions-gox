package test_generic

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/io"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestCppCommentStateNextToken(t *testing.T) {
	state := generic.NewCppCommentState()

	scanner := io.NewStringScanner("-- Comment \n Comment ")
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

	scanner = io.NewStringScanner("// Comment \n Comment ")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, "// Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())

	scanner = io.NewStringScanner("/* Comment \n Comment */#")
	token = state.NextToken(scanner, nil)
	assert.Equal(t, "/* Comment \n Comment */", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())
}
