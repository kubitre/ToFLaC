package states

type AllStates struct {
	ZeroInstance *ZeroState
	Z1Instance   *Z1State
	Z2Instance   *Z2State
	Z21Instance  *Z21State
	Z22Instance  *Z22State
	Z3Instance   *Z3State
	Z31Instance  *Z31State
	Z32Instance  *Z32State
}

func (s *AllStates) GetZeroInstance() *ZeroState {
	return s.ZeroInstance
}

func (s *AllStates) GetZ1Instance() *Z1State {
	return s.Z1Instance
}

func (s *AllStates) GetZ2Instance() *Z2State {
	return s.Z2Instance
}

func (s *AllStates) GetZ21Instance() *Z21State {
	return s.Z21Instance
}

func (s *AllStates) GetZ22Instance() *Z22State {
	return s.Z22Instance
}

func (s *AllStates) GetZ3Instance() *Z3State {
	return s.Z3Instance
}

func (s *AllStates) GetZ31Instance() *Z31State {
	return s.Z31Instance
}

func (s *AllStates) GetZ32Instance() *Z32State {
	return s.Z32Instance
}
