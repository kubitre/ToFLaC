package syntax

import (
	"strconv"
	"tflac_cw/token"
)

/*CommaState - окончания expression в виде ;*/
type CommaState struct {
	StateName string
}

/*New - инициализация состояния Comma*/
func (state *CommaState) New() *CommaState {
	state.StateName = "SYNTAX_AUTOMAT_COMMA_STATE"
	return state
}

/*NextState - следующее состояние:

 */
func (state *CommaState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.POINTER:
		if context.PointerValue() == 0 {
			context.NewError(tok, "it is impossible to declare a pointer, because no first pointer", 1, 2, -1)
			context.SetState(states.ERROR)
			return
		} else {
			context.PointerLastCounter()
			context.SetState(states.POINTER)
			return
		}
	case token.IDENTIFIER:
		if context.PointerAvailable() {
			context.SetState(states.IDENT)
		} else {
			context.NewError(tok, "requested a identifier but we expected that you write non-equal first amount of pointers with before the current identifier pointers. First Amount of Pointers: "+strconv.Itoa(context.PointerValue()), 1, 2, -1)
			context.SetState(states.ERROR)
		}
		return
	case token.SPACE:
		context.SetState(states.COMMA)
		return
	case token.NEWLINE:
		context.SetState(states.COMMA)
		return
	case token.ENDSTATEMENT:
		context.NewError(tok, "expected id, but was not detected id", 0, 0, 3)
		context.SetState(states.ERROR)
		return
	case token.FLOAT:
		context.NewError(tok, "cannot list variables with different types", 1, -1, 2)
		context.SetState(states.ERROR)
		return
	case token.INT:
		context.NewError(tok, "cannot list variables with different types", 1, -1, 2)
		context.SetState(states.ERROR)
		return
	case token.NONTYPE:
		context.NewError(tok, "cannot recognize nontype token", 1, -1, 2)
		context.SetState(states.ERROR)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *CommaState) GetCurrentStateName() string {
	return state.StateName
}
