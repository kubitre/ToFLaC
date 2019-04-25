package regularAutomat

/*AllStates - объединение всех состояний в одном адрессном пространстве*/
type AllStates struct {
	Z0  *Z0State
	Z1  *Z1State
	Z2  *Z2State
	Z3  *Z3State
	Z4  *Z4State
	Z5  *Z5State
	Z6  *Z6State
	Z8  *Z8State
	Z9  *Z9State
	ERR *ErrorState
}
