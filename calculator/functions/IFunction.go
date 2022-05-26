package functions

import "github.com/pip-services3-gox/pip-services3-expressions-gox/variants"

// IFunction вefines an interface for expression function.
// </summary>
type IFunction interface {
	// Name The function name.
	Name() string

	// Calculate the function calculation method.
	// Parameters:
	//		- parameters: A list with function parameters<
	//		- variantOperations: Variants operations manager.
	Calculate(parameters []*variants.Variant,
		variantOperations variants.IVariantOperations) (*variants.Variant, error)
}
