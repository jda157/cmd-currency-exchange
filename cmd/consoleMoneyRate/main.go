package main

import (
	"fmt"
	"os"

	ce "github.com/jda157/cmd-currency-exchange/internal/app/console-exchanger"
	"github.com/jda157/cmd-currency-exchange/internal/pkg/prepare"
)

var (
	apiURL = "https://api.exchangeratesapi.io/latest?base="
)

func main() {
	ex := ce.NewConsoleExchanger()

	if err := doPreparation(ex); err != nil {
		fmt.Printf("%v\n %s", err, ce.UsageAnswer)
		return
	}
	response, err := ex.GetResponse()
	if err != nil {
		fmt.Printf("%v\n%s", err, ce.WrongArgs)
	} else {
		fmt.Println(response)
	}

}

func doPreparation(ex *ce.Exchanger) error {
	args := os.Args
	err := prepare.Exchanger(ex, apiURL, args)
	if err != nil {
		return err
	}
	return nil
}
