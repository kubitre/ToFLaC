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
func (state *IdentifierState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.COMMA:
		context.SetState(states.COMMA)
		return

	case token.ENDSTATEMENT:
		context.SetState(states.END)
		return
	case token.POINTER:
		context.NewError(tok, "expected pointer but you should type space or end of statement", 0, 2, 4)
		context.SetState(states.ERROR)

	case token.SPACE:
		context.NewError(tok, "expected comma, but was detected space", 0, 0, 6)
		context.SetState(states.ERROR)
		return
	case token.NEWLINE:
		context.NewError(tok, "expected comma or end statement, but was detected new line", 0, 1, 4)
		context.SetState(states.ERROR)
		return
	case token.INT:
		context.NewError(tok, "expected end statement before declare new type varibale", 0, 0, 4)
		context.SetState(states.ERROR)
		return
	case token.FLOAT:
		context.NewError(tok, "expected end statement before declare new type varibale", 0, 0, 4)
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "expected end of statement or comma or space, but requested a error token! Check lexer output", 1, -1, -1)
		context.SetState(states.ERROR)
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *IdentifierState) GetCurrentStateName() string {
	return state.StateName
}
