package typeautomat

/*SuccessState - успешный пасринг типа автоматом*/
type SuccessState struct {
	StateName string
}

/*New - инициализация состояния автомата SuccessState*/
func (state *SuccessState) New() *SuccessState {
	state.StateName = "TYPE_AUTOMAT_SUCCESS_STATE"
	return state
}

/*NextState - следующее состояние:
INITState - следующее состояние автомата
// INTState - следующий символ на чтение i
// FLOATState - следующий символ на чтение f
// ErrorState - ошибка в распознании типа int
*/
func (state *SuccessState) NextState(states *AllStates, context Context, mark rune) {
	context.SetState(states.INIT)
}

/*GetCurrentStateName - получить текущее имя состояния*/
func (state *SuccessState) GetCurrentStateName() string {
	return string(state.StateName)
}
