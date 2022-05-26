package test_variants

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-expressions-gox/variants"
	"github.com/stretchr/testify/assert"
)

func TestVariants(t *testing.T) {
	a := variants.NewVariant(123)
	assert.Equal(t, variants.Integer, a.Type())
	assert.Equal(t, 123, a.AsInteger())
	assert.Equal(t, 123, a.AsObject())

	b := variants.NewVariant("xyz")
	assert.Equal(t, variants.String, b.Type())
	assert.Equal(t, "xyz", b.AsString())
	assert.Equal(t, "xyz", b.AsObject())
}
