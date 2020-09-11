package prepare

import (
	"fmt"
	"github.com/jda157/cmd-currency-exchange/internal/pkg/cmd-worker"
)

type CmdSetter interface {
	SetCmdWorker(cw *cmd_worker.CmdWorker)
	//SetHttpWorker(hw *http.Client)
}

func Exchanger(ex CmdSetter, api string, args []string) error {
	cw := cmd_worker.NewCmdWorker()
	if cw == nil {
		return fmt.Errorf("can't create cmd-worker worker")
	}
	cw.Init(args)
	ex.SetCmdWorker(cw)
	return nil
}
