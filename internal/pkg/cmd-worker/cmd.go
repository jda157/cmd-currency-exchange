package cmd_worker

import (
	"fmt"
)

var wrongArgsError = fmt.Errorf("wrong args")

type CmdWorker struct {
	args []string
}

func NewCmdWorker() *CmdWorker {
	return &CmdWorker{}
}

func (cw *CmdWorker) Init(args []string) error {
	if len(args) == 4 {
		cw.setArguments(args)
                return nil
	} else {
		return fmt.Errorf("wrong args")
	}
}

func (cw *CmdWorker) GetArgs() []string {
	return cw.args
}

func (cw *CmdWorker) setArguments(args []string) {
	cw.args = args[1:]
}
