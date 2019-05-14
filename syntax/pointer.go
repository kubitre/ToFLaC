package syntax

import (
	"strconv"
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type PointerState struct {
	StateName string
}

func (state *PointerState) New() *PointerState {
	state.StateName = "SYNTAX_AUTOMAT_POINTER_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *PointerState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.POINTER:
		if context.PointerCheckState() {
			context.PointerFirstCounter()
			context.SetState(states.POINTER)
			return
		} else {
			context.PointerLastCounter()
			context.SetState(states.POINTER)
		}
		return
	case token.IDENTIFIER:
		if context.PointerCheckState() {
			context.PointerExpect()
		}

		if context.itFirstSectionPointer() {
			context.SetState(states.IDENT)
		} else {
			if context.PointerAvailable() {
				context.SetState(states.IDENT)
			} else {
				context.NewError(tok, "expected identifier, but you write a non equal amount pointers before the first varibale name!", 1, -1, -1)
				context.SetState(states.ERROR)
			}
		}
		return
	case token.SPACE:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		if context.PointerAvailable() {
			context.SetState(states.IDENT)
		} else {
			context.NewError(tok, "expected "+strconv.Itoa(context.PointerValue())+" but request at: "+strconv.Itoa(context.PointerLasts()), 1, -1, -1)
			context.SetState(states.ERROR)
		}
		return
	case token.ENDSTATEMENT:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.NewError(tok, "expected id", 0, 0, 3)
		context.SetState(states.ERROR)
		return
	case token.INT:
		context.NewError(tok, " requested a type but you should end of current expression", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	case token.FLOAT:
		context.NewError(tok, " requested a type but you should end of current expression", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	case token.COMMA:
		context.NewError(tok, " requested a comma but you should end of current expression", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "unable to process token", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	case token.NEWLINE:
		context.NewError(tok, "you should add identificator before this new line", 1, -1, -1)
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *PointerState) GetCurrentStateName() string {
	return state.StateName
}
