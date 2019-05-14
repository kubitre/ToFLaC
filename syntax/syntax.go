package syntax

import (
	"fmt"
	"strings"
	"tflac_cw/error_"
	"tflac_cw/token"
	"tflac_cw/warning_"
)

/*SyntaxAnalys - синтаксический анализатор*/
type SyntaxAnalys struct {
	Tokens         []token.Token
	SyntaxAnalyser *AutomatState
}

const (
	prefix = "[SYNTAX]: "
)

/*New - инициализация объекта анализатора*/
func (syntx *SyntaxAnalys) New() *SyntaxAnalys {
	analyser := &AutomatState{}
	syntx.SyntaxAnalyser = analyser.NewAutomat()
	return syntx
}

/*Analys - точка входа для синтаксического анализа*/
func (syntx *SyntaxAnalys) Analys(tokens []token.Token) {
	for _, val := range tokens {
		fmt.Println(prefix, val)

		syntx.SyntaxAnalyser.NewSymb(val)

		fmt.Println(prefix + syntx.SyntaxAnalyser.GetLog())

		if strings.Compare(syntx.SyntaxAnalyser.State.GetCurrentStateName(), syntx.SyntaxAnalyser.AllStates.ERROR.GetCurrentStateName()) == 0 {
			fmt.Println(prefix+"have error: ", syntx.SyntaxAnalyser.Errors[len(syntx.SyntaxAnalyser.Errors)-1])
		}
	}

}

/*GetErrors - получить все ошибки с этапа синтаксического анализа*/
func (syntx *SyntaxAnalys) GetErrors() []error_.ErrorModel {
	resp := syntx.SyntaxAnalyser.Errors
	return resp
}

/*GetWarnings - получить все предупреждения с этапа синтаксического анализа*/
func (syntx *SyntaxAnalys) GetWarnings() []warning_.WarningModel {
	resp := syntx.SyntaxAnalyser.Warnings
	return resp
}

/*GetRepaired - получить все исправления*/
func (syntx *SyntaxAnalys) GetRepaired() []token.Token {
	resp := syntx.SyntaxAnalyser.TokensRepaired
	return resp
}

/*Reset - сброс */
func (syntx *SyntaxAnalys) Reset() {
	syntx.SyntaxAnalyser.Reset()
}
