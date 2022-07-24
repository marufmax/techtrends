package postgres

import "time"

type Option func(*Postgres)

func MaxIdleConnection(idleConn int) Option {
	return func(c *Postgres) {
		c.maxIdleConnection = idleConn
	}
}

func MaxOpenConnection(openConn int) Option {
	return func(c *Postgres) {
		c.maxOpenConnection = openConn
	}
}

func maxConnectionLifeTime(time time.Duration) Option {
	return func(c *Postgres) {
		c.maxConnectionLifeTime = time
	}
}
