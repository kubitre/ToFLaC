package syntax

import (
	"bytes"
	"tflac_cw/error_"
	"tflac_cw/token"
)

type (
	/*AutomatInterface - основной интерфейс для любого состояния автомата*/
	AutomatInterface interface {
		NextState(states *AllStates, context Context, token token.Token)
		GetCurrentStateName() string
	}

	/*AutomatState - Текущее состояние автомата*/
	AutomatState struct {
		State     AutomatInterface    // текущее состояние автомата
		AllStates *AllStates          // сборка всех состояний автомата
		Tokens    []token.Token       // список токенов
		Errors    []error_.ErrorModel // полученные ошибки
		Logs      []string            // логи
	}

	/*Context - контекст для обращения к изменению состояния из состояния*/
	Context interface {
		RecordLog(log string)               // записать лог
		SetState(newState AutomatInterface) // установка текущего состояния автомата
		GetLog() string                     // получить логи
		NewSymb(token token.Token)          // новый вход
		NewError(token token.Token)         // добавление новой ошибки в стек ошибок
		ClearLogs()                         // чистка логов\
		Reset()                             // сброс состояния автомата
	}
)

/*Reset - сброс состояния автомата*/
func (automat *AutomatState) Reset() {
	automat.State = automat.AllStates.INIT
	automat.Errors = []error_.ErrorModel{}
	automat.Tokens = []token.Token{}
	automat.ClearLogs()
}

/*NewAutomat - создание нового автомата*/
func (automat *AutomatState) NewAutomat() *AutomatState {
	automat.AllStates = &AllStates{}
	automat.AllStates.New()
	automat.State = automat.AllStates.INIT

	return automat
}

/*ClearLogs - отчистка логов*/
func (automat *AutomatState) ClearLogs() {
	automat.Logs = []string{}
}

/*RecordLog - записать лог в память автомата*/
func (automat *AutomatState) RecordLog(log string) {
	automat.Logs = append(automat.Logs, log)
}

/*NewSymb - новый символ на вход автомата*/
func (automat *AutomatState) NewSymb(token token.Token) {
	automat.State.NextState(automat.AllStates, automat, token)
}

/*NewError - добавление новой ошибки через контекст*/
func (automat *AutomatState) NewError(token token.Token) {
	automat.Errors = append(automat.Errors, error_.ErrorModel{Type: 1, Token: token})
}

/*GetLog - получить все логи из автомата*/
func (automat *AutomatState) GetLog() string {
	buff := bytes.Buffer{}
	for _, val := range automat.Logs {
		buff.WriteString(val + "\n")
	}
	automat.Logs = []string{}
	return buff.String()
}

/*SetState - установка состояния newState у автомата*/
func (automat *AutomatState) SetState(newState AutomatInterface) {
	automat.RecordLog("[" + automat.State.GetCurrentStateName() + "]: change state to= " + newState.GetCurrentStateName())
	automat.State = newState

	switch newState.GetCurrentStateName() {
	case automat.AllStates.END.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, token.Token{})
		return
	}
	// if strings.Compare(newState.GetCurrentStateName(), automat.AllStates.ERROR.GetCurrentStateName()) == 0 {

	// }
}
