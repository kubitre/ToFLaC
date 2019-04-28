package lexer

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type EndState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
}

func (state *EndState) New() *EndState {
	state.StateName = "LEXER_AUTOMAT_END_STATE"
	state.ExampleOfState = []rune{';'}
	state.Length = 1
	state.CurrentPosition = 0
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *EndState) NextState(states *AllStates, context Context, mark rune) {
	if mark == ';' {
		context.SetMem(token.ENDSTATEMENT)
		context.SetState(states.INIT)
	} else {
		context.SetMem(token.ERROR)
		context.SetState(states.INIT)

	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *EndState) GetCurrentStateName() string {
	return state.StateName
}
