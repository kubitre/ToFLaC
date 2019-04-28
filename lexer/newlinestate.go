package lexer

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type NewLineState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
}

func (state *NewLineState) New() *NewLineState {
	state.StateName = "LEXER_AUTOMAT_NEW_LINE_STATE"
	state.ExampleOfState = []rune{'\n'}
	state.Length = 1
	state.CurrentPosition = 0
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *NewLineState) NextState(states *AllStates, context Context, mark rune) {
	if mark == 10 {
		context.SetMem(token.NEWLINE)
		context.SetState(states.INIT)
	} else {
		context.SetMem(token.ERROR)
		context.SetState(states.INIT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *NewLineState) GetCurrentStateName() string {
	return state.StateName
}
