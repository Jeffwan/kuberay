package cluster

import (
	"context"
	"github.com/ray-project/ray-contrib/api/go_client"
	"github.com/ray-project/ray-contrib/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"time"
)

type CreateOptions struct {
	name string
	baseImage string
	pipPackages []string
	condaPackages []string
	systemPackages []string
	envs map[string]string
	customCommand string
}

func newCmdCreate() *cobra.Command {
	opts := CreateOptions{}

	cmd := &cobra.Command{
		Use:   "create <command>",
		Short: "Create Cluster Runtime",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createClusterRuntime(opts)
		},
	}

	// TODO: mark some of them optional and required
	cmd.Flags().StringVar(&opts.name, "name", "", "name of the compute runtime")
	cmd.Flags().StringVar(&opts.baseImage , "base-image", "", "name of the cloud")
	//cmd.Flags().StringArray(&opts.pipPackages , "region", "us-west-2", "name of the region")
	//cmd.Flags().StringArray(&opts.condaPackages , "az", "us-west-2a", "name of the sz")
	//cmd.Flags().StringArray(&opts.systemPackages , "head-resource-cpu", "1", "name of the sz")
	//cmd.Flags().StringArray(&opts.envs , "head-resource-memory", "1g", "name of the sz")
	cmd.Flags().StringVar(&opts.customCommand , "custom-command", "1", "name of the sz")

	return cmd
}

func createClusterRuntime(opts CreateOptions) error {
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewClusterRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	runtime := &go_client.ClusterRuntime{
		Name: "name",
		BaseImage: opts.baseImage,
		PipPackages: []string{},
		CondaPackages: []string{},
		SystemPackages: []string{},
		EnvironmentVariables: map[string]string{
			"key": "value",
		},
		CustomCommands: "",
	}

	r, err := client.CreateClusterRuntime(ctx, &go_client.CreateClusterRuntimeRequest{
		ClusterRuntime: runtime,
	})
	if err != nil {
		log.Fatalf("could not list compute runtime %v", err)
	}

	log.Printf("compute runtime %v is created", r.Id)
	return nil
}
