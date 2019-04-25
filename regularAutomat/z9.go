package regularAutomat

import (
	"strconv"
)

type Z9State struct{
	StateName string
}

/*New - инициализация нового инстанса состояния*/
func (state *Z9State) New() *Z9State {
	state.StateName = "Z8State"
	return state
}

func (state *Z9State) GetCurrentStateName() string {
	return state.StateName
}

func (state *Z9State) NextState(AllState *AllStates, context Context, mark string) {
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

	if number >= 0 && number <= 9 {

		context.SetMem(mark)
		context.RecordLog("input mark=$x input in range 0<=$x<=9")
		context.SetState(AllState.Z6, mark)
	} else {
		context.SetMessage("input mark=$x does not input in range 0 <= $x <= 9")
		context.SetState(AllState.ERR, mark)
	}


	return
}
