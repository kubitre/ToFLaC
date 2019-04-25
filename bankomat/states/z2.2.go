package states

import "strconv"

type Z22State struct{}

/* NextState -
Метка:
  1 - переход к состоянию ZeroState
  2 - переход к состоянию Z1State

*/
func (state *Z22State) NextState(states *AllStates, context Context, mark string) {
	elementID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("[Z22State]: Client " + context.GetCurrentAccount() + " enter error mark")
		context.SetState(states.GetZeroInstance())
		return
	}

	state.DoWork("", context)

	switch elementID {
	case 1:
		context.RecordLog("[Z22State]: Client " + context.GetCurrentAccount() + " was choose ZeroState")
		context.SetState(states.GetZeroInstance())
		return
	case 2:
		context.RecordLog("[Z22State]: Client " + context.GetCurrentAccount() + " was choose Z1State")
		context.SetState(states.GetZ1Instance())
		return
	default:
		context.RecordLog("[Z22State]: Client " + context.GetCurrentAccount() + " choose no handle mark. Send to ZeroState")
		context.SetState(states.GetZeroInstance())
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z22State) DoWork(data string, context Context) string {

	userAcc := context.GetCurrentUserAccount()

	return strconv.Itoa(userAcc.GetCurrentBalance())
}
