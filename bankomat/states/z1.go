package states

import "strconv"

type Z1State struct {
}

/*NextState - следующее состояние автомата -
Метка
  1 - проверка баланса
  2 - выдача наличных
  3 - выход из состояния обслуживания
  * - замыкание
*/
func (state *Z1State) NextState(states *AllStates, context Context, mark string) {

	elementMenuID, err := strconv.Atoi(mark)

	if err != nil {
		context.RecordLog("[Z1State]: Error choose of menu item, was choose: " + mark)
		return
	}

	switch elementMenuID {
	// Проверка баланса
	case 1:
		context.SetState(states.GetZ2Instance())
		context.RecordLog("[Z1State]: Client " + context.GetCurrentAccount() + " was choose checking account balance")
		return
	case 2:
		// Выдача наличных
		context.SetState(states.GetZ3Instance())
		context.RecordLog("[Z1State]: Client " + context.GetCurrentAccount() + " was choose getting money")
		return
	case 3:
		context.SetState(states.GetZeroInstance())
		context.RecordLog("[Z1State]: Client " + context.GetCurrentAccount() + "was exit")
		return
	default:
		context.RecordLog("[Z1State]: Undefined mark")
		return
	}
}

/*DoWork - выполняемая рутина в состоянии
 */
func (state *Z1State) DoWork(data string, context Context) string {
	// not implemented
	return ""
}
