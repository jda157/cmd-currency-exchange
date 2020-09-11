package cmd_worker

import (
	"fmt"
)
var wrongArgsError error = fmt.Errorf("wrong args")

type CmdWorker struct {
	args []string
}

func NewCmdWorker() *CmdWorker {
	return &CmdWorker{}
}

func (cw *CmdWorker) Init(args []string) error {
	switch len(args) {
	case 1:
		return fmt.Errorf("wrong args")
	case 4:
		cw.setArguments(args)
	default:
		return fmt.Errorf("wrong args")
	}
	return nil
}

func (cw *CmdWorker) GetArgs() []string {
	return cw.args
}

func (cw *CmdWorker) setArguments(args []string) {
	cw.args = args[1:]
}
