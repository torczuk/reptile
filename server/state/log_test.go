package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CommitOperationInTheLog(t *testing.T) {
	log := NewLog()
	log.Add("client-1", "1+1")

	assert.Equal(t, false, log.Get(0).Committed, "operation can't committed")
	commit, err := log.Commit(0)

	assert.Nil(t, err)
	assert.Equal(t, 0, commit)
}

func Test_CommitOperationNotInTheLog(t *testing.T) {
	log := NewLog()

	_, err := log.Commit(1)

	assert.NotNil(t, err)
}
