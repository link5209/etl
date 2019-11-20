package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv("GO_ENV")
	env := viper.GetString("GO_ENV")
	if env == "" {
		env = "development"
	}

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(workPath, "config")

	viper.SetConfigName("default")
	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName(env)
	// viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	err = viper.MergeInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("init config success!")
}
