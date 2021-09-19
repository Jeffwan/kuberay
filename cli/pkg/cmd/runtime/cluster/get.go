package cluster

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/ray-project/ray-contrib/api/go_client"
	"github.com/ray-project/ray-contrib/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

func newCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <command>",
		Short: "Get Cluster Runtime",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRuntime(args[0])
		},
	}

	return cmd
}

func getRuntime(id string) error {
	// Get gRPC connection
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewClusterRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetClusterRuntime(ctx, &go_client.GetClusterRuntimeRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not list cluster runtime %v", err)
	}
	rows := [][]string{
		convertRuntimesToString(r),
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "BaseImage", "Image"})
	table.AppendBulk(rows)
	table.Render()

	return nil
}