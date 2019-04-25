package states

import "strconv"

type Z2State struct{}

/* NextState -
Метка:
  1 - переход к состоянию z2'1
  2 - переход к состоянию z2'2
  3 - возврат обратно к Z1State
  4 - переход к ZeroState

*/
func (state *Z2State) NextState(states *AllStates, context Context, mark string) {
	elementID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("[Z2State]: Enter error mark!")
		return
	}

	switch elementID {
	case 1:
		context.SetState(states.GetZ21Instance())
		context.RecordLog("[Z2State]: Client " + context.GetCurrentAccount() + " was choosed print balance on screen")
		return

	case 2:
		context.SetState(states.GetZ22Instance())
		context.RecordLog("[Z2State]: Client " + context.GetCurrentAccount() + " was choosed print balance on paper(screen)")
		return

	case 3:
		context.SetState(states.GetZ1Instance())
		context.RecordLog("[Z2State]: Client " + context.GetCurrentAccount() + " was redo to Z1State")
		return

	case 4:
		context.SetState(states.GetZeroInstance())
		context.RecordLog("[Z2State]: Client " + context.GetCurrentAccount() + " was exited from system")
		return
	default:
		context.RecordLog("[Z2State]: Carring")
		return
	}

}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z2State) DoWork(data string, context Context) string {
	return ""
}
