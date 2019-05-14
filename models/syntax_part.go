package models

import (
	"tflac_cw/error_"
	"tflac_cw/token"
	"tflac_cw/warning_"
)

/*SyntaxPart - результат синтаксического анализатора*/
type SyntaxPart struct {
	Errors   []error_.ErrorModel
	Warnings []warning_.WarningModel
	Repair   []token.Token `json:"repair"`
}
