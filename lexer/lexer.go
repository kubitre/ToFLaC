package lexer

import (
	"errors"
	"log"
	"tflac_cw/error_"
	"tflac_cw/token"
)

/*Lexer - лексический анализ введённого текста*/
type Lexer struct {
	AnalysingString string              // анализируемая строка
	Debug           bool                // вывод промежуточных логов и ошибок
	Tokens          []token.Token       // слайс токенов
	Errors          []error_.ErrorModel // слайс с ошибками на этапе лексического анализа
	Length          int                 // длина
	Position        int                 // текущая позиция в вводимой строке
	LexerAutomat    *AutomatState       // конечный автомат
}

const (
	prefix = "[LEXER]: "
)

/*New - инициализация нового лексера*/
func (lex *Lexer) New(debug bool) {
	autm := &AutomatState{}
	lex.LexerAutomat = autm.NewAutomat()
	lex.Debug = debug
}

/*Input - добавление новой строки для анализа*/
func (lex *Lexer) Input(str string) {
	lex.AnalysingString = str
	lex.Position = 0
	lex.Length = len(lex.AnalysingString)
}

/*Peek - получить символ в позиции lex.Position + position*/
func (lex *Lexer) Peek(position int) (rune, error) {
	for index, value := range lex.AnalysingString {
		if index == position+lex.Position {
			return value, nil
		}
	}
	lex.NextSym()
	return ' ', errors.New(prefix + "out of range")
}

/*Tokenize - токенизация входной строки*/
func (lex *Lexer) Tokenize() {
	startPosition := 0
	endPosition := 0
	newstate := false
	for lex.Position < lex.Length {
		currentRune, _ := lex.Peek(0)
		if newstate {
			startPosition = lex.Position
			newstate = false
		}
		endPosition = lex.Position

		if lex.Debug {
			log.Println(prefix+"Current Rune: ", string(currentRune))
		}
		lex.LexerAutomat.NewSymb(currentRune)

		switch lex.LexerAutomat.Buffer {
		case token.POINTER:
			lex.AddTokenReservered(token.POINTER, startPosition, endPosition)
			newstate = true
			break

		case token.FLOAT:
			lex.AddTokenReservered(token.FLOAT, startPosition, endPosition)
			newstate = true

			break

		case token.INT:
			lex.AddTokenReservered(token.INT, startPosition, endPosition)
			newstate = true

			break

		case token.ENDSTATEMENT:
			lex.AddTokenReservered(token.ENDSTATEMENT, startPosition, endPosition)
			newstate = true

			break

		case token.IDENTIFIER:
			lex.AddTokenUnreserved(token.IDENTIFIER, startPosition, endPosition, lex.LexerAutomat.CacheMemory)
			lex.PrevSym()
			newstate = true

			break
		case token.SPACE:
			lex.AddTokenReservered(token.SPACE, startPosition, endPosition)
			newstate = true

			break

		case token.NEWLINE:
			lex.AddTokenReservered(token.NEWLINE, startPosition, endPosition)
			newstate = true

			break

		case token.COMMA:
			lex.AddTokenReservered(token.COMMA, startPosition, endPosition)
			newstate = true

			break
		case token.ERROR:
			lex.AddTokenReservered(token.ERROR, startPosition, endPosition)
			newstate = true

			break
		}

		lex.NextSym()

		// if string(currentRune) == " "{
		// 	err1 := lex.NextSym()

		// 	if err1 != nil {
		// 		log.Println(prefix + "EOF")
		// 	}
		// } else {
		// 	lex.TaskStart()
		// }
	}
	if lex.LexerAutomat.CacheMemory != "" {
		lex.AddTokenUnreserved(lex.LexerAutomat.Buffer, startPosition, lex.Position, lex.LexerAutomat.CacheMemory)
	}

	lex.AggregateErrors()
}

