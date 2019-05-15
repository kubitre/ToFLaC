package syntax

import (
	"tflac_cw/token"
)

/*InitState - окончания expression в виде ;*/
type InitState struct {
	StateName string
}

/*New - инициализация состояния старта*/
func (state *InitState) New() *InitState {
	state.StateName = "SYNTAX_AUTOMAT_INIT_STATE"
	return state
}

/*NextState - [InitState] следующее состояние:
INT, FLOAT - в сосстояние TypeState
SPACE,NEWLINE - в состояние InitState
ENDSTATEMENT - в состояние InitState с предупреждением
IDENTIFIER, POINTER - в состояние ErrorState с ошибкой
*/
func (state *InitState) NextState(states *AllStates, context Context, tok token.Token) {
	switch tok.Type {
	case token.INT:
		context.SetState(states.TYPE)
		return
	case token.FLOAT:
		context.SetState(states.TYPE)
		return
	case token.SPACE:
		context.SetState(states.INIT)
		return
	case token.NEWLINE:
		context.SetState(states.INIT)
		return
	case token.ENDSTATEMENT:
		context.NewWarn(tok, "before have not expression for do this token")
		context.SetState(states.INIT)
		return
	case token.IDENTIFIER:
		context.NewError(tok, "Unexpected identifier! expected type| space| new line. You should add type of this identifier", 0, 0, token.IDENTIFIER)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.POINTER:
		context.NewError(tok, "Unexpected pointer! you should add type before this pointer", 0, 0, token.IDENTIFIER)
		context.SetChangeState(states.INIT.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.COMMA:
		context.NewError(tok, "Unexpected comma! you should add type", 0, 0, token.IDENTIFIER)
		context.SetChangeState(states.TYPE.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	case token.ERROR:
		context.NewError(tok, "You use type with error syms!", tok.Action, tok.Position, tok.Token)
		context.SetChangeState(states.TYPE.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	default:
		context.NewError(tok, "Unexpected error!", 1, -1, -1)
		context.SetChangeState(states.TYPE.GetCurrentStateName())
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *InitState) GetCurrentStateName() string {
	return state.StateName
}
