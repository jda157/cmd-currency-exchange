package console_exchanger

import (
	"github.com/jda157/cmd-currency-exchange/internal/pkg/cmd-worker"
)

type Exchanger struct {
	cw *cmd_worker.CmdWorker
}

func (e *Exchanger) SetCmdWorker(cw *cmd_worker.CmdWorker) {
	e.cw = cw
}

func (e *Exchanger) GetCmdWorker() *cmd_worker.CmdWorker{
	return e.cw
}

// NewConsoleMoneyRate create new Exchanger
func NewConsoleExchanger() *Exchanger {
	return &Exchanger{}
}
