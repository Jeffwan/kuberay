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
	namespace string
	environment string
	version string
	user string
	computeRuntime string
	clusterRuntime string
}

func newCmdCreate() *cobra.Command {
	opts := CreateOptions{}

	cmd := &cobra.Command{
		Use:   "create <command>",
		Short: "Create Cluster Runtime",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createCluster(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", "name of the compute runtime")
	cmd.Flags().StringVar(&opts.namespace , "namespace", "", "name of the cloud")
	cmd.Flags().StringVar(&opts.environment , "environment", "", "environment of the cluster")
	cmd.Flags().StringVar(&opts.version , "version", "", "version of the ray cluster")
	cmd.Flags().StringVar(&opts.computeRuntime , "compute-runtime", "", "name of the cloud")
	cmd.Flags().StringVar(&opts.clusterRuntime , "cluster-runtime", "", "name of the cloud")

	// handle user from auth and inject it.

	return cmd
}

func createCluster(opts CreateOptions) error {
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewClusterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	cluster := &go_client.Cluster{
		Name: opts.name,
		Namespace: "ray-system", // make this configuration in the future
		User: "jiaxin.shan",
		Version: opts.version,
		Environment: go_client.Cluster_DEV,
		ComputeRuntime: opts.computeRuntime,
		ClusterRuntime: opts.clusterRuntime,
	}

	r, err := client.CreateCluster(ctx, &go_client.CreateClusterRequest{
		Cluster: cluster,
	})
	if err != nil {
		log.Fatalf("could not create cluster %v", err)
	}

	log.Printf("cluster %v %v is created", r.Id, r.Name)
	return nil
}