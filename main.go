package main

import (
	"os"

	"github.com/dougfort/spindrift/config"
	"github.com/dougfort/spindrift/incoming"
	"github.com/dougfort/spindrift/rpterr"
	"github.com/dougfort/spindrift/types"
)

func main() {
	os.Exit(run())
}

func run() int {
	var cfg types.Config
	var err error

	if cfg, err = config.Load(); err != nil {
		rpterr.ReportError(err)
		return -1
	}

	// Listen will block as long as our server is running
	if err = incoming.Listen(cfg); err != nil {
		rpterr.ReportError(err)
		return -1
	}

	return 0
}
