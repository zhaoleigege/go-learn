package fsm

import "errors"

var UnDefineState = errors.New("未定义的状态")

// Transition 状态变化元信息
type Transition struct {
	From   string                          // 开始状态
	To     string                          // 结束状态
	Event  string                          // 发生什么事件会导致状态的变化
	Action func(args ...interface{}) error // 状态变化执行什么样的操作
}

type StateMachine struct {
	transitions []*Transition
}

func NewStateMachine(transitions []*Transition) *StateMachine {
	if len(transitions) <= 0 {
		return &StateMachine{transitions: make([]*Transition, 0)}
	}

	return &StateMachine{
		transitions: transitions,
	}
}

func (s *StateMachine) Process(curState, event string, args ...interface{}) error {
	newTran := s.findTransition(curState, event)
	if newTran == nil {
		return UnDefineState
	}

	return newTran.Action(args...)
}

func (s *StateMachine) findTransition(from, event string) *Transition {
	for _, tran := range s.transitions {
		if tran.From == from && tran.Event == event {
			return tran
		}
	}

	return nil
}
