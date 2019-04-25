package regularAutomat

import "strconv"

type Z1State struct {
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *Z1State) New() *Z1State {
	state.StateName = "Z1State"
	return state
}

func (state *Z1State) GetCurrentStateName() string {
	return state.StateName
}

func (state *Z1State) NextState(AllState *AllStates, context Context, mark string) {
	number, err := strconv.Atoi(mark)

	if err != nil {
		context.SetMessage("mark[" + mark + "] error parsing! Does not have a number")
		context.SetState(AllState.ERR, mark)
		return
	}

	context.SetMem(strconv.Itoa(number))

	switch number {
	case 0:
		context.SetMessage("input mark is 0")
		context.SetState(AllState.Z5, mark)
	case 1:
		context.SetMessage("input mark is 1")
		context.SetState(AllState.Z4, mark)
	case 2:
		context.SetMessage("input mark is 2")
		context.SetState(AllState.Z3, mark)
	default:

		context.SetMessage("input mark[" + mark + "] input in range [3,9]")
		context.SetState(AllState.Z2, mark)
	}
}
