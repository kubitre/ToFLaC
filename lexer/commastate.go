package lexer

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type CommaState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
}

func (state *CommaState) New() *CommaState {
	state.StateName = "LEXER_AUTOMAT_COMMA_STATE"
	state.ExampleOfState = []rune{','}
	state.Length = 1
	state.CurrentPosition = 0
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *CommaState) NextState(states *AllStates, context Context, mark rune) {
	if mark == ',' {
		context.SetMem(token.COMMA)
		context.SetState(states.INIT)
	} else {
		context.SetMem(token.ERROR)
		context.SetState(states.INIT)

	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *CommaState) GetCurrentStateName() string {
	return state.StateName
}
