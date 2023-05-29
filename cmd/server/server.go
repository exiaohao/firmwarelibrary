package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/server"
	"github.com/spf13/cobra"
)

var (
	opts    server.InitialOpts
	rootCmd = &cobra.Command{
		Use:          "webserver",
		Short:        "Firmware Library Server",
		Long:         "Firmware Library Server",
		SilenceUsage: true,
	}
	serverCmd = &cobra.Command{
		Use:   "run",
		Short: "Run API server",
		Long:  "Run API server",
		RunE: func(*cobra.Command, []string) error {
			server := new(server.Server)
			stopCh := setupSignalHandler()
			server.Initialize(opts)
			server.Run(stopCh)
			return nil
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringVar(&opts.ListenAddr, "address", "0.0.0.0:8080", "HTTP server port")
	serverCmd.PersistentFlags().StringVar(&opts.MySQLDSN, "mysql-dsn", "firmware:123456@tcp(127.0.0.1:13306)/firmwarelibrary?charset=utf8mb4", "MySQL DSN")
}

func main() {
	rootCmd.AddCommand(serverCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(stop)
		<-sigs
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
