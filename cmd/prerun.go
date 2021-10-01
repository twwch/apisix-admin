package cmd

import (
	apisix_sdk "apisix-admin/apisix-sdk"
	"apisix-admin/component/repo"
	"apisix-admin/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfigFile() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}

var preRun = func(cmd *cobra.Command, args []string) {
	initConfigFile()
	if _, err := config.Load(""); err != nil {
		panic(fmt.Errorf("config load failed: %s", err.Error()))
	}
	if err := apisix_sdk.NewApiSixClient(config.Get().ApisixHost, config.Get().ApisixKey); err != nil {
		panic(err)
	}
	repo.Init()
}
