package console_exchanger

import (
	"fmt"
	cw "github.com/jda157/cmd-currency-exchange/internal/pkg/cmd-worker"
	hw "github.com/jda157/cmd-currency-exchange/internal/pkg/http-worker"
)

type emptyData struct{}

type Exchanger struct {
	api   string
	cw    *cw.CmdWorker
	hw    *hw.HttpWorker
}

func (e *Exchanger) SetHttpWorker(hw *hw.HttpWorker) {
	e.hw = hw
}

func (e *Exchanger) GetHttpWorker() *cw.CmdWorker {
	return e.cw
}

func (e *Exchanger) SetApi(api string) {
	e.api = api
}

func (e *Exchanger) GetApi() string {
	return e.api
}

func (e *Exchanger) SetCmdWorker(cw *cw.CmdWorker) {
	e.cw = cw
}

func (e *Exchanger) GetCmdWorker() *cw.CmdWorker {
	return e.cw
}

func (e *Exchanger) GetResponse() (string, error) {
	args := e.cw.GetArgs()
	amountStr := args[0]
	src := args[1]
	dst := args[2]

	amountFloat, err := toFloatAmount(amountStr)
	if err != nil {
		return "", err
	}

	resp, err := e.hw.GetResponse(fmt.Sprintf("%s%s", e.GetApi(), src))
	if err != nil {
		return "", err
	}

	currentDst, err := getCurrentDst(resp, src, dst)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", currentDst*amountFloat), nil
}

// NewConsoleMoneyRate create new Exchanger
func NewConsoleExchanger() *Exchanger {
	return &Exchanger{}
}
