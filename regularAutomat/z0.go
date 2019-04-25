package regularAutomat

import "strconv"

type Z0State struct {
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *Z0State) New() *Z0State {
	state.StateName = "Z0State"
	return state
}

func (state *Z0State) GetCurrentStateName() string {
	return state.StateName
}

func (state *Z0State) NextState(AllState *AllStates, context Context, mark string) {
	_, err := strconv.Atoi(mark)

	if err != nil {
		context.SetMessage("in input mark[" + mark + "] does not recognize number")
		return
	}

	context.SetMessage("number $x in range of 0<=$x<=9")
	context.SetState(AllState.Z1, mark)
	return
}
