package variables

import "github.com/pip-services3-gox/pip-services3-expressions-gox/variants"

// IVariable defines a variable interface.
type IVariable interface {
	// Name the variable name.
	Name() string

	// Value gets the variable value.
	Value() *variants.Variant

	// SetValue sets the variable value.
	SetValue(value *variants.Variant)
}
