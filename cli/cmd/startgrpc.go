package cmd

import (
	"context"
	"flag"
	"fmt"
	"net"

	schemapb "github.com/francois-poidevin/briefly/internal/services/v1/schema"
	pb "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
	port        = flag.Int("port", 5556, "The server port")
)

var startGrpcCmd = &cobra.Command{
	Use:   "startGrpc",
	Short: "Allow to start gRPC service",
	Long:  `The HTTP gRPC service start with config parameters.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Initialize config
		// initConfig()

		fmt.Println("Start grpc Here")

		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		srv := schemapb.New()
		pb.RegisterSchemaAPIServer(s, srv)
		log.Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	},
}
