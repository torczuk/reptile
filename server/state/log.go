package state

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
