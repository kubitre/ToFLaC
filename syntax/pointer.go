package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type PointerState struct {
	StateName string
}

func (state *PointerState) New() *PointerState {
	state.StateName = "SYNTAX_AUTOMAT_POINTER_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *PointerState) NextState(states *AllStates, context Context, tok token.Token) {

	switch tok.Type {
	case token.POINTER:
		context.SetState(states.POINTER)
		return
	case token.IDENTIFIER:
		context.SetState(states.IDENT)
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
func (state *PointerState) GetCurrentStateName() string {
	return state.StateName
}
