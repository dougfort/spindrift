package config

import (
	"flag"
	"strings"

	"github.com/dougfort/spindrift/types"
)

// func parseComandline to override static config items
func parseComandLine(cfg *types.Config) error {
	flag.StringVar(
		&cfg.Address,
		"address",
		cfg.Address,
		"IP address of the form <host>:<port>",
	)

	defaultSeedPeers := strings.Join(cfg.SeedPeers, ",")
	var seedPeers string
	flag.StringVar(
		&seedPeers,
		"seed-peers",
		defaultSeedPeers,
		"comma separted list of peer addresses to seed the initial pool",
	)
	if seedPeers != "" {
		cfg.SeedPeers = strings.Split(seedPeers, ",")
	}

	flag.Parse()

	return nil
}
