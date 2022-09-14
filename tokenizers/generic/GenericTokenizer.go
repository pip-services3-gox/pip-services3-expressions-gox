package generic

import "github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"

// GenericTokenizer implements a default tokenizer class.
type GenericTokenizer struct {
	*tokenizers.AbstractTokenizer
}

func NewGenericTokenizer() *GenericTokenizer {
	c := &GenericTokenizer{}
	c.AbstractTokenizer = tokenizers.InheritAbstractTokenizer(c)

	c.SetSymbolState(NewGenericSymbolState())
	c.SymbolState().Add("<>", tokenizers.Symbol)
	c.SymbolState().Add("<=", tokenizers.Symbol)
	c.SymbolState().Add(">=", tokenizers.Symbol)

	c.SetNumberState(NewGenericNumberState())
	c.SetQuoteState(NewGenericQuoteState())
	c.SetWhitespaceState(NewGenericWhitespaceState())
	c.SetWordState(NewGenericWordState())
	c.SetCommentState(NewGenericCommentState())

	c.ClearCharacterStates()
	c.SetCharacterState(0x0000, 0x00ff, c.SymbolState())
	c.SetCharacterState(0x0000, ' ', c.WhitespaceState())

	c.SetCharacterState('a', 'z', c.WordState())
	c.SetCharacterState('A', 'Z', c.WordState())
	c.SetCharacterState(0x00c0, 0x00ff, c.WordState())
	c.SetCharacterState(0x0100, 0xffff, c.WordState())

	c.SetCharacterState('-', '-', c.NumberState())
	c.SetCharacterState('0', '9', c.NumberState())
	c.SetCharacterState('.', '.', c.NumberState())

	c.SetCharacterState('"', '"', c.QuoteState())
	c.SetCharacterState('\'', '\'', c.QuoteState())

	c.SetCharacterState('#', '#', c.CommentState())

	return c
}
