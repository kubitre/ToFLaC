package lexer

import (
	"log"
	"strings"
	"tflac_cw/token"
)

/*IntState - окончания expression в виде ;*/
type IntState struct {
	StateName       string
	ExampleOfState  []rune
	Length          int
	CurrentPosition int
	StateParse      bool
}

func (state *IntState) New() *IntState {
	state.StateName = "LEXER_AUTOMAT_INT_STATE"
	state.ExampleOfState = []rune{'i', 'n', 't'}
	state.Length = 3
	state.CurrentPosition = 0
	return state
}

func (state *IntState) Reset() {
	state.CurrentPosition = 0
}

/*GetLastRune - получение последнего символа по позиции*/
func (state *IntState) GetLastRune() rune {
	pos := state.CurrentPosition

	if state.CurrentPosition+1 >= state.Length {
		state.StateParse = true
	} else {
		state.CurrentPosition++
	}

	return state.ExampleOfState[pos]
}

func (state *IntState) GetFirstRune() rune {
	return state.ExampleOfState[0]
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *IntState) NextState(states *AllStates, context Context, mark rune) {
	currentRune := state.GetLastRune()

	log.Println("[LEXER_INT]: Get last rune: ", string(currentRune))

	if state.StateParse {

		context.SetMem(token.INT)
		state.Reset()

		context.SetState(states.INIT)
		return
	}

	context.SetCache(mark)

	if strings.Compare(string(mark), string(currentRune)) == 0 {
		context.SetState(states.INT)
	} else {
		context.SetState(states.IDENT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *IntState) GetCurrentStateName() string {
	return state.StateName
}
