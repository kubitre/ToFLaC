package syntax

import (
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
		// fmt.Println(prefix, val)

		syntx.SyntaxAnalyser.NewSymb(val)

		// fmt.Println(prefix + syntx.SyntaxAnalyser.GetLog())
	}

	syntx.correctionEndSequence()
}

/*correctionEndSequence - корректировка завершающего состояния*/
func (syntx *SyntaxAnalys) correctionEndSequence() {

	switch syntx.SyntaxAnalyser.State.GetCurrentStateName() {
	case syntx.SyntaxAnalyser.AllStates.INIT.GetCurrentStateName():
		return
	case syntx.SyntaxAnalyser.AllStates.COMMA.GetCurrentStateName():
		syntx.correctorRemoverSequence(token.COMMA, token.ENDSTATEMENT)
		return
	case syntx.SyntaxAnalyser.AllStates.END.GetCurrentStateName():
		return
	case syntx.SyntaxAnalyser.AllStates.POINTER.GetCurrentStateName():
		syntx.correctorRemoverSequence(token.COMMA, token.ENDSTATEMENT)
		return
	default:
		getLastToken := syntx.SyntaxAnalyser.TokensRepaired[len(syntx.SyntaxAnalyser.TokensRepaired)-1]
		syntx.SyntaxAnalyser.NewError(&getLastToken, "you should add ; for ending our sentence", 0, 1, token.ENDSTATEMENT)
	}

	// if strings.Compare(, ) != 0 {
	// 	getLastToken := syntx.SyntaxAnalyser.TokensRepaired[len(syntx.SyntaxAnalyser.TokensRepaired)-1]
	// 	syntx.SyntaxAnalyser.NewSymb(token.Token{Type: token.ENDSTATEMENT, Line: getLastToken.Line, Column: getLastToken.Column + 1, StartPosition: getLastToken.EndPosition + 1, EndPosition: getLastToken.EndPosition + 1})
	// }
}

func (syntx *SyntaxAnalys) correctorRemoverSequence(typeToken int, typeTokenEnd int) {
	var getReverseTokens []token.Token
	getReverseTokens = syntx.shakeToken()
	getRepairedTokens := syntx.GetRepaired()
	newTokens := []token.Token{}

	for index, val := range getReverseTokens {
		if val.Type == typeToken {
			newTokens = getRepairedTokens[:len(getRepairedTokens)-1-index]
			newTokens = append(newTokens, token.Token{Type: typeTokenEnd})
			break
		} else {
			continue
		}
	}
	syntx.SyntaxAnalyser.TokensRepaired = newTokens
}

/*shakeToken - переворот слайса с токенами*/
func (syntx *SyntaxAnalys) shakeToken() []token.Token {
	tokensRepaired := syntx.GetRepaired()
	newTokens := []token.Token{}
	for i := len(tokensRepaired) - 1; i >= 0; i-- {
		newTokens = append(newTokens, tokensRepaired[i])
	}
	return newTokens
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
