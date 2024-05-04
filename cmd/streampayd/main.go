package main

import (
	"os"

	"github.com/furysport/streampay/v2/app"
	"github.com/furysport/streampay/v2/cmd/streampayd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
