package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cfgFile string

// InitConfig reads config file
func InitConfig() {
	ReadConfigFiles("./etc/", cfgFile)
}

// ReadConfigFiles viper actions
func ReadConfigFiles(path, file string) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(path)
	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("common")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	if file != "" {
		file = path + file
		var _, err = os.Stat(file)
		if !os.IsNotExist(err) {
			viper.SetConfigFile(file)
			err := viper.MergeInConfig()
			if err != nil {
				return
			}
		} else {
			panic(fmt.Sprintf("Config file \"%s\" was not found", file))
		}
	}
}
