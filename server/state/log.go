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

func (l *Log) Add(ClientId string, op string) int {
	l.Sequence = append(l.Sequence, &Operation{Committed: false, Operation: op, ClientId: ClientId})
	return len(l.Sequence)
}

func (l *Log) Get(sequenceNum int) *Operation {
	return l.Sequence[sequenceNum]
}

func (l *Log) Commit(commitNum int) (int, error) {
	if commitNum < len(l.Sequence) {
		l.Sequence[commitNum].Committed = true
		return commitNum, nil
	}
	return commitNum, fmt.Errorf("commitNum: %v bigger than log size %v", commitNum, len(l.Sequence))
}
