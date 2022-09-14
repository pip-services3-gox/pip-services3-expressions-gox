package tokenizers

import (
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers"
	"github.com/pip-services3-gox/pip-services3-expressions-gox/tokenizers/generic"
)

type MustacheTokenizer struct {
	*tokenizers.AbstractTokenizer
	special      bool
	specialState tokenizers.ITokenizerState
}

func NewMustacheTokenizer() *MustacheTokenizer {
	c := &MustacheTokenizer{
		special: true,
	}
	c.AbstractTokenizer = tokenizers.InheritAbstractTokenizer(c)

	c.SetSymbolState(generic.NewGenericSymbolState())
	c.SymbolState().Add("{{", tokenizers.Symbol)
	c.SymbolState().Add("}}", tokenizers.Symbol)
	c.SymbolState().Add("{{{", tokenizers.Symbol)
	c.SymbolState().Add("}}}", tokenizers.Symbol)

	c.SetNumberState(nil)
	c.SetQuoteState(generic.NewGenericQuoteState())
	c.SetWhitespaceState(generic.NewGenericWhitespaceState())
	c.SetWordState(generic.NewGenericWordState())
	c.SetCommentState(nil)
	c.specialState = NewMustacheSpecialState()

	c.ClearCharacterStates()
	c.SetCharacterState(0x0000, 0x00ff, c.SymbolState())
	c.SetCharacterState(0x0000, ' ', c.WhitespaceState())

	c.SetCharacterState('a', 'z', c.WordState())
	c.SetCharacterState('A', 'Z', c.WordState())
	c.SetCharacterState('0', '9', c.WordState())
	c.SetCharacterState('_', '_', c.WordState())
	c.SetCharacterState(0x00c0, 0x00ff, c.WordState())
	c.SetCharacterState(0x0100, 0xfffe, c.WordState())

	c.SetCharacterState('"', '"', c.QuoteState())
	c.SetCharacterState('\'', '\'', c.QuoteState())

	c.SetSkipWhitespaces(true)
	c.SetSkipComments(true)
	c.SetSkipEof(true)

	return c
}

func (c *MustacheTokenizer) ReadNextToken() *tokenizers.Token {
	if c.Scanner == nil {
		return nil
	}

	// Check for initial state
	if c.NextTokenValue == nil && c.LastTokenType == tokenizers.Unknown {
		c.special = true
	}

	// Process quotes
	if c.special {
		token := c.specialState.NextToken(c.Scanner, c)
		if token != nil && token.Value() != "" {
			return token
		}
	}

	// Proces other tokens
	c.special = false
	token := c.AbstractTokenizer.ReadNextToken()
	// Switch to quote when '{{' or '{{{' symbols found
	if token != nil && (token.Value() == "}}" || token.Value() == "}}}") {
		c.special = true
	}
	return token
}
