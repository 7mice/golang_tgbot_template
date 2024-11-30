package connections

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"goTgTemplate/internal/config"
	"goTgTemplate/pkg/logging"
)

func ConnectToDB() *sqlx.DB {
	logging.DefaultLogger.Info().Msg("Connecting to database...")
	var err error
	dbDsn := config.DefaultConfig.DatabaseDsn
	db, err := sqlx.Connect("postgres", dbDsn)
	if err != nil {
		logging.DefaultLogger.Fatal().Msgf("Error connecting to database. Error: ", err)
	}

	return db
}
