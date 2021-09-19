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

func newCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List Cluster Runtime",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listClusterRuntimes()
		},
	}

	return cmd
}

func listClusterRuntimes() error {
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

	r, err := client.ListClusterRuntimes(ctx, &go_client.ListClusterRuntimesRequest{})
	if err != nil {
		log.Fatalf("could not list cluster runtime %v", err)
	}
	runtime := r.GetRuntimes()
	rows := convertRuntimesToStrings(runtime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "BaseImage", "Image"})
	table.AppendBulk(rows)
	table.Render()

	return nil
}

func convertRuntimesToStrings(runtimes []*go_client.ClusterRuntime) [][]string {
	var data [][]string

	for _, r := range runtimes {
		data = append(data, convertRuntimesToString(r))
	}

	return data
}

func convertRuntimesToString(r *go_client.ClusterRuntime) []string {
	line := []string{r.Id, r.Name, r.BaseImage, r.Image}
	return line
}
