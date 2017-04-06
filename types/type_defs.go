package types

// Config contains static configuration information
type Config struct {

	// The address this host listens on
	Address string

	// SeedPeers is a list of fixed addresses for contacting peers initially
	SeedPeers []string
}
