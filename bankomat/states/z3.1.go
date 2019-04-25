package states

import "strconv"

type Z31State struct{}

/* NextState -
Метка
  1 - переход к состояни печати чека
  2 - переход к состоянию Z3State
  3 - переход к ZeroState
  4 - переход к Z1State
*/
func (state *Z31State) NextState(states *AllStates, context Context, mark string) {
	elementID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("Client " + context.GetCurrentAccount() + " entered not handled mark!")
		return
	}

	switch elementID {
	case 1:
		context.RecordLog("[Z31State]: Client " + context.GetCurrentAccount() + " was choosed Z32State[print check]")
		context.SetState(states.GetZ32Instance())
		return
	case 2:
		context.RecordLog("[Z31State]: Cleint " + context.GetCurrentAccount() + " was choosed Z3State[cancel]")
		context.SetState(states.GetZ3Instance())
		return
	case 3:
		context.RecordLog("[Z31State]: Client " + context.GetCurrentAccount() + " was choosed ZeroState[exit]")
		context.SetState(states.GetZeroInstance())
		return
	case 4:
		context.RecordLog("[Z31State]: Client " + context.GetCurrentAccount() + " was choosed Z1State[main menu]")
		context.SetState(states.GetZ1Instance())
		return
	default:
		context.RecordLog("[Z31State]: Client " + context.GetCurrentAccount() + " was choosed unhandled mark!")
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z31State) DoWork(data string, context Context) string {

	dataNum, _ := strconv.Atoi(data)
	// context.
	err := context.GetCash(context.GetCurrentUserAccount().ID, dataNum)
	if err != nil {
		return err.Error()
	}
	return ""
}
