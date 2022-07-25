package postgres

import (
	"github.com/marufmax/techtrends/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Postgres struct {
	maxIdleConnection     int
	maxOpenConnection     int
	maxConnectionLifeTime time.Duration
}

// Connection New Database connection which return ORM
func Connection() (db *gorm.DB) {
	pg := &Postgres{
		maxIdleConnection:     config.App.DBMaxIdleConnection,
		maxOpenConnection:     config.App.DBMaxOpenConnection,
		maxConnectionLifeTime: config.App.DBMaxConnectionLifeTime,
	}

	db, err := gorm.Open(postgres.Open(config.Env.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting postgres", err)
	}

	sqldb, err := db.DB()

	if err != nil {
		log.Fatal("Error connecting postgres", err)
	}

	// Setting connection Pool configuration
	sqldb.SetMaxOpenConns(pg.maxOpenConnection)
	sqldb.SetConnMaxLifetime(pg.maxConnectionLifeTime)
	sqldb.SetMaxIdleConns(pg.maxIdleConnection)

	return db
}
