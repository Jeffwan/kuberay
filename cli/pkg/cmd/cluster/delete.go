package cluster

import (
	"context"
	"github.com/ray-project/ray-contrib/api/go_client"
	"github.com/ray-project/ray-contrib/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <command>",
		Short: "Delete Cluster Runtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteCluster(args[0])
		},
	}

	return cmd
}

func deleteCluster(id string) error {
	// Get gRPC connection
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewClusterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &go_client.DeleteClusterRequest{
		Id: id,
	}
	if _, err := client.DeleteCluster(ctx, request); err != nil {
		log.Fatalf("could not delete cluster %v", err)
	}

	log.Printf("cluster %v has been deleted", id)
	return nil
}