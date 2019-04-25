package states

/*AutomatInterface - интерфейс для любого состояния автомата*/
type AutomatInterface interface {
	NextState()
}

/*AutomatState - нулевое состояние автомата*/
type AutomatState struct {
	state AutomatInterface
}

/*NextState - Следующее состояние автомата*/
func (am *AutomatState) NextState() {
	// return NextState()
}

/*SetState - установка нового состояния автомата*/
func (am *AutomatState) SetState(newState AutomatInterface) {
	am.state = newState
}

/*NewAutomat - создание нового автомата с нулевым состоянием */
func NewAutomat() *AutomatState {
	return &AutomatState{
		// state: &ZeroState()}
	}
}
