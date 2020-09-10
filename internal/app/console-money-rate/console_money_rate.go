package console_money_rate

import (
	"consoleMoneyRate/internal/pkg/console"
)

type Implementation struct {
	cw *console.Worker
}

func (i *Implementation) SetConsoleWorker(cw *console.Worker) {
	i.cw = cw
}

func (i *Implementation) GetConsoleWorker() *console.Worker{
	return i.cw
}

// NewConsoleMoneyRate create new Implementation
func NewConsoleMoneyRate() *Implementation {
	return &Implementation{}
}
