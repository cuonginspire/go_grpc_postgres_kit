package migration

import (
	"log"

	"github.com/103cuong/grpc_go_kit/configs"
	"github.com/pressly/goose"

	// postgres driver
	_ "github.com/lib/pq"
)

// MigrateDB migrate the database
func MigrateDB(command string) {
	dbstring := configs.BuildDSN()
	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, "./internal/migration"); err != nil {
		log.Fatalf("goose run failed :%v\n", err)
	}
}
