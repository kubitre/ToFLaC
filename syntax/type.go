package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type TypeState struct {
	StateName string
}

func (state *TypeState) New() *TypeState {
	state.StateName = "SYNTAX_AUTOMAT_TYPE_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *TypeState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.POINTER:
		context.SetState(states.POINTER)
		return
	case token.IDENTIFIER:
		context.SetState(states.IDENT)
		return
	case token.SPACE:
		context.SetState(states.TYPE)
		return
	default:
		context.NewError(tok)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *TypeState) GetCurrentStateName() string {
	return state.StateName
}
