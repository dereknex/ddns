package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddns",
	Short: "DDNS client with Lightsail Networking DNS as backend",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute commands
func Execute(args []string) error {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
