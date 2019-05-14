package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type TypeState struct {
	StateName string
}

func (state *TypeState) New() *TypeState {
	state.StateName = "SYNTAX_AUTOMAT_TYPE_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *TypeState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.POINTER:
		context.PointerExpect()
		context.PointerFirstCounter()
		context.SetState(states.POINTER)
		return
	case token.IDENTIFIER:
		if context.CurrentSpacePosition() == 0 {
			context.SetState(states.IDENT)
		} else {
			context.NewError(tok, "expected space or pointer before declare varibale name", 1, 0, 2)
			context.SetState(states.ERROR)
		}
		return
	case token.SPACE:
		context.ChangeSpacePosition(0)
		context.SetState(states.TYPE)
		return
	case token.NEWLINE:
		context.NewError(tok, "identifier was expected but was not detected", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	case token.ENDSTATEMENT:
		context.NewError(tok, "to end the announcement you must enter an identifier", 0, 0, 3)
		context.SetState(states.ERROR)
		return
	case token.COMMA:
		context.NewError(tok, "you can not use comma between types!", 1, -1, 4)
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "expected true value, but we expect error!", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *TypeState) GetCurrentStateName() string {
	return state.StateName
}