/*AddTokenReservered - добавить зарезервированный токен в слайс с токенами*/
func (lex *Lexer) AddTokenReservered(tokenType, start, end int) {
	if lex.Debug {
		log.Println(prefix + lex.LexerAutomat.GetLog())
	}
	lex.Tokens = append(lex.Tokens, token.Token{Value: "", Type: tokenType, StartPosition: start, EndPosition: end})
	lex.LexerAutomat.Reset()
}

/*AddTokenUnreserved - добавить незарезервированный токен в слайс с токенами*/
func (lex *Lexer) AddTokenUnreserved(tokenType, start, end int, val string) {
	if lex.Debug {
		log.Println(prefix + lex.LexerAutomat.GetLog())
	}
	log.Println(prefix + lex.LexerAutomat.GetLog())
	lex.Tokens = append(lex.Tokens, token.Token{Value: val, Type: tokenType, StartPosition: start, EndPosition: end})
	lex.LexerAutomat.Reset()
}

// func (lex *Lexer) TaskStart() error {
// 	err1 := TaskForTokenType()

// 	if err1 != nil {
// 		return err1
// 	}

// 	err2 := TaskForPointer()

// 	if err2 != nil {
// 		return err2
// 	}

// 	err3 := TaskForIdentifiers()

// 	if err3 != nil {
// 		return err3
// 	}

// 	err4 := TaskForEndStatement()

// 	if err4 != nil {
// 		return err4
// 	}
// }

// func (lex *Lexer) TaskForTokenType() error {
// 	for{
// 		currentRune, err := lex.Peek(0)

// 		if err {
// 			return err
// 		}
// 		lex.ParserTypesToken.NewSymb(currentRune)
// 		switch(lex.ParserTypesToken.Buffer){
// 		case token.FLOAT:
// 			lex.Tokens = append(lex.Tokens, &token.Token{Value: lex.ParserTypesToken.CacheMemory, Type: token.FLOAT})
// 			return nil
// 		case token.INT:
// 			lex.Tokens = append(lex.Tokens, &token.Token{Value: lex.ParserTypesToken.CacheMemory, Type: token.INT})
// 			return nil
// 		case token.NONTYPE:
// 			continue
// 		}
// 	}
// }

// func (lex *Lexer) TaskForPointer() error {
// 	for {}
// }

// func (lex *Lexer) TaskForIdentifiers() error {

// }

// func (lex *Lexer) TaskForEndStatement() error {

// }

/*AggregateErrors - сбор ошибок*/
func (lex *Lexer) AggregateErrors() {
	for _, val := range lex.Tokens {
		if val.Type == token.ERROR {
			lex.AddError(-3, val)
		} else if val.Type == token.NONTYPE {
			lex.AddError(-4, val)
		}
	}
}

/*AddError - добавить новую ошибку в слайс с ошибками*/
func (lex *Lexer) AddError(errorType int, tok token.Token) {
	lex.Errors = append(lex.Errors, error_.ErrorModel{Type: errorType, Token: tok})
}

/*NextSym - смещение каретки на следующий символ*/
func (lex *Lexer) NextSym() error {
	if lex.Position < lex.Length-1 {
		lex.Position++
		return nil
	} else {
		lex.Position++
		return errors.New(prefix + "Current position is last position of analysing string")
	}
}

func (lex *Lexer) GetErrors() []error_.ErrorModel {
	resp := lex.Errors
	lex.Errors = []error_.ErrorModel{}
	lex.LexerAutomat.Reset()
	return resp
}

func (lex *Lexer) GetTokens() []token.Token {
	resp := lex.Tokens
	lex.Tokens = []token.Token{}
	return resp
}

/*PrevSym - смещение каретки на предыдущий символ*/
func (lex *Lexer) PrevSym() error {
	if lex.Position >= 0 {
		lex.Position--
		return nil
	} else {
		return errors.New(prefix + "Current position is first position of analysing string")
	}
}