package error_

import (
	"tflac_cw/token"
)

/*ErrorModel - ошибка*/
type ErrorModel struct {
	Token   token.Token // токен, в котором произошла ошибка
	Message string      // сообщеине ошибки
}
