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

func newCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <command>",
		Short: "Get Compute Runtime",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRuntime(args[0])
		},
	}

	//TODO: use flag as web query params
	//cmd.Flags().BoolVar(&namespace, "namespace", "n", "ray-system", "Namespace of the configs")

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
	client := go_client.NewComputeRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetComputeRuntime(ctx, &go_client.GetComputeRuntimeRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not list compute runtime %v", err)
	}
	rows := [][]string{
		convertRuntimesToString(r),
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Region", "AZ", "ServiceType"})
	table.AppendBulk(rows)
	table.Render()

	return nil
}