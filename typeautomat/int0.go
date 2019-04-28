package typeautomat

import (
	"log"
	"strings"
	"tflac_cw/token"
)

/*INTState - состяние автомата на тип int*/
type INTState struct {
	CurrentPosition int
	ExampleOfType   []rune
	Length          int
	StateComplete   bool
}

/*GetLastRune - получить символ по его позиции*/
func (state *INTState) GetLastRune() rune {
	position := state.CurrentPosition

	if state.CurrentPosition+1 >= state.Length {
		state.StateComplete = true
	} else {
		state.CurrentPosition++
	}

	return state.ExampleOfType[position]
}

/*New - инициализация состояния*/
func (state *INTState) New() *INTState {
	state.CurrentPosition = 0
	state.ExampleOfType = []rune{'i', 'n', 't'}
	state.StateComplete = false
	state.Length = 3
	return state
}

/*NextState - следующее состояние:
INTState - следующий символ на чтение i
// FLOATState - следующий символ на чтение f
ErrorState - ошибка в распознании типа int
*/
func (state *INTState) NextState(states *AllStates, context Context, mark rune) {
	runeInAutomatQueue := string(state.GetLastRune())
	log.Println("[INT]: Next rune is: ", runeInAutomatQueue)

	if state.StateComplete {
		context.SetState(states.COMPLETE)
		context.SetMem(token.INT)
		return
	}

	if strings.Compare(string(mark), runeInAutomatQueue) == 0 {
		context.SetState(states.INT)
	} else {
		context.SetState(states.ERROR)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *INTState) GetCurrentStateName() string {
	return string(state.ExampleOfType)
}
