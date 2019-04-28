package lexer

import (
	"strings"
	"tflac_cw/token"
)

/*IdentState - идентификатор*/
type FloatState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
	StateParse      bool
}

func (state *FloatState) New() *FloatState {
	state.StateName = "LEXER_AUTOMAT_IDENTIFIER_STATE"
	state.ExampleOfState = []rune{'f', 'l', 'o', 'a', 't'}
	state.Length = 5
	state.CurrentPosition = 0
	state.StateParse = false
	return state
}

func (state *FloatState) Reset() {
	state.CurrentPosition = 0
}

func (state *FloatState) GetFirstRune() rune {
	return state.ExampleOfState[0]
}

/*GetLsatRune - получение последнего символа по позиции*/
func (state *FloatState) GetLastRune() rune {

	pos := state.CurrentPosition

	if state.CurrentPosition+1 >= state.Length {
		state.StateParse = true
		return ' '
	} else {
		state.CurrentPosition++
	}

	return state.ExampleOfState[pos]
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *FloatState) NextState(states *AllStates, context Context, mark rune) {
	currentRune := state.GetLastRune()

	if state.StateParse {
		context.SetMem(token.FLOAT)
		state.Reset()

		context.SetState(states.INIT)
		return
	}

	context.SetCache(mark)

	if strings.Compare(string(mark), string(currentRune)) == 0 {
		context.SetState(states.FLOAT)
	} else {
		context.SetState(states.IDENT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *FloatState) GetCurrentStateName() string {
	return state.StateName
}
