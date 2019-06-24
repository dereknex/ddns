package cmd

import (
	"fmt"
	"github.com/dereknex/ddns/pkg/v1/config"
	"github.com/dereknex/ddns/pkg/v1/provider"
	"github.com/dereknex/ddns/pkg/v1/solver"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
const duration = 5 * time.Second

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
	if err := viper.Unmarshal(config.ClientConfig); err != nil {
		logrus.Error("Can't read config: ", err)
		os.Exit(1)
	}
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func setupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func run() {
	manager, err := solver.NewSolverManager(config.ClientConfig.Solvers)
	if err != nil {
		logrus.Error("Can not create solve manager: ", err)
		os.Exit(1)
	}
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for now := range ticker.C {
		fmt.Printf("Working at %v\n", now)
		ip, err := manager.Run()
		if err != nil {
			logrus.Error("Can not get IP address: ", err)
		}
		for k, d := range config.ClientConfig.Domains {
			domain := strings.ReplaceAll(k, "#", ".")
			p, err := provider.NewProvider(d.Provider)
			if err != nil {
				logrus.Error("Can not new domain provider: ", err)
				continue
			}
			for _, entry := range d.Entrys {
				err = p.SetRecord(entry, ip, domain)
				if err != nil {
					logrus.Error("Can not set IP: ", err)
					continue
				}
			}
		}

	}
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "DDNS run in client mode",
	Run: func(cmd *cobra.Command, args []string) {

		stopCh := setupSignalHandler()
		go run()
		//go daemon.SdNotify(false, "READY=1")

		<-stopCh
	},
}
