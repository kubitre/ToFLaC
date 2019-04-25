package states

import (
	"errors"
	"strconv"
	"strings"

	"tflac_labs/bankomat"
)

type (

	/*AutomatInterface - интерфейс для любого состояния автомата*/
	AutomatInterface interface {
		NextState(states *AllStates, context Context, mark string)
		DoWork(data string, context Context) string
	}

	/*AutomatState - нулевое состояние автомата*/
	AutomatState struct {
		State          AutomatInterface       // текущее состояние
		Bank           *bankomat.UniverseBank // аккаунты пользователей с баблом
		Statess        *AllStates             // набор инстансов состояний
		AmountPapers   int                    // количество доступной бумаги для печати чеков
		Logs           []string               // логи
		CurrentAccount int                    // идентификатор текущего обслуживаемого клиента
		CashMemory     []string               // Ответы пользователю (для интерфейса)
	}

	/*Context - интерфейс для обращения к во всей системе из разных состояний*/
	Context interface {
		RecordLog(log string)                             // запись лога
		SetState(newState AutomatInterface)               // изменить состояние
		GetCash(id, amount int) error                     // выдать деньги
		GetPaper() error                                  // удалить лист для печати
		NewEnter(mark string)                             // новая обработка
		GetAccount(id int) (*bankomat.UserAccount, error) // получить аккаунт по его идентификатору
		GetLog() string                                   // получение логов обработки
		SetCurrentAccount(id int)                         // установка текущего аккаунт обслуживания
		GetCurrentAccount() string                        // получить идентификатор текущего аккаунта
		GetCurrentUserAccount() *bankomat.UserAccount     // получить аккаунт пользователя и интерфейс для взаимодействия с ним
	}
)

/*SetState - установка нового состояния автомата*/
func (am *AutomatState) SetState(newState AutomatInterface) {
	am.State = newState
}

/*GetCurrentUserAccount - получение аккаунта пользователя для взаимодействия с ним*/
func (am *AutomatState) GetCurrentUserAccount() *bankomat.UserAccount {
	for _, element := range am.Bank.Accounts {
		if element.ID == am.CurrentAccount {
			return &element
		}
	}

	return nil
}

/*RecordLog - запись операций пользователя*/
func (am *AutomatState) RecordLog(log string) {
	am.Logs = append(am.Logs, log)
}

/*GetLog - получить все логи*/
func (am *AutomatState) GetLog() string {
	return strings.Join(am.Logs, " \n")
}

/*GetAccount - получить аккаунт по id*/
func (am *AutomatState) GetAccount(id int) (*bankomat.UserAccount, error) {
	for _, element := range am.Bank.Accounts {
		if element.ID == id {
			return &element, nil
		}
	}

	return nil, errors.New("does not have account in bank")
}

// GetCash - получить дегьги со счёта
func (am *AutomatState) GetCash(id, amount int) error {
	// account, err := am.GetAccount(id)
	// if err != nil {
	// 	am.RecordLog("Error on state z1\nchange state to zero")
	// 	am.SetState(am.Statess.GetZeroInstance())

	// 	return errors.New("can not get cash from account, becouse not have money")
	// }

	// account.GetCash(amount)
	am.Bank.GetCashClient(id, amount)
	am.RecordLog("Client " + strconv.Itoa(id) + " get cash: " + strconv.Itoa(amount))
	return nil
}

/*GetPaper - печать чека -1 лист из банка листов*/
func (am *AutomatState) GetPaper() error {
	if am.AmountPapers-1 >= 0 {
		am.AmountPapers -= 1
		am.RecordLog("Get Paper to print")
		return nil
	}

	am.RecordLog("Paper does not have")
	return errors.New("papers: 0")
}

/*NewEnter - новый ввод*/
func (am *AutomatState) NewEnter(mark string) {
	// am.state.NextState(am.Statess, am, strconv.Itoa(id))
	am.State.NextState(am.Statess, am, mark)
	workStr := am.State.DoWork(mark, am)
	if workStr != "" {
		am.RecordLog("Client " + am.GetCurrentAccount() + " have: " + strconv.Itoa(am.GetCurrentUserAccount().GetCurrentBalance()) + " coins on bank payment")
	}
}

/*SetCurrentAccount - установка текущего аккаунта обслуживания*/
func (am *AutomatState) SetCurrentAccount(id int) {
	am.CurrentAccount = id
}

/*GetCurrentAccount - получение id текущего обрабатываемоего аккаунта*/
func (am *AutomatState) GetCurrentAccount() string {
	return strconv.Itoa(am.CurrentAccount)
}
