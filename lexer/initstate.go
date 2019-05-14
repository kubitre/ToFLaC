package lexer

import (
	"tflac_cw/token"
)

/*InitState - начальное состояние автомата*/
type InitState struct {
	StateName string
}

func (state *InitState) New() *InitState {
	state.StateName = "LEXER_AUTOMAT_INIT_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *InitState) NextState(states *AllStates, context Context, mark rune) {
	context.SetCache(mark)
	switch mark {
	case states.INT.GetFirstRune():
		states.INT.GetLastRune()
		context.SetState(states.INT)
		return
	case states.FLOAT.GetFirstRune():
		states.FLOAT.GetLastRune()
		context.SetState(states.FLOAT)
		return
	case ' ':
		context.SetState(states.SPACE)
		return
	case ';':
		context.SetState(states.END)
		return
	case 10:
		context.SetState(states.NEWLINE)
		return
	case ',':
		context.SetState(states.COMMA)
		return
	case '*':
		context.SetState(states.POINTER)
		return
	default:
		for _, val := range []rune{'~', '`', '!', '@', '#', '$', '%', '^', '&', '(', ')', '{', '}', '[', ']', '?', '<', '>'} {
			if mark == val {
				context.SetMem(token.ERROR)
				context.SetState(states.INIT)
				return
			}
		}
		context.SetState(states.IDENT)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *InitState) GetCurrentStateName() string {
	return state.StateName
}
