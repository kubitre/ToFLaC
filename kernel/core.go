package kernel

import (
	"fmt"
	"log"
	"strconv"
	"tflac_cw/lexer"
	"tflac_cw/syntax"
	"tflac_cw/token"
)

/*Kernel - ядро языка*/
type Kernel struct {
	LexerAnalys  *lexer.Lexer
	SyntaxAnalys *syntax.SyntaxAnalys
}

const (
	prefix        = "[KERNEL]: "
	prefix_lexer  = "{LEXER}"
	prefix_syntax = "{SYNTAX}"
)

/*New - инициализация ядра*/
func (ker *Kernel) New() *Kernel {
	ker.LexerAnalys = &lexer.Lexer{}
	syntx := &syntax.SyntaxAnalys{}
	ker.SyntaxAnalys = syntx.New()
	return ker
}

/*Input - ввод нового текста*/
func (ker *Kernel) Input(str string) {
	ker.LexerAnalys.New(false)
	ker.LexerAnalys.Input(str)
	ker.LexerAnalys.Tokenize()

	ker.print()
	ker.SyntaxAnalys.Analys(ker.LexerAnalys.Tokens)
	ker.printErrors()
}

func (ker *Kernel) printErrors() {
	for _, val := range ker.LexerAnalys.Errors {
		fmt.Println(prefix, prefix_lexer, val)
	}
	for _, val := range ker.SyntaxAnalys.SyntaxAnalyser.Errors {
		fmt.Println(prefix, prefix_syntax, val)
	}
}

func (ker *Kernel) print() {
	for _, val := range ker.LexerAnalys.Tokens {
		switch val.Type {
		case token.COMMA:
			log.Println("[TOKEN: COMMA, POSITION: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.NEWLINE:
			log.Println("[TOKEN: NEW LINE: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.SPACE:
			log.Println("[TOKEN: SPACE: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.POINTER:
			log.Println("[TOKEN: POINTER: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.NONTYPE:
			log.Println("[TOKEN: NON TYPE: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.INT:
			log.Println("[TOKEN: INTEGER TYPE: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.IDENTIFIER:
			log.Println("[TOKEN: IDENTIFIER: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}, VALUE: " + val.Value + "]")
		case token.FLOAT:
			log.Println("[TOKEN: FLOAT TYPE: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.ENDSTATEMENT:
			log.Println("[TOKEN: END OF STATEMENT: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		case token.ERROR:
			log.Println("[TOKEN: ERROR: {" + strconv.Itoa(val.StartPosition) + ";" + strconv.Itoa(val.EndPosition) + "}]")
		}
	}
}
