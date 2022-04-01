package cmd

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	schemapb "github.com/francois-poidevin/briefly/internal/services/v1/schema"
	pb "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
		initConfig()

		err := setup()
		if err != nil {
			log.Fatalf("failed to setup the service: %v", err)
		}

		port := flag.Int("port", conf.Briefly.GRPC.ListenPort, "The server port")

		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.WithContext(ctx).Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		srv := schemapb.New(ctx, log)
		pb.RegisterSchemaAPIServer(s, srv)

		log.WithContext(ctx).Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.WithContext(ctx).Fatalf("failed to serve: %v", err)
		}

	},
}

func setup() error {

	//log handling
	log = logrus.New()

	fmt.Println(fmt.Sprintf("log format json: %t ", conf.Log.JSONFormatter))
	if conf.Log.JSONFormatter {
		fmt.Println("log format: Json")
		log.Formatter = new(logrus.JSONFormatter)
	} else {
		fmt.Println("log format: Text")
		log.Formatter = new(logrus.TextFormatter)                     //default
		log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
		log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	}

	lvl, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Fatal("Not success to parse logrus log level")
		return err
	}
	log.Level = lvl
	log.Out = os.Stdout

	return nil
}

func init() {
	startGrpcCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
}
