package cmd_worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmdWorker(t *testing.T) {
	t.Parallel()
	
	cw := NewCmdWorker()
	assert.NotNil(t, cw)
}

func TestCmdWorker_Init(t *testing.T) {
	t.Parallel()
	
	cw := NewCmdWorker()
	
	args := []string{"./currency", "1000", "USD", "RUB"}
	assert.Equal(t, nil, cw.Init(args))
	assert.NoError(t, cw.Init(args))
	
	args = []string{"woop"}
	assert.Error(t, cw.Init(args))
	
	args = []string{"./currency", "1000", "USD", "RUB", "KEK"}
	assert.Error(t, cw.Init(args))
}

func TestCmdWorker_SetArguments(t *testing.T) {
	t.Parallel()
	
	args := []string{"./currency", "1000", "USD", "RUB"}
	cw := NewCmdWorker()
	cw.setArguments(args)
	
	assert.Equal(t, args[1:], cw.args)
}

func TestCmdWorker_GetArgs(t *testing.T) {
	t.Parallel()
	
	args := []string{"./currency", "100", "RUB", "ISK"}
	
	cw := NewCmdWorker()
	cw.args = args
	
	assert.Equal(t, args, cw.GetArgs())
}
