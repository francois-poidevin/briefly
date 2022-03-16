package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
)

var startGrpcCmd = &cobra.Command{
	Use:   "startGrpc",
	Short: "Allow to start gRPC service",
	Long:  `The HTTP gRPC service start with config parameters.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Initialize config
		// initConfig()
		fmt.Println("Start grpc Here")

	},
}
