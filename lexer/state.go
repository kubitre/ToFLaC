package lexer

import (
	"bytes"
	"tflac_cw/token"
)

type (
	/*AutomatInterface - основной интерфейс для любого состояния автомата*/
	AutomatInterface interface {
		NextState(states *AllStates, context Context, mark rune)
		GetCurrentStateName() string
	}

	/*AutomatState - Текущее состояние автомата*/
	AutomatState struct {
		State                    AutomatInterface // текущее состояние автомата
		AllStates                *AllStates       // сборка всех состояний автомата
		Buffer                   int              // буфер для автомата
		BufferForRepairStage     int              // заменяемый токен
		NeedForRepairChange      bool             // требуется ли замена на BufferForRepairStage
		ContinueAdd              bool             // продолжение парсинга
		CacheMemory              string           // текущее значение токена
		BufferForContinueParsing int
		CacheMemoryCopy          string
		Logs                     []string // логи
	}

	/*Context - контекст для обращения к изменению состояния из состояния*/
	Context interface {
		RecordLog(log string)               // записать лог
		SetState(newState AutomatInterface) // установка текущего состояния автомата
		SetCache(mark rune)                 // установка значения токена
		GetLog() string                     // получить логи
		SetMem(value int)                   // устанлока текущего выделяемого типа
		SetTokenForRepairStage(typ int)     // установка значения для токена после синтаксического анализа
		ChangeNeedRepair()                  // изменить значение добавления изменения на стадии исправления ошибок
		NewSymb(mark rune)                  // новый вход
		ClearLogs()                         // чистка логов\
		Reset()                             // сброс состояния автомата
		SetContinue(flag bool)              // переключить на продолжение парса в заданный токен
		SetCacheMemCopy()                   // установка значения в кэш
		ResetNeedRepair()
	}
)

func (automat *AutomatState) ResetNeedRepair() {
	automat.NeedForRepairChange = false
	automat.BufferForRepairStage = -1
}

func (automat *AutomatState) SetCacheMemCopy() {
	automat.CacheMemoryCopy += automat.CacheMemory
}

/*SetContinue - продолжение*/
func (automat *AutomatState) SetContinue(flag bool) {
	automat.ContinueAdd = flag
	automat.CacheMemoryCopy = automat.CacheMemory
	automat.CacheMemory = ""
	automat.BufferForContinueParsing = automat.Buffer
}

/*ChangeNeedRepair - изменить текущее состояние флага*/
func (automat *AutomatState) ChangeNeedRepair() {
	automat.NeedForRepairChange = false
	automat.BufferForRepairStage = 0
}

/*SetTokenForRepairStage - установка токена для его вставки на этапе нейтрализации ошибок после синтаксического анализа*/
func (automat *AutomatState) SetTokenForRepairStage(typ int) {
	automat.NeedForRepairChange = true
	automat.BufferForRepairStage = typ
}

/*Reset - сброс значений текущего автомата*/
func (automat *AutomatState) Reset() {
	automat.State = automat.AllStates.INIT
	automat.Buffer = token.NONTYPE
	automat.CacheMemory = ""
	automat.ClearLogs()
}

/*SetCache - установка значения токена*/
func (automat *AutomatState) SetCache(mark rune) {
	automat.CacheMemory += string(mark)
}

/*SetMem - установка текущего выделяемого типа*/
func (automat *AutomatState) SetMem(value int) {
	automat.Buffer = value
}

/*NewAutomat - создание нового автомата*/
func (automat *AutomatState) NewAutomat() *AutomatState {
	automat.AllStates = &AllStates{}
	automat.AllStates.New()
	automat.State = automat.AllStates.INIT
	automat.Buffer = token.NONTYPE

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
func (automat *AutomatState) NewSymb(mark rune) {
	automat.State.NextState(automat.AllStates, automat, mark)
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
	case automat.AllStates.SPACE.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, ' ')
		return
	case automat.AllStates.END.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, ';')
		return
	case automat.AllStates.POINTER.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, '*')
		return
	case automat.AllStates.NEWLINE.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, '\n')
		return
	case automat.AllStates.COMMA.GetCurrentStateName():
		automat.State.NextState(automat.AllStates, automat, ',')
		return
	}

}
