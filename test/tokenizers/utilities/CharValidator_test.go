package test_utilities

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/utilities"
	"github.com/stretchr/testify/assert"
)

func TestCharValidatorIsEof(t *testing.T) {
	assert.True(t, utilities.CharValidator.IsEof(-1))
	assert.False(t, utilities.CharValidator.IsEof('A'))
}

func TestCharValidatorIsEol(t *testing.T) {
	assert.True(t, utilities.CharValidator.IsEol(10))
	assert.True(t, utilities.CharValidator.IsEol(13))
	assert.False(t, utilities.CharValidator.IsEof('A'))
}

func TestCharValidatorIsDigit(t *testing.T) {
	assert.True(t, utilities.CharValidator.IsDigit('0'))
	assert.True(t, utilities.CharValidator.IsDigit('7'))
	assert.True(t, utilities.CharValidator.IsDigit('9'))
	assert.False(t, utilities.CharValidator.IsDigit('A'))
}
