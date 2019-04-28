package syntax

/*AllStates - состояния конечного автомата синтаксического анализатора*/
type AllStates struct {
	INIT    *InitState
	TYPE    *TypeState
	POINTER *PointerState
	IDENT   *IdentifierState
	COMMA   *CommaState
	END     *EndState
	ERROR   *ErrorState
}

/*New - инициализация всех состояний для автомата*/
func (states *AllStates) New() *AllStates {
	initState := &InitState{}
	typeState := &TypeState{}
	pointerState := &PointerState{}
	identState := &IdentifierState{}
	commaState := &CommaState{}
	endState := &EndState{}
	errorState := &ErrorState{}

	states.INIT = initState.New()
	states.TYPE = typeState.New()
	states.POINTER = pointerState.New()
	states.IDENT = identState.New()
	states.COMMA = commaState.New()
	states.END = endState.New()
	states.ERROR = errorState.New()

	return states
}
