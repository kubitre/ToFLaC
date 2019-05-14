package warning_

import (
	"tflac_cw/token"
)

/*WarningModel - модель предупреждения*/
type WarningModel struct {
	Token   token.Token
	Message string
}
