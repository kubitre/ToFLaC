package lexer

import (
	"errors"
	"log"
	"strings"
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
	CurrentLine     int                 // текущая строка
	CurrentColumn   int                 // текущий номер символа в строке
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
	lex.CurrentLine = 1
	lex.CurrentColumn = 0
	lex.Length = len(lex.AnalysingString)
}

/*AddNewLine - добавление новой строки*/
func (lex *Lexer) AddNewLine() {
	lex.CurrentLine = lex.CurrentLine + 1
	lex.CurrentColumn = 1
}

/*AddNewColumn - добавление нового символа*/
func (lex *Lexer) AddNewColumn() {
	lex.CurrentColumn = lex.CurrentColumn + 1
}

/*RemoveOneColumnPosition - удаление одной позиции в строке*/
func (lex *Lexer) RemoveOneColumnPosition() {
	lex.CurrentColumn = lex.CurrentColumn - 1
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
	tokenNumber := 0
	startPosition := 0
	endPosition := 0
	newstate := false

	for lex.Position < lex.Length {
		currentRune, _ := lex.Peek(0)

		lex.AddNewColumn()

		if newstate {
			startPosition = lex.Position
			newstate = false
			tokenNumber++
		}
		endPosition = lex.Position

		if lex.Debug {
			log.Println(prefix+"Current Rune: ", string(currentRune))
		}

		lex.LexerAutomat.NewSymb(currentRune)

		if lex.LexerAutomat.ContinueAdd {
			flag := false
			for _, value := range []rune{',', ';', ' ', '*', '\n'} {
				if strings.Compare(string(currentRune), string(value)) == 0 {
					flag = true
					break
				}
			}

			if !flag {
				lex.LexerAutomat.SetCacheMemCopy()
				lex.NextSym()
				continue
			} else {
				lex.LexerAutomat.SetCache(currentRune)
				lex.AddTokenReservedWithRepairSettings(lex.LexerAutomat.BufferForContinueParsing, startPosition, endPosition, lex.LexerAutomat.CacheMemory, tokenNumber, 2, lex.LexerAutomat.BufferForRepairStage)
				tokenNumber++
			}
		}

		switch lex.LexerAutomat.Buffer {
		case token.POINTER:
			lex.AddTokenReservered(token.POINTER, startPosition, endPosition, tokenNumber)
			newstate = true
			break

		case token.FLOAT:
			lex.AddTokenReservered(token.FLOAT, startPosition, endPosition, tokenNumber)
			newstate = true

			break

		case token.INT:
			lex.AddTokenReservered(token.INT, startPosition, endPosition, tokenNumber)
			newstate = true

			break

		case token.ENDSTATEMENT:
			lex.AddTokenReservered(token.ENDSTATEMENT, startPosition, endPosition, tokenNumber)
			newstate = true

			break

		case token.IDENTIFIER:
			if lex.LexerAutomat.NeedForRepairChange {
				lex.AddTokenReservedWithRepairSettings(token.IDENTIFIER, startPosition, endPosition, lex.LexerAutomat.CacheMemory, tokenNumber, 2, lex.LexerAutomat.BufferForRepairStage)
				lex.LexerAutomat.ChangeNeedRepair()
			} else {
				lex.AddTokenUnreserved(token.IDENTIFIER, startPosition, endPosition, lex.LexerAutomat.CacheMemory, tokenNumber)
			}
			lex.PrevSym()
			lex.RemoveOneColumnPosition()

			newstate = true

			break
		case token.SPACE:
			lex.AddTokenReservered(token.SPACE, startPosition, endPosition, tokenNumber)
			newstate = true

			break

		case token.NEWLINE:
			lex.AddTokenReservered(token.NEWLINE, startPosition, endPosition, tokenNumber)
			newstate = true
			lex.AddNewLine()

			break

		case token.COMMA:
			lex.AddTokenReservered(token.COMMA, startPosition, endPosition, tokenNumber)
			newstate = true

			break
		case token.ERROR:
			lex.AddTokenUnreserved(token.ERROR, startPosition, endPosition, lex.LexerAutomat.CacheMemory, tokenNumber)
			if endPosition != startPosition {
				lex.PrevSym()
			}
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
		lex.AddTokenUnreserved(lex.LexerAutomat.Buffer, startPosition, lex.Position, lex.LexerAutomat.CacheMemory, tokenNumber)
	}

	lex.AggregateErrors()
}

/*AddTokenReservered - добавить зарезервированный токен в слайс с токенами*/
func (lex *Lexer) AddTokenReservered(tokenType, start, end, tokenNumber int) {
	if lex.Debug {
		log.Println(prefix + lex.LexerAutomat.GetLog())
	}
	lex.Tokens = append(lex.Tokens, token.Token{Value: "", Type: tokenType, StartPosition: start, EndPosition: end, Line: lex.CurrentLine, Column: lex.CurrentColumn - (end - start), Index: tokenNumber})
	lex.LexerAutomat.Reset()
}

/*AddTokenUnreserved - добавить незарезервированный токен в слайс с токенами*/
func (lex *Lexer) AddTokenUnreserved(tokenType, start, end int, val string, tokenNumber int) {
	if lex.Debug {
		log.Println(prefix + lex.LexerAutomat.GetLog())
	}
	lex.Tokens = append(lex.Tokens, token.Token{Value: val, Type: tokenType, StartPosition: start, EndPosition: end, Line: lex.CurrentLine, Column: lex.CurrentColumn - (end - start), Index: tokenNumber})
	lex.LexerAutomat.Reset()
}

/*AddTokenReservedWithRepairSettings - установка токену действий по его замене на этапе нейтрализации ошибок*/
func (lex *Lexer) AddTokenReservedWithRepairSettings(tokenType, start, end int, val string, tokenNumber int, action, payload int) {
	lex.LexerAutomat.SetContinue(false)
	lex.LexerAutomat.ResetNeedRepair()
	if lex.Debug {
		log.Println(prefix + lex.LexerAutomat.GetLog())
	}
	lex.Tokens = append(lex.Tokens, token.Token{Value: val, Type: tokenType, StartPosition: start, EndPosition: end, Line: lex.CurrentLine, Column: lex.CurrentColumn - (end - start), Index: tokenNumber, Action: action, Token: payload, Position: 2})
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
	lex.Errors = append(lex.Errors, error_.ErrorModel{Token: tok, Message: "unsupporting declared token"})
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
