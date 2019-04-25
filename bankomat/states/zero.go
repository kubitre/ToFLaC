package states

import "strconv"

type ZeroState struct {
}

/*NextState - следующее состояние автомата - меню*/
func (state *ZeroState) NextState(states *AllStates, context Context, mark string) {
	valueId, err := strconv.Atoi(mark)
	if err != nil {
		context.RecordLog("[ZeroState]: Enter type but not a valid id!")
		return
	}
	context.SetCurrentAccount(valueId)
	_, ошибка := context.GetAccount(valueId)

	if ошибка != nil {
		context.RecordLog("[ZeroState]: Client does not enter, because his not exist in Bank System. Error message: " + ошибка.Error())
		return
	}

	context.SetState(states.GetZ1Instance())
	context.RecordLog("[ZeroState]: Client " + mark + " enter to menu")
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *ZeroState) DoWork(data string, context Context) string {
	// not implemented
	return ""
}
