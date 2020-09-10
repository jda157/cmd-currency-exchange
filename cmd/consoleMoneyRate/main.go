package main

import (
	"flag"
	"fmt"
	"github.com/jda157/cmd-currency-exchange/internal/pkg/console"
	"os"
)

var (
	apiURL = flag.String("api <URL>", "https://exchangeratesapi.io", "Exchange service URL.")
)

func main() {
	flag.Parse()
	args := os.Args
	
	cw := console.NewCmdWorker()
	cw.Init(args)
	
	fmt.Println(cw.GetArgs())
}
