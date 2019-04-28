package error_

import (
	"tflac_cw/token"
)

/*ErrorModel - ошибка*/
type ErrorModel struct {
	Type  int         // тип ошибки (0 - лексическая, синтаксическая)
	Token token.Token // токен, в котором произошла ошибка
}
