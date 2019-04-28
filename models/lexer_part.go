package models

import (
	"tflac_cw/error_"
	"tflac_cw/token"
)

type LexerPart struct {
	Tokens []token.Token
	Errors []error_.ErrorModel
}
