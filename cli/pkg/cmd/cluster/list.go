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
			return listCluster()
		},
	}

	return cmd
}

func listCluster() error {
	// Get gRPC connection
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewClusterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	r, err := client.ListCluster(ctx, &go_client.ListClustersRequest{})
	if err != nil {
		log.Fatalf("could not list cluster clusters %v", err)
	}
	clusters := r.GetClusters()
	log.Println("no clusters found")
	rows := convertClustersToStrings(clusters)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "User", "Environment", "Version", "ClusterRuntime", "ComputeRuntime"})
	table.AppendBulk(rows)
	table.Render()

	return nil
}

func convertClustersToStrings(runtimes []*go_client.Cluster) [][]string {
	var data [][]string

	for _, r := range runtimes {
		data = append(data, convertClusterToString(r))
	}

	return data
}

func convertClusterToString(r *go_client.Cluster) []string {
	line := []string{r.Id, r.Name, r.User, r.Environment.String(), r.Version, r.ClusterRuntime, r.ComputeRuntime}
	return line
}
