package postgres

import (
	"github.com/marufmax/techtrends/api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Postgres struct {
	maxIdleConnection     int
	maxOpenConnection     int
	maxConnectionLifeTime time.Duration
}

const (
	_defaultMaxIdleConnection     = 10
	_defaultMaxOpenConnection     = 100
	_defaultMaxConnectionLifeTime = time.Hour
)

// Connection New Database connection which return ORM
func Connection(opts ...Option) (db *gorm.DB) {
	pg := &Postgres{
		maxIdleConnection:     _defaultMaxIdleConnection,
		maxOpenConnection:     _defaultMaxOpenConnection,
		maxConnectionLifeTime: _defaultMaxConnectionLifeTime,
	}

	// custom option
	for _, opt := range opts {
		opt(pg)
	}
	gorm.Open(postgres.Open(config.Env.DBDSN), &gorm.Config{})

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})

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
