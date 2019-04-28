package lexer

import (
	"tflac_cw/token"
)

/*IntState - окончания expression в виде ;*/
type PointerState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
	StateParse      bool
}

func (state *PointerState) New() *PointerState {
	state.StateName = "LEXER_AUTOMAT_POINTER_STATE"
	state.ExampleOfState = []rune{'*'}
	state.Length = 1
	state.CurrentPosition = 0
	return state
}

func (state *PointerState) Reset() {
	state.CurrentPosition = 0
}

/*GetLastRune - получение последнего символа по позиции*/
func (state *PointerState) GetLastRune() rune {
	pos := state.CurrentPosition

	if state.CurrentPosition+1 >= state.Length {
		state.StateParse = true
	} else {
		state.CurrentPosition++
	}

	return state.ExampleOfState[pos]
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *PointerState) NextState(states *AllStates, context Context, mark rune) {
	if mark == '*' {
		context.SetMem(token.POINTER)
		context.SetState(states.INIT)
	} else {
		context.SetMem(token.ERROR)
		context.SetState(states.INIT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *PointerState) GetCurrentStateName() string {
	return state.StateName
}
