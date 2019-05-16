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
func (state *ErrorState) NextState(states *AllStates, context Context, token *token.Token) {
	switch context.GetChangeState() {
	case states.TYPE.GetCurrentStateName():
		context.SetState(states.TYPE)
		return
	case states.POINTER.GetCurrentStateName():
		context.SetState(states.POINTER)
		return
	case states.IDENT.GetCurrentStateName():
		context.SetState(states.IDENT)
		return
	case states.END.GetCurrentStateName():
		context.SetState(states.END)
		return
	case states.COMMA.GetCurrentStateName():
		context.SetState(states.COMMA)
		return
	default:
		context.SetState(states.INIT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *ErrorState) GetCurrentStateName() string {
	return state.StateName
}
