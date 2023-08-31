package configs

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func GetPgOptions() *pg.Options {
	return &pg.Options{
		User:     viper.GetString("db.user"),
		Database: viper.GetString("db.database"),
		Password: viper.GetString("db.password"),
		Addr:     viper.GetString("db.host"),
	}
}

func GetApiPort() int {
	return viper.GetInt("api.port")
}

func GetEnabledListeners() []string {
	return viper.GetStringSlice("api.listeners")
}
