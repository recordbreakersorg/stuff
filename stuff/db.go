package stuff

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/recordbreakersorg/ins-stuff/stuff/db"
)

var (
	dbConn *pgx.Conn       = nil
	dbCtx  context.Context = nil
	dbQ    *db.Queries     = nil
)

func SetupDB() error {
	var err error
	dbCtx = context.Background()
	connString := os.Getenv("PSQL_INS_DB_PATH")
	dbConn, err = pgx.Connect(dbCtx, connString)
	if err != nil {
		return err
	}
	dbQ = db.New(dbConn)
	return nil
}

func UnsetDB() error {
	return dbConn.Close(dbCtx)
}
