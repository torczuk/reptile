package state

import "fmt"

type Log struct {
	Sequence []*Operation
}

type Operation struct {
	Committed bool
	Operation string
	ClientId  string
}

func NewLog() *Log {
	return &Log{Sequence: make([]*Operation, 0)}
}

func (l *Log) Add(ClientId string, op string) uint32 {
	l.Sequence = append(l.Sequence, &Operation{Committed: false, Operation: op, ClientId: ClientId})
	return uint32(len(l.Sequence) - 1)
}

func (l *Log) Get(sequenceNum int) *Operation {
	return l.Sequence[sequenceNum]
}

func (l *Log) Commit(operationNum int) (int, error) {
	if operationNum < len(l.Sequence) {
		l.Sequence[operationNum].Committed = true
		return operationNum, nil
	}
	return operationNum, fmt.Errorf("operationNum: %v bigger than log size %v", operationNum, len(l.Sequence))
}

func (l *Log) IsCommitted(operationNum int) bool {
	if operationNum < len(l.Sequence) {
		return l.Sequence[operationNum].Committed
	}
	return false
}
