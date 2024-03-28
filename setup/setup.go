package setup

import "github.com/spf13/viper"

func Init() {
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}