package main

import (
	"github.com/moto-taka/opencode/cmd"
	"github.com/moto-taka/opencode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
