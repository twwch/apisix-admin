package cmd

import (
	"apisix-admin/config"
	"apisix-admin/server"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/twwch/gin-sdk/constant"
)

var cfgFile string


func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config/config.json| /etc/config/config.json)")
}

var (
	rootCmd = &cobra.Command{
		Use: "apisix-admin",
		//Short:   fmt.Sprintf("TEMPLATE\nCommit: %s\nDate: %s", Commit, Date),
		//Version: Version,
		PreRun: preRun,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				logger = log.WithField("service", "init")
				ctx    = context.Background()
				conf   = config.Get()
			)

			go func() {
				httpServer := server.New(constant.HTTPProtocal, conf.HttpListen).(*server.HTTP)
				logger.Error(ctx, httpServer.Server.Run(ctx))
			}()
			select {}
		},
	}
)

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(context.Background(), err)
	}
}
