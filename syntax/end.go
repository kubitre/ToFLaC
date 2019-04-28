package syntax

import "tflac_cw/token"

/*EndState - окончания expression в виде ;*/
type EndState struct {
	StateName string
}

func (state *EndState) New() *EndState {
	state.StateName = "SYNTAX_AUTOMAT_END_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *EndState) NextState(states *AllStates, context Context, tok token.Token) {
	context.SetState(states.INIT)
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *EndState) GetCurrentStateName() string {
	return state.StateName
}
