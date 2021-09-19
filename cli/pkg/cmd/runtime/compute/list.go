package compute

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/ray-project/kuberay/api/go_client"
	"github.com/ray-project/kuberay/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

func newCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List Compute Runtime",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listComputeRuntimes()
		},
	}

	return cmd
}

func listComputeRuntimes() error {
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

	r, err := client.ListComputeRuntimes(ctx, &go_client.ListComputeRuntimesRequest{})
	if err != nil {
		log.Fatalf("could not list compute runtime %v", err)
	}
	runtime := r.GetRuntimes()
	rows := convertRuntimesToStrings(runtime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Region", "AZ", "ServiceType"})
	table.AppendBulk(rows)
	table.Render()

	return nil
}

func convertRuntimesToStrings(runtimes []*go_client.ComputeRuntime) [][]string {
	var data [][]string

	for _, r := range runtimes {
		data = append(data, convertRuntimesToString(r))
	}

	return data
}

func convertRuntimesToString(r *go_client.ComputeRuntime) []string {
	line := []string{r.Id, r.Name, r.Region, r.AvailabilityZone, r.HeadGroupSpec.ServiceType}
	return line
}
