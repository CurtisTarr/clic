package tokenizer

// Literal constant for the literal Token type
const Literal string = "literal"

// Operator constant for the operator Token type
const Operator string = "operator"

// OpeningBrace constant for the openingBrace Token type
const OpeningBrace string = "openingBrace"

// ClosingBrace constant for the closingBrace Token type
const ClosingBrace string = "closingBrace"

// Token object that store the tokenType and its value
type Token struct {
	tokenType string
	value     string
}

// NewToken constructor for Token object
func NewToken(tokenType string, value string) *Token {
	token := new(Token)
	token.tokenType = tokenType
	token.value = value
	return token
}
