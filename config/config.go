package config

import "github.com/dougfort/spindrift/types"

// Load populates the config struct from a stored file (.spindrift)
func Load() (types.Config, error) {

	// start with hard coded values, worry about the file later
	cfg := types.Config{
		Address:   "127.0.0.1:9000",
		SeedPeers: []string{"127.0.0.1:9001"},
	}

	if err := parseComandLine(&cfg); err != nil {
		return types.Config{}, err
	}

	return cfg, nil
}
