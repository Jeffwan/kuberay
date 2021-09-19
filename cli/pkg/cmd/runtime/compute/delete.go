package compute

import (
	"context"
	"github.com/ray-project/kuberay/api/go_client"
	"github.com/ray-project/kuberay/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <command>",
		Short: "Delete Compute Runtime",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRuntime(args[0])
		},
	}

	return cmd
}

func deleteRuntime(id string) error {
	// Get gRPC connection
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewComputeRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &go_client.DeleteComputeRuntimeRequest{
		Id: id,
	}
	if _, err := client.DeleteComputeRuntime(ctx, request); err != nil {
		log.Fatalf("could not delete compute runtime %v", err)
	}

	log.Printf("compute runtime %v has been deleted", id)
	return nil
}