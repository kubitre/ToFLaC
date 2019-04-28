package models

import (
	"tflac_cw/error_"
)

type SyntaxPart struct {
	Errors []error_.ErrorModel
}
