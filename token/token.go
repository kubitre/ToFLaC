package token

/*Token - токен для парса*/
type Token struct {
	Value         string
	Type          int
	StartPosition int
	EndPosition   int
}

/*New - создание нового токена*/
func (tok *Token) New(tp int, val string, start, end int) {
	tok.Value = val
	tok.Type = tp
	tok.StartPosition = start
	tok.EndPosition = end
}
