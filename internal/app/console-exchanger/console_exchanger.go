package console_exchanger

import (
	"github.com/jda157/consoleMoneyRate/internal/pkg/console"
)

type Exchanger struct {
	cw *console.CmdWorker
}

func (e *Exchanger) SetCmdWorker(cw *console.CmdWorker) {
	e.cw = cw
}

func (e *Exchanger) GetCmdWorker() *console.CmdWorker{
	return e.cw
}

// NewConsoleMoneyRate create new Exchanger
func NewConsoleExchanger() *Exchanger {
	return &Exchanger{}
}
