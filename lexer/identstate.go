package lexer

import (
	"tflac_cw/token"
)

/*IdentState - идентификатор по следующему шаблону: [A-Za-z0-9]*/
type IdentState struct {
	StateName string
}

func (state *IdentState) New() *IdentState {
	state.StateName = "LEXER_AUTOMAT_IDENTIFIER_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *IdentState) NextState(states *AllStates, context Context, mark rune) {

	switch mark {
	case ' ':
		context.SetMem(token.IDENTIFIER)
		context.SetState(states.INIT)
		return
	case ';':
		context.SetMem(token.IDENTIFIER)
		context.SetState(states.INIT)
		return
	case 10:
		context.SetMem(token.IDENTIFIER)
		context.SetState(states.INIT)
		return
	case ',':
		context.SetMem(token.IDENTIFIER)
		context.SetState(states.INIT)
		return
	default:
		for _, val := range []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '{', '}', '[', ']', '?', '`', '~'} {
			if mark == val {
				context.SetMem(token.IDENTIFIER)
				context.SetState(states.INIT)
				return
			}
		}
		context.SetCache(mark)
		context.SetState(states.IDENT)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *IdentState) GetCurrentStateName() string {
	return state.StateName
}
