package models

import (
	"tflac_cw/error_"
	"tflac_cw/token"
)

/*LexerPart - часть лексического анализатора*/
type LexerPart struct {
	Tokens []token.Token
	Errors []error_.ErrorModel
}
