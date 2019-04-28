package typeautomat

/*AllStates - все состояния текущего автомата*/
type AllStates struct {
	INIT     *InitState
	FLOAT    *FLOATState
	INT      *INTState
	COMPLETE *SuccessState
	ERROR    *ErrorState
}

/*Init - инициализация состояний*/
func (state *AllStates) Init() {
	initState := &InitState{}
	intState := &INTState{}
	floatState := &FLOATState{}
	errorState := &ErrorState{}
	completeState := &SuccessState{}

	state.INIT = initState.New()
	state.INT = intState.New()
	state.FLOAT = floatState.New()
	state.ERROR = errorState.New()
	state.COMPLETE = completeState.New()
}
