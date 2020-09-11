package main

import (
	"flag"
	"fmt"
	"os"

	ce "github.com/jda157/cmd-currency-exchange/internal/app/console-exchanger"
	"github.com/jda157/cmd-currency-exchange/internal/pkg/prepare"
)

var (
	apiURL = flag.String("api <URL>", "https://exchangeratesapi.io", "Exchange service URL.")
)

func main() {
	flag.Parse()

	ex := ce.NewConsoleExchanger()
	doPreparation(ex)
	fmt.Println(ex.GetCmdWorker().GetArgs())

}

func doPreparation(ex *ce.Exchanger) {
	args := os.Args
	err := prepare.Exchanger(ex, *apiURL, args)
	if err != nil {
		panic(err)
	}
}
