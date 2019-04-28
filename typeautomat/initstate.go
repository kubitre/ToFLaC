package typeautomat

/*InitState - начальное состояние автомата*/
type InitState struct {
	StateName string
}

/*New - инициализация начального состояния автомата*/
func (state *InitState) New() *InitState {
	state.StateName = "TYPE_AUTOMAT_INIT"
	return state
}

/*NextState - следующее состояние:
INTState1 - на вход поступил символ i
FLOATState1 - на вход поступил символ f
CycleSpaces - на вход поступил символ ' '
*/
func (state *InitState) NextState(states *AllStates, context Context, mark rune) {
	switch mark {
	case states.INT.GetLastRune():
		context.SetState(states.INT)
		return
	case states.FLOAT.GetLastRune():
		// next state : float state
		context.SetState(states.FLOAT)
		return
	case ' ':
		// next state : space state
	default:
		// next state : error state
		context.SetState(states.ERROR)
		return
	}
}

/*GetCurrentStateName - получение имени состояния InitState*/
func (state *InitState) GetCurrentStateName() string {
	return state.StateName
}
