package syntax

import (
	"tflac_cw/token"
)

/*CommaState - окончания expression в виде ;*/
type CommaState struct {
	StateName string
}

/*New - инициализация состояния Comma*/
func (state *CommaState) New() *CommaState {
	state.StateName = "SYNTAX_AUTOMAT_COMMA_STATE"
	return state
}

/*NextState - следующее состояние:

 */
func (state *CommaState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.IDENTIFIER:
		context.SetState(states.IDENT)
		return
	case token.SPACE:
		context.SetState(states.COMMA)
		return
	default:
		context.NewError(tok)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *CommaState) GetCurrentStateName() string {
	return state.StateName
}
