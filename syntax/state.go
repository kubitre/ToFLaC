package syntax

import (
	"bytes"
	"tflac_cw/error_"
	"tflac_cw/token"
	"tflac_cw/warning_"
)

type (
	/*AutomatInterface - основной интерфейс для любого состояния автомата*/
	AutomatInterface interface {
		NextState(states *AllStates, context Context, token token.Token)
		GetCurrentStateName() string
	}

	/*AutomatState - Текущее состояние автомата*/
	AutomatState struct {
		State               AutomatInterface        // текущее состояние автомата
		AllStates           *AllStates              // сборка всех состояний автомата
		Tokens              []token.Token           // список токенов
		PointerFirstAmounts int                     // количество указателей перед первым идентификатором
		PointerAmountNext   int                     // количество указателей после последующих идентификаторов
		PointerWasDetected  bool                    // указатель был обнаружен
		SpacePosition       int                     // положение пробела (0 - после типа, 1- после идентификатора)
		Errors              []error_.ErrorModel     // полученные ошибки
		Warnings            []warning_.WarningModel // полученые предупреждения
		TokensRepaired      []token.Token           // нейтрализованные продукции
		Logs                []string                // логи
	}

	/*Context - контекст для обращения к изменению состояния из состояния*/
	Context interface {
		RecordLog(log string)                                              // записать лог
		SetState(newState AutomatInterface)                                // установка текущего состояния автомата
		GetLog() string                                                    // получить логи
		NewSymb(token token.Token)                                         // новый вход
		NewError(token token.Token, msg string, action, position, tok int) // добавление новой ошибки в стек ошибок
		NewWarn(token token.Token, msg string)                             // добавление нового предупреждения в стек предупреждений
		ClearLogs()                                                        // чистка логов\
		Reset()                                                            // сброс состояния автомата

		ChangeSpacePosition(newpos int) // присвоение позиции пробела
		CurrentSpacePosition() int      // получение текущий позиции пробела

		PointerExpect()          // обновление флажка
		PointerFirstCounter()    // обновление счётчика количества указателей для 1го идентификатора
		PointerLastCounter()     // обновление счётчика количества указателей для 2го и последующих идентификаторов
		PointerValue() int       // текущее количество указателей перед первым идентификатором
		PointerReset()           // сброс счётчика
		PointerCheckState() bool // проверка значения счётчика
		PointerAvailable() bool
		PointerLasts() int
		itFirstSectionPointer() bool
	}
)

func (automat *AutomatState) itFirstSectionPointer() bool {
	if automat.PointerAmountNext == 0 {
		return true
	} else {
		return false
	}
}

/*PointerLasts - счётчик количества указателей после первого*/
func (automat *AutomatState) PointerLasts() int {
	return automat.PointerAmountNext
}

/*PointerAvailable - количество указателей совпадает*/
func (automat *AutomatState) PointerAvailable() bool {
	if automat.PointerFirstAmounts == automat.PointerAmountNext {
		return true
	} else {
		return false
	}
}

/*PointerValue - текущее количество указателей перед первым идентификатором*/
func (automat *AutomatState) PointerValue() int {
	return automat.PointerFirstAmounts
}

/*CurrentSpacePosition - получить текущую позицию пробела (0 - после типа, -1 не было пробела, 1 - после запятой)*/
func (automat *AutomatState) CurrentSpacePosition() int {
	return automat.SpacePosition
}

/*NewWarn - добавление нового предупреждения в коллекцию предупреждений*/
func (automat *AutomatState) NewWarn(token token.Token, msg string) {
	automat.Warnings = append(automat.Warnings, warning_.WarningModel{Token: token, Message: msg})
}

/*ChangeSpacePosition - присвоение положения пробелу (после типа, после запятой)*/
func (automat *AutomatState) ChangeSpacePosition(newpos int) {
	automat.SpacePosition = newpos
}

/*PointerCheckState - проверка счётчика*/
func (automat *AutomatState) PointerCheckState() bool {
	if automat.PointerWasDetected {
		return false
	} else {
		return true
	}
}

/*PointerExpect - обновление указателя*/
func (automat *AutomatState) PointerExpect() {
	automat.PointerWasDetected = !automat.PointerWasDetected
}

/*PointerFirstCounter - увеличение счётчика указателя*/
func (automat *AutomatState) PointerFirstCounter() {
	automat.PointerFirstAmounts = automat.PointerFirstAmounts + 1
}

/*PointerLastCounter - обновление счётчика последующих за 1 идентификтаором указателей*/
func (automat *AutomatState) PointerLastCounter() {
	automat.PointerAmountNext = automat.PointerAmountNext + 1
}

/*PointerReset - сброс счётчика*/
func (automat *AutomatState) PointerReset() {
	automat.PointerFirstAmounts = 0
	automat.PointerWasDetected = false
}

