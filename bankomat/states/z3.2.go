package states

import "strconv"

type Z32State struct{}

/* NextState -
Метка
  1 - переход к ZeroState
  2 - переход к Z1State
*/
func (state *Z32State) NextState(states *AllStates, context Context, mark string) {
	elementID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("Client " + context.GetCurrentAccount() + " entere unhandled mark")
		context.SetState(states.GetZeroInstance())
		return
	}

	switch elementID {
	case 1:
		context.RecordLog("[Z32State]: Client " + context.GetCurrentAccount() + " was choosed ZeroState")
		context.SetState(states.GetZeroInstance())
		return

	case 2:
		context.RecordLog("[Z32State]: Client " + context.GetCurrentAccount() + " was choosed Z1State")
		context.SetState(states.GetZ1Instance())
		return
	default:
		context.RecordLog("[Z32State]: Cleitn " + context.GetCurrentAccount() + " was choosed unhandled mark!")
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z32State) DoWork(data string, context Context) string {
	return "Transaction info\n" + "User: " + context.GetCurrentAccount() + "\nOperation: Withdraw\n" + "Value: " + data + "\n"
}
