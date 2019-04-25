package states

import "strconv"

type Z3State struct{}

/* NextState -
Метка
  %100==0 - переход к подтверждению
  %100!=0 - ошибка, переход к корретировке
  -1 - выход

*/
func (state *Z3State) NextState(states *AllStates, context Context, mark string) {
	moneyValue, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("[Z3State]: Client " + context.GetCurrentAccount() + " was input error mark")
		context.SetState(states.GetZeroInstance())
		return
	}

	if moneyValue%100 == 0 && moneyValue >= 0 {
		context.RecordLog("[Z3State]: Client " + context.GetCurrentAccount() + " was send to Z31State")
		context.SetState(states.GetZ31Instance())
		return
	} else if moneyValue == -1 {
		context.RecordLog("[Z3State]: Client " + context.GetCurrentAccount() + " was cancel operation. Send to Z1State")
		context.SetState(states.GetZ1Instance())
		return
	} else if moneyValue%100 != 0 && moneyValue >= 0 {
		context.RecordLog("[Z3State]: Client " + context.GetCurrentAccount() + " was error money value entered! ")
		return
	} else {
		context.RecordLog("[Z3State]: Client " + context.GetCurrentAccount() + " was choosed error mark! Send to ZeroState")
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z3State) DoWork(data string, context Context) string {
	return ""
}
