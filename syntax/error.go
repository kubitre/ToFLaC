package syntax

import "tflac_cw/token"

/*EndState - окончания expression в виде ;*/
type ErrorState struct {
	StateName string
}

func (state *ErrorState) New() *ErrorState {
	state.StateName = "SYNTAX_AUTOMAT_ERROR_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *ErrorState) NextState(states *AllStates, context Context, token token.Token) {
	context.SetState(states.INIT)
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *ErrorState) GetCurrentStateName() string {
	return state.StateName
}
