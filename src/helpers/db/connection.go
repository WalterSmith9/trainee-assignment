package db

import (
	"github.com/WalterSmith9/trainee-assignment/src/helpers/configs"
	"github.com/go-pg/pg/v10"
	"sync"
)

var (
	connection *pg.DB
	once       sync.Once
)

func GetConnection() *pg.DB {
	once.Do(func() {
		connection = pg.Connect(configs.GetPgOptions())
	})
	return connection
}
