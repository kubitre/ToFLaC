package regularAutomat

type (
	/*AutomatInterface - основной интерфейс для любого состояния автомата*/
	AutomatInterface interface {
		NextState(states *AllStates, context Context, mark string)
		GetCurrentStateName() string
	}

	/*AutomatState - Текущее состояние автомата*/
	AutomatState struct {
		State           AutomatInterface // текущее состояние автомата
		AllStates       *AllStates       // сборка всех состояний автомата
		CacheMemory     string           // текущий ip
		AllMatches      []string         // все найденные ip адреса
		CacheAll        string           // кэш со всеми
		CurrentActet    int              // текущий актет в разбираемом выражении
		CurrentPosition int              // текщая позиция цифры в актете
		Logs            []string         // логи
	}

	/*Context - контекст для обращения к изменению состояния из состояния*/
	Context interface {
		RecordLog(log string)                            // записать лог
		SetState(newState AutomatInterface, mark string) // установка текущего состояния автомата
		GetLog() string                                  // получить логи
		NewSymbs(str string)                             // новый вход
		SuccessActetPosition() bool                      // текущая позиция в актете <=3
		CurrentActetIsLast() bool                        // текущий актет последний
		SetMessage(str string)                           // установка сообщения лога в память
		SetMem(str string)                               // Установка
		SetDump(str string)                              // дамп текущего адреса в память адресов
		ClearMem()                                       // чистка
		ClearLogs()                                      // чистка логов
		SetNewFindingIP()                                // запись нового найденного p адреса в слайс с найденными адрессами allMathces
	}
)

/*SetNewFindingIP -*/
func (am *AutomatState) SetNewFindingIP() {
	am.AllMatches = append(am.AllMatches, am.CacheMemory)
	am.CacheMemory = ""
}

/*CurrentActetIsLast - проверка на последний актет в введёном ip адресе*/
func (am *AutomatState) CurrentActetIsLast() bool {
	if am.CurrentActet == 3 {
		return true
	}
	return false
}

/*SuccessActetPosition - проверка теущей позиции в актете на*/
func (am *AutomatState) SuccessActetPosition() bool {
	if am.CurrentPosition <= 3 {
		return true
	}
	return false
}

/*SetMem - установка нового знака в ip*/
func (am *AutomatState) SetMem(str string) {
	am.CacheMemory += str
	am.CurrentPosition += 1
}

/*ClearMem - отчистка накопленной строки*/
func (am *AutomatState) ClearMem() {
	am.CacheMemory = ""
	am.CurrentPosition = 0
}

/*SetMessage - установка сообщения лога в память */
func (am *AutomatState) SetMessage(str string) {
	am.Logs = append(am.Logs, str)
}

/*SetDump - создание дампа распарсенных IP адресов*/
func (am *AutomatState) SetDump(str string) {
	am.CacheAll += am.CacheMemory + str
	am.ClearMem()
}

/*SetState - установка нового состояния автомата*/
func (am *AutomatState) SetState(newState AutomatInterface, mark string) {
	switch am.State.GetCurrentStateName() {
	case "Z0State":
		am.State = newState
		am.RecordLog("[SetState]: Start parsing, change state to Z1State...")
		am.State.NextState(am.AllStates, am, mark)
		return
	case "Z8State":
		am.ClearMem()
		am.CurrentActet = 0
		am.CurrentPosition = 0
	}

	switch newState.GetCurrentStateName() {
	case "Z1State":
		if am.CurrentActet >= 3 {
			am.RecordLog("[SetState]: Error actet increasing!")
			am.SetState(am.AllStates.ERR, mark)
			return
		}
		am.CurrentActet += 1
	}

	am.RecordLog("[" + am.State.GetCurrentStateName() + "]: " + am.CacheMemory)
	am.State = newState
}

/*RecordLog - запись лога в базу логов*/
func (am *AutomatState) RecordLog(log string) {
	am.Logs = append(am.Logs, log+"\n")
	return
}

/*GetLog - получить список логов*/
func (am *AutomatState) GetLog() string {
	var str string
	for i := 0; i < len(am.Logs); i++ {
		str += am.Logs[i]
	}
	return str
}

/*ClearLogs - отчистка логов автомата*/
func (am *AutomatState) ClearLogs() {
	am.Logs = []string{}
}

/*NewSymbs - поступление нового символа*/
func (am *AutomatState) NewSymbs(str string) {
	am.State.NextState(am.AllStates, am, str)
}

func CreateAutomat() *AutomatState {
	allStates := &AllStates{
		Z0: &Z0State{
			StateName: "Z0State",
		},
		Z1: &Z1State{
			StateName: "Z1State",
		},
		Z2: &Z2State{
			StateName: "Z2State",
		},
		Z3: &Z3State{
			StateName: "Z3State",
		},
		Z4: &Z4State{
			StateName: "Z4State",
		},
		Z5: &Z5State{
			StateName: "Z5State",
		},
		Z6: &Z6State{
			StateName: "Z6State",
		},
		Z8: &Z8State{
			StateName: "Z8State",
		},
		Z9: &Z9State{
			StateName: "Z9State",
		},
		ERR: &ErrorState{
			StateName: "ErrorState",
		},
	}
	automat := &AutomatState{
		State:        allStates.Z0,
		AllStates:    allStates,
		CacheAll:     "",
		CacheMemory:  "",
		CurrentActet: 0,
		Logs:         []string{"[AUTOMAT]: initial automat state for parsing"},
	}

	return automat
}
