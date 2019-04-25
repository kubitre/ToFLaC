package states

import "strconv"

type Z21State struct {
}

/* NextState -
Метка:
  1 - переход к состоянию ZeroState
  2 - переход к состоянию Z1State
*/
func (state *Z21State) NextState(states *AllStates, context Context, mark string) {
	elementID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("[Z21State]: Client " + context.GetCurrentAccount() + " enter error mark!")
		return
	}

	state.DoWork("", context)

	switch elementID {
	case 1:
		context.SetState(states.GetZeroInstance())
		context.RecordLog("[Z21State]: Client " + context.GetCurrentAccount() + " redo to zeroState")
		return

	case 2:
		context.SetState(states.GetZ1Instance())
		context.RecordLog("[Z21State]: Client " + context.GetCurrentAccount() + " redo to Z1State")
		return
	default:
		context.RecordLog("[Z21State]: Client " + context.GetCurrentAccount() + " enter mark does't handled")
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z21State) DoWork(data string, context Context) string {
	userAccount := context.GetCurrentUserAccount()
	context.GetPaper()

	return strconv.Itoa(userAccount.GetCurrentBalance())
}
