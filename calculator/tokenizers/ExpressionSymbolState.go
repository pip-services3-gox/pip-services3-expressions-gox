package tokenizers

import (
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
)

// ExpressionSymbolState implements a symbol state object.
type ExpressionSymbolState struct {
	generic.GenericSymbolState
}

// NewExpressionSymbolState constructs an instance of this class.
func NewExpressionSymbolState() *ExpressionSymbolState {
	c := &ExpressionSymbolState{
		GenericSymbolState: *generic.NewGenericSymbolState(),
	}

	c.Add("<=", tokenizers.Symbol)
	c.Add(">=", tokenizers.Symbol)
	c.Add("<>", tokenizers.Symbol)
	c.Add("!=", tokenizers.Symbol)
	c.Add(">>", tokenizers.Symbol)
	c.Add("<<", tokenizers.Symbol)

	return c
}
