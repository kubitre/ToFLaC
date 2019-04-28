package typeautomat

/*ErrorState - состояние ошибки автомата typeAutomat*/
type ErrorState struct {
	StateName string
}

/*New - инициализация состояния ошибки*/
func (state *ErrorState) New() *ErrorState {
	state.StateName = "TYPE_AUTOMAT_ERROR"
	return state
}

/*NextState - следующее состояние:
InitState - в начальное состояние
*/
func (state *ErrorState) NextState(states *AllStates, context Context, mark rune) {
	context.SetState(states.INIT)
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *ErrorState) GetCurrentStateName() string {
	return state.StateName
}
