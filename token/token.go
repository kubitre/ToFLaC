package token

/*Token - токен для парса*/
type Token struct {
	Value         string
	Type          int
	StartPosition int
	EndPosition   int
	Line          int
	Column        int
	Action        int // действие 0 - вставить, 1- удалить, 2 - заменить
	Position      int // 0 - перед, 1 - после 2 - на место -1 не использовать
	Token         int // токен, который вставить, 0 - int, 1 - float, 2 - space, 3 - rand_id, 4 - endstatement, 5 - Pointer -1 - не использовать 6 - comma
	Index         int // индекс распознаноого токена
}

/*New - создание нового токена*/
func (tok *Token) New(tp int, val string, start, end int, line, col int) {
	tok.Value = val
	tok.Type = tp
	tok.StartPosition = start
	tok.EndPosition = end
	tok.Line = line
	tok.Column = col
}

/*Equals - проверка на равенство токенов*/
func (tok *Token) Equals(toks Token) bool {
	if tok.Index == toks.Index {
		return true
	} else {
		return false
	}
}
