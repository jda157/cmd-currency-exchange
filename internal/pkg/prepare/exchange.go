package prepare

import (
	"fmt"
	"github.com/jda157/cmd-currency-exchange/internal/pkg/cmd-worker"
	http_worker "github.com/jda157/cmd-currency-exchange/internal/pkg/http-worker"
)

type CmdSetter interface {
	SetCmdWorker(cw *cmd_worker.CmdWorker)
	SetHttpWorker(hw *http_worker.HttpWorker)
	SetApi(api string)
}

func Exchanger(ex CmdSetter, api string, args []string) error {
	ex.SetApi(api)

	cw := cmd_worker.NewCmdWorker()
	if cw == nil {
		return fmt.Errorf("can't create cmd-worker worker")
	}

	if err := cw.Init(args); err != nil {
		return err
	}
	ex.SetCmdWorker(cw)

	hw := http_worker.NewHttpWorker()

	if hw == nil {
		return fmt.Errorf("can't create http-worker worker")
	}
	ex.SetHttpWorker(hw)

	return nil
}
