package constants

import "time"

// Define constants
const (
	LocalEnvName           = "development"
	ProductionEnvName      = "produduction"
	DefaultPort            = "0.0.0.0:8000"
	DefaultReadTimeout     = 5 * time.Second
	DefaultWriteTimeout    = 5 * time.Second
	DefaultShutdownTimeout = 3 * time.Second
)
