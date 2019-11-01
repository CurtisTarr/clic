package tokenizer

const Literal string = "literal"
const Operator string = "operator"
const OpeningBrace string = "openingBrace"
const ClosingBrace string = "closingBrace"

var assoc = map[string] string{
	"^" : "right",
	"*" : "left",
	"/" : "left",
	"+" : "left",
	"-" : "left",
}

var prec = map[string] int{
	"^" : 4,
	"*" : 3,
	"/" : 3,
	"+" : 2,
	"-" : 2,
}

type Token struct {
	TokenType string
	Value     string
}

func NewToken(tokenType string, value string) *Token {
	token := new(Token)
	token.TokenType = tokenType
	token.Value = value
	return token
}

func GetAssoc(t *Token) string {
	return assoc[t.Value]
}

func GetPrec(t *Token) int {
	return prec[t.Value]
}