package event

import "github.com/okx/zkevm-node/db"

// Config for event
type Config struct {
	// DB is the database configuration
	DB db.Config `mapstructure:"DB"`
}
