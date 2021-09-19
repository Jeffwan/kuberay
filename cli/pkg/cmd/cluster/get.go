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
		Short: "Get Cluster",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getCluster(args[0])
		},
	}

	return cmd
}

func getCluster(id string) error {
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

	r, err := client.GetCluster(ctx, &go_client.GetClusterRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not list clusters: %v", err)
	}
	rows := [][]string{
		convertClusterToString(r),
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "User", "Environment", "Version", "ClusterRuntime", "ComputeRuntime"})
	table.AppendBulk(rows)
	table.Render()

	return nil

}