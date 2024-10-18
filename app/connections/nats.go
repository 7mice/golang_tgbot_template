package connections

import (
	"github.com/nats-io/nats.go"
	"goTgTemplate/internal/config"
	"goTgTemplate/pkg/logging"
)

func ConnectToNatsBroker() *nats.Conn {
	logging.DefaultLogger.Info().Msg("Connecting to nats...")
	if config.DefaultConfig.NatsDsn == nil {
		logging.DefaultLogger.Fatal().Msg("Nats DSN is nil")
		return nil
	}
	nc, _ := nats.Connect(*config.DefaultConfig.NatsDsn)
	return nc
}
