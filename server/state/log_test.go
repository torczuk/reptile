package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddNewLogEntryAndCommitIt(t *testing.T) {
	log := NewLog()
	log.Add("client-1", "1+1")

	assert.False(t, false, log.IsCommitted(0), "operation can't committed")
	commitNum, err := log.Commit(0)

	assert.Nil(t, err)
	assert.Equal(t, 0, commitNum)
	assert.True(t, true, log.IsCommitted(0), "operation must be committed")
}

func Test_OperationCantBeCommitted(t *testing.T) {
	log := NewLog()

	_, err := log.Commit(1)

	assert.NotNil(t, err)
}

func Test_OperationIsNotCommitted(t *testing.T) {
	log := NewLog()

	committed := log.IsCommitted(1)

	assert.False(t, false, committed, "operation can't committed")
}
