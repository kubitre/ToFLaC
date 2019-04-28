package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type InitState struct {
	StateName string
}

func (state *InitState) New() *InitState {
	state.StateName = "SYNTAX_AUTOMAT_INIT_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *InitState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.INT:
	case token.FLOAT:
		context.SetState(states.TYPE)
		return
	case token.SPACE:
		context.SetState(states.INIT)
		return
	default:
		context.NewError(tok)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *InitState) GetCurrentStateName() string {
	return state.StateName
}
