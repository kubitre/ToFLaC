package syntax

import (
	"tflac_cw/token"
)

/*EndState - окончания expression в виде ;*/
type IdentifierState struct {
	StateName string
}

func (state *IdentifierState) New() *IdentifierState {
	state.StateName = "SYNTAX_AUTOMAT_IDENTIFIER_STATE"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *IdentifierState) NextState(states *AllStates, context Context, tok *token.Token) {
	switch tok.Type {
	case token.COMMA:
		context.SetState(states.COMMA)
		return

	case token.ENDSTATEMENT:
		context.SetState(states.END)
		return
	case token.POINTER:
		context.NewError(tok, "expected pointer but you should add comma", 0, 2, token.COMMA)
		context.SetChangeState(states.COMMA.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.SPACE:
		// context.NewError(tok, "expected comma or end of statement, but was detected space", 0, 0, token.COMMA)
		// context.SetChangeState(states.IDENT.GetCurrentStateName())
		// context.SetState(states.ERROR)
		context.SetState(states.IDENT)
		return
	case token.NEWLINE:
		context.NewError(tok, "expected comma or end statement, but was detected new line", 0, 0, token.ENDSTATEMENT)
		context.SetChangeState(states.END.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.INT:
		context.NewError(tok, "expected end statement before declare new type varibale", 0, 0, token.ENDSTATEMENT)
		context.SetChangeState(states.END.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.FLOAT:
		context.NewError(tok, "expected end statement before declare new type varibale", 0, 0, token.ENDSTATEMENT)
		context.SetChangeState(states.END.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "expected end of statement or comma or space, but requested a error token! Check lexer output", 1, -1, -1)
		context.SetChangeState(states.IDENT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.IDENTIFIER:
		context.NewError(tok, "before declaring new varible, you should add comma betwen them identifiers", 0, 0, token.COMMA)
		context.SetChangeState(states.IDENT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *IdentifierState) GetCurrentStateName() string {
	return state.StateName
}
