package syntax

import "tflac_cw/token"

/*EndState - окончания expression в виде ;*/
type EndState struct {
	StateName string
}

func (state *EndState) New() *EndState {
	state.StateName = "SYNTAX_AUTOMAT_END_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *EndState) NextState(states *AllStates, context Context, tok *token.Token) {
	switch tok.Type {
	case token.ENDSTATEMENT:
		context.NewError(tok, "Unexpected end, expected either a space or type of new identifier, or transfer to a new line", 1, 2, -1)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "Unable to handle token", tok.Action, tok.Position, tok.Token)
		context.SetChangeState(states.TYPE.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.COMMA:
		context.NewError(tok, "comma after; is prohibited", 1, 2, -1)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return

	case token.POINTER:
		context.NewError(tok, "you should add type before declaring this pointer!", 1, 2, -1)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return

	case token.NONTYPE:
		context.NewError(tok, "you can use varible types, spaces, newlines but you use unhandled token!", 1, 2, -1)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return

	case token.IDENTIFIER:
		context.NewError(tok, "you should add type before this identificator", 1, 2, -1)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.INT:
		context.SetState(states.TYPE)
		return
	case token.FLOAT:
		context.SetState(states.TYPE)
	default:
		context.SetState(states.INIT)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *EndState) GetCurrentStateName() string {
	return state.StateName
}
