package setup

import "github.com/spf13/viper"

func EnvEthNet() string {
	return viper.GetString("ETH_NET")
}