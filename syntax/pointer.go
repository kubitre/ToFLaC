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
			if !context.PointerAvailable() {
				context.PointerLastRemoveOne()
				context.NewError(tok, "you can not use more pointers then you declared before the first identifier", 1, 2, -1)
				context.SetChangeState(states.POINTER.GetCurrentStateName())
				context.SetState(states.ERROR)
			}
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
			// if context.PointerAvailableLast() {
			// 	context.NewError(tok, "you use a no equal amount of pointers with the first declaring state", 0, 0, token.POINTER)
			// 	context.SetChangeState(states.POINTER.GetCurrentStateName())
			// 	context.SetState(states.ERROR)
			// 	return
			// }
			context.ResetLastPointerCounter()
			if context.PointerAvailable() {
				context.SetState(states.IDENT)
			} else {
				context.NewError(tok, "expected identifier, but you write a non equal amount pointers before the first varibale name!", 1, -1, -1)
				context.SetChangeState(states.TYPE.GetCurrentStateName())
				context.SetState(states.ERROR)
			}
		}
		return
	case token.SPACE:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		if context.PointerAvailable() {
			context.ResetLastPointerCounter()
			context.SetState(states.IDENT)
		} else {
			context.NewError(tok, "expected "+strconv.Itoa(context.PointerValue())+" but request at: "+strconv.Itoa(context.PointerLasts()), 1, -1, -1)
			context.SetChangeState(states.IDENT.GetCurrentStateName())
			context.SetState(states.ERROR)
		}
		return
	case token.ENDSTATEMENT:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, "expected identifier, but detect end of statemetn", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.INT:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, " requested a type but you should end of current expression", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.FLOAT:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, " requested a type but you should end of current expression", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.COMMA:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, " requested a comma but you should end of current expression", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, "unable to process token", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.NEWLINE:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, "you should add identificator before this new line", 1, -1, -1)
		context.SetChangeState(states.POINTER.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.NONTYPE:
		if context.PointerCheckState() {
			context.PointerExpect()
		}
		context.ResetLastPointerCounter()
		context.NewError(tok, "you param is not identifier", 2, 2, token.IDENTIFIER)
		// context.NewError(tok, "you should add end of statement", 0, 1, token.ENDSTATEMENT)
		context.SetChangeState(states.IDENT.GetCurrentStateName())
		context.SetState(states.ERROR)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *PointerState) GetCurrentStateName() string {
	return state.StateName
}
