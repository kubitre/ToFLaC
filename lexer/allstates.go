package lexer

type AllStates struct {
	INIT    *InitState
	SPACE   *SpaceState
	END     *EndState
	IDENT   *IdentState
	FLOAT   *FloatState
	INT     *IntState
	POINTER *PointerState
	NEWLINE *NewLineState
	COMMA   *CommaState
}

func (states *AllStates) New() *AllStates {
	initState := &InitState{}
	spaceState := &SpaceState{}
	endState := &EndState{}
	identState := &IdentState{}
	floatState := &FloatState{}
	intState := &IntState{}
	pointerState := &PointerState{}
	newLineState := &NewLineState{}
	commaState := &CommaState{}

	states.INIT = initState.New()
	states.INT = intState.New()
	states.FLOAT = floatState.New()
	states.IDENT = identState.New()
	states.SPACE = spaceState.New()
	states.END = endState.New()
	states.POINTER = pointerState.New()
	states.NEWLINE = newLineState.New()
	states.COMMA = commaState.New()

	return states
}
