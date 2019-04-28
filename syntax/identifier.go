package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type IdentifierState struct {
	StateName string
}

func (state *IdentifierState) New() *IdentifierState {
	state.StateName = "SYNTAX_AUTOMAT_IDENTIFIER_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *IdentifierState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.COMMA:
		context.SetState(states.COMMA)
		return

	case token.ENDSTATEMENT:
		context.SetState(states.END)
		return
	case token.SPACE:
		context.SetState(states.IDENT)
		return
	default:
		context.NewError(tok)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *IdentifierState) GetCurrentStateName() string {
	return state.StateName
}
