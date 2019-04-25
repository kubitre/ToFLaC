package regularAutomat

import "strconv"

type Z6State struct {
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *Z6State) New() *Z6State {
	state.StateName = "Z6State"
	return state
}

func (state *Z6State) GetCurrentStateName() string {
	return state.StateName
}

func (state *Z6State) NextState(AllState *AllStates, context Context, mark string) {
	number, err := strconv.Atoi(mark)

	if err != nil {
		switch mark {
		case ".":
			context.SetMem(mark)
			context.SetMessage("mark is point")
			context.SetState(AllState.Z1, mark)
			return

		case " ":
			if context.CurrentActetIsLast() {
				context.SetMem(mark)
				context.SetMessage("success parsing ip")
				context.SetState(AllState.Z8, mark)
				return
			} else {
				context.SetMessage("mark[" + mark + "] does not a number and point!")
				context.SetState(AllState.ERR, mark)	
				return
			}
		default:
			context.SetMessage("mark[" + mark + "] does not a number and point!")
			context.SetState(AllState.ERR, mark)
			return
		}
	}

	if context.SuccessActetPosition() {
		context.SetMem(strconv.Itoa(number))
		context.SetMessage("input mark[" + mark + "] is number in range [0,9]")
		context.SetState(AllState.Z6 , mark)
	} else {
		context.SetMessage("error position of new mark")
		context.SetState(AllState.ERR , mark)
	}

	return
}