/*Reset - сброс состояния автомата*/
func (automat *AutomatState) Reset() {
	automat.State = automat.AllStates.INIT
	automat.Errors = []error_.ErrorModel{}
	automat.Tokens = []token.Token{}
	automat.Warnings = []warning_.WarningModel{}
	automat.TokensRepaired = []token.Token{}
	automat.ClearLogs()
	automat.PointerReset()
}

/*NewAutomat - создание нового автомата*/
func (automat *AutomatState) NewAutomat() *AutomatState {
	automat.AllStates = &AllStates{}
	automat.AllStates.New()
	automat.State = automat.AllStates.INIT
	automat.Errors = []error_.ErrorModel{}
	automat.Warnings = []warning_.WarningModel{}
	automat.TokensRepaired = []token.Token{}
	automat.PointerReset()
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
	automat.TokensRepaired = append(automat.TokensRepaired, token)
	automat.State.NextState(automat.AllStates, automat, token)
}

/*NewError - добавление новой ошибки через контекст*/
func (automat *AutomatState) NewError(token token.Token, errormsg string, action, position, tok int) {
	// token.Action = action
	// token.Position = position
	// token.Token = tok
	automat.RepairSentence(token, action, position, tok)
	automat.Errors = append(automat.Errors, error_.ErrorModel{Token: token, Message: errormsg})
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

/*RepairSentence - исправление предложения*/
func (automat *AutomatState) RepairSentence(tok token.Token, action, position, tp int) {

	posit := automat.GetPositionByDefinedPos(position, tok)

	if tok.Action != 0 {
		switch tok.Action {
		case 2:
			automat.ChangeTokenByPosition(tok.Index, tok.Token)
			break
		}
	} else {
		switch action {
		case 0:
			automat.AddTokenByPosition(posit, tp)
			return
		case 1:
			automat.RemoveTokenByPosition(posit)
			return
		case 2:
			automat.ChangeTokenByPosition(posit, tp)
			return
		}
	}
}

/*ChangeTokenByPosition - изменить токен по его позции на тип tp*/
func (automat *AutomatState) ChangeTokenByPosition(position, tp int) {
	for _, val := range automat.TokensRepaired {
		if val.Index == position {
			val.Type = tp
			val.Value = ""
		}
	}
}

/*GetPositionByDefinedPos - получение позиции по отношению к токену*/
func (automat *AutomatState) GetPositionByDefinedPos(position int, token token.Token) int {
	switch position {
	case 0:
		return token.Index

	case 1:
		return token.Index + 1

	case 2:
		return token.Index

	default:
		return -1
	}
}

/*AddTokenByPosition - добавление токена*/
func (automat *AutomatState) AddTokenByPosition(position, tp int) {
	newToken := automat.GetTokenByTypeNum(tp)
	newToken.Index = position
	var newTokens []token.Token

	for _, val := range automat.TokensRepaired {
		if val.Index == position {
			val.Index++
			newTokens = append(newTokens, *newToken, val)
		} else {
			val.Index++
			newTokens = append(newTokens, val)
		}
	}
	automat.TokensRepaired = newTokens
	return
}

/*RemoveTokenByPosition - удаление токена*/
func (automat *AutomatState) RemoveTokenByPosition(position int) {
	var newTokens []token.Token
	for _, val := range automat.TokensRepaired {
		if val.Index == position {
			continue
		} else {
			val.Index--
			newTokens = append(newTokens, val)
		}
	}
	automat.TokensRepaired = newTokens
	return
}

/*GetTokenByTypeNum - получение токена по его номеру в enum*/
func (automat *AutomatState) GetTokenByTypeNum(tp int) *token.Token {
	switch tp {
	case token.COMMA:
		return &token.Token{Type: token.COMMA}
	case token.ENDSTATEMENT:
		return &token.Token{Type: token.ENDSTATEMENT}
	case token.FLOAT:
		return &token.Token{Type: token.FLOAT}
	case token.IDENTIFIER:
		return &token.Token{Type: token.IDENTIFIER, Value: "test"}
	case token.INT:
		return &token.Token{Type: token.INT}
	case token.NEWLINE:
		return &token.Token{Type: token.NEWLINE}
	case token.POINTER:
		return &token.Token{Type: token.POINTER}
	case token.SPACE:
		return &token.Token{Type: token.SPACE}
	default:
		return &token.Token{Type: token.NONTYPE}
	}
}

/*SetState - установка состояния newState у автомата*/
func (automat *AutomatState) SetState(newState AutomatInterface) {
	automat.RecordLog("[" + automat.State.GetCurrentStateName() + "]: change state to= " + newState.GetCurrentStateName())
	automat.State = newState

	switch newState.GetCurrentStateName() {
	case automat.AllStates.INIT.GetCurrentStateName():
		automat.PointerReset()
		return
	case automat.AllStates.END.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, token.Token{})
		return
	case automat.AllStates.ERROR.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, token.Token{})
		return

	}
	// if strings.Compare(newState.GetCurrentStateName(), automat.AllStates.ERROR.GetCurrentStateName()) == 0 {

	// }
}
