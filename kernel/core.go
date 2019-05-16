package kernel

import (
	"fmt"
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
	ker.LexerAnalys.New(false)
	ker.LexerAnalys.Input(str)
	ker.LexerAnalys.Tokenize()

	// ker.print(ker.LexerAnalys.Tokens)
	ker.SyntaxAnalys.Analys(ker.LexerAnalys.Tokens)
	// ker.printErrors()
	// ker.printRepairedTokens()
	// ker.printRepairedSentence()
	// ker.Neitralize()
	// ker.print(ker.RepairTokens)
}

func (ker *Kernel) PrintRepairedSentence() {
	fmt.Print("Repaired sentence: ", ker.RepairActual())
}

func (ker *Kernel) RepairActual() string {
	str := ""
	for _, value := range ker.SyntaxAnalys.GetRepaired() {
		str += ker.getTokenStringNotation(value)
	}
	return str
}

func (ker *Kernel) getTokenStringNotation(tok token.Token) string {
	switch tok.Type {
	case token.COMMA:
		return ","
	case token.ENDSTATEMENT:
		return ";"
	case token.INT:
		return "int"
	case token.FLOAT:
		return "float"
	case token.IDENTIFIER:
		return tok.Value
	case token.SPACE:
		return " "
	case token.NEWLINE:
		return "\n"
	case token.POINTER:
		return "*"
	case token.TYPE:
		return "<type>"
	}
	return ""
}

func (ker *Kernel) printRepairedTokens() {
	fmt.Println("Repaired for: ")
	for _, value := range ker.SyntaxAnalys.GetRepaired() {
		fmt.Println(ker.getTokenStringType(value.Type, value.StartPosition, value.EndPosition, value.Value))
	}
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
		fmt.Println(ker.getTokenStringType(val.Type, val.StartPosition, val.EndPosition, val.Value))
	}
}

func (ker *Kernel) getTokenStringType(tp, start, end int, value string) string {
	switch tp {
	case token.COMMA:
		return "[TOKEN: COMMA, POSITION: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.NEWLINE:
		return "[TOKEN: NEW LINE: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.SPACE:
		return "[TOKEN: SPACE: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.POINTER:
		return "[TOKEN: POINTER: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.NONTYPE:
		return "[TOKEN: NON TYPE: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.INT:
		return "[TOKEN: INTEGER TYPE: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.IDENTIFIER:
		return "[TOKEN: IDENTIFIER: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}, VALUE: " + value + "]"
	case token.FLOAT:
		return "[TOKEN: FLOAT TYPE: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.ENDSTATEMENT:
		return "[TOKEN: END OF STATEMENT: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	case token.ERROR:
		return "[TOKEN: ERROR: {" + strconv.Itoa(start) + ";" + strconv.Itoa(end) + "}]"
	default:
		return ""
	}
}
