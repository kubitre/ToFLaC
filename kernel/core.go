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
	RepairTokens []token.Token
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
	ker.LexerAnalys.New(true)
	ker.LexerAnalys.Input(str)
	ker.LexerAnalys.Tokenize()

	ker.print(ker.LexerAnalys.Tokens)
	ker.SyntaxAnalys.Analys(ker.LexerAnalys.Tokens)
	ker.printErrors()
	ker.printRepairedSentence()
	// ker.Neitralize()
	// ker.print(ker.RepairTokens)
}

func (ker *Kernel) printRepairedSentence() {

}

func (ker *Kernel) printErrors() {
	for _, val := range ker.LexerAnalys.Errors {
		fmt.Println(prefix, prefix_lexer, val)
	}
	for _, val := range ker.SyntaxAnalys.SyntaxAnalyser.Errors {
		fmt.Println(prefix, prefix_syntax, val)
	}
}

// func (ker *Kernel) Neitralize() {
// 	var tokensRepaired []token.Token
// 	for _, val := range ker.SyntaxAnalys.SyntaxAnalyser.Errors {
// 		fmt.Println(prefix, prefix_syntax, val)
// 		switch val.Token.Action {
// 		case 0:
// 			var tokenTemp token.Token
// 			var indexTemp int
// 			for index, val2 := range ker.LexerAnalys.Tokens {

// 				if index <= val.Token.Index {
// 					if index == val.Token.Index {
// 						tokenTemp = val2
// 						indexTemp = index
// 					} else {
// 						tokensRepaired = append(tokensRepaired, val2)
// 					}

// 				} else {
// 					break
// 				}
// 			}

// 			switch val.Token.Position {
// 			case -1:
// 				break
// 			case 0:
// 				tokNew := ker.GetTokenByTypeNumber(val.Token.Token)
// 				tokNew.Index = indexTemp
// 				tokensRepaired = append(tokensRepaired, *tokNew)
// 				tokenTemp.Index = indexTemp + 1
// 				tokensRepaired = append(tokensRepaired, tokenTemp)
// 			case 1:
// 				tokNew := ker.GetTokenByTypeNumber(val.Token.Token)
// 				tokenTemp.Index = indexTemp
// 				tokensRepaired = append(tokensRepaired, tokenTemp)
// 				tokNew.Index = indexTemp + 1
// 				tokensRepaired = append(tokensRepaired, *tokNew)

// 			}

// 		case 1:
// 			for index, val2 := range ker.LexerAnalys.Tokens {
// 				if index < val.Token.Index {
// 					tokensRepaired = append(tokensRepaired, val2)
// 				} else {
// 					if index == val.Token.Index {
// 						continue
// 					} else {
// 						val2.Index--
// 						tokensRepaired = append(tokensRepaired, val2)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	ker.RepairTokens = tokensRepaired
// }

func (ker *Kernel) print(tokens []token.Token) {
	for _, val := range tokens {
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
