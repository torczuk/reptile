package executor

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_RunningExecutor(t *testing.T) {
	var counter = 0
	task := func() {
		counter = counter + 1
	}

	executor := NewExecutor(task, 100*time.Millisecond)
	go executor.Start()
	time.Sleep(time.Second)
	executor.Stop()

	assert.Greater(t, counter, 0)
	assert.LessOrEqual(t, counter, 10)
}
