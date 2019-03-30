package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	rootCmd.AddCommand(clientCmd)
	cobra.OnInitialize(initConfig)
	clientCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Client config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		logrus.Error("Miss config file")
		os.Exit(1)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Can't read config: ", err)
		os.Exit(1)
	}
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "DDNS run in client mode",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
