package db

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
	"strconv"
)

type databaseCofig struct {
	dsn string
}

//var (
//	dbConnOnce sync.Once
//)

func DB() *bun.DB {
	var dbConn *bun.DB

	dbConfig := databaseCofig{
		dsn: os.Getenv("DB_DSN"),
	}

	// Below lines should be refactored to App struct
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		isDebug = false
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbConfig.dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("DEBUG"),
		bundebug.WithVerbose(isDebug),
	))

	dbConn = db

	return dbConn
}
