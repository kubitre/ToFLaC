package regularAutomat

type ErrorState struct {
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *ErrorState) New() *ErrorState {
	state.StateName = "ErrorState"
	return state
}

func (state *ErrorState) NextState(allStates *AllStates, context Context, mark string) {
	context.SetState(allStates.Z0, mark)
}

func (state *ErrorState) GetCurrentStateName() string {
	return state.StateName
}
