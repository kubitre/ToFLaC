package regularAutomat

type Z8State struct {
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *Z8State) New() *Z8State {
	state.StateName = "Z8State"
	return state
}

func (state *Z8State) GetCurrentStateName() string {
	return state.StateName
}

func (state *Z8State) NextState(AllState *AllStates, context Context, mark string) {
	context.SetDump("\n")
	context.SetMessage("complete parsing ip")
	context.SetState(AllState.Z0, mark)
}
