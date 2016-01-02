package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/bbbs/backend/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "backend"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
