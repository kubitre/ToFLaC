package typeautomat

import (
	"strings"
	"tflac_cw/token"
)

/*FLOATState - состояние типа float*/
type FLOATState struct {
	ExampleOfType   []rune
	Length          int // длина
	CurrentPosition int
	StateComplete   bool // завершение действия
}

/*GetLastRune - получить символ по его позиции*/
func (state *FLOATState) GetLastRune() rune {
	position := state.CurrentPosition

	if state.CurrentPosition+1 >= state.Length {
		state.StateComplete = true
	} else {
		state.CurrentPosition++
	}

	return state.ExampleOfType[position]
}

/*New - инициализация состояния*/
func (state *FLOATState) New() *FLOATState {
	state.CurrentPosition = 0
	state.ExampleOfType = []rune{'f', 'l', 'o', 'a', 't'}
	state.StateComplete = false
	state.Length = 5
	return state
}

/*NextState - следующее состояние:
FLOATState - следующий символ на чтение f
ErrorState - ошибка в распознании типа float
*/
func (state *FLOATState) NextState(states *AllStates, context Context, mark rune) {
	runeInAutomat := string(state.GetLastRune())

	if state.StateComplete {
		context.SetState(states.COMPLETE)
		context.SetMem(token.FLOAT)
		return
	}

	if strings.Compare(string(mark), runeInAutomat) == 0 {
		context.SetState(states.FLOAT)
	} else {
		context.SetState(states.ERROR)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *FLOATState) GetCurrentStateName() string {
	return string(state.ExampleOfType)
}
