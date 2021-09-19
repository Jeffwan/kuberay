package compute

import (
	"context"
	"github.com/ray-project/kuberay/api/go_client"
	"github.com/ray-project/kuberay/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
	"time"
)

type CreateOptions struct {
	name string
	cloud string
	region string
	az string
	headResourceCpu string
	headResourceMemory string
	workerResourceCpu string
	workerResourceMemory string
}

func newCmdCreate() *cobra.Command {
	opts := CreateOptions{}

	cmd := &cobra.Command{
		Use:   "create <command>",
		Short: "Create Compute Runtime",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createComputeRuntime(opts)
		},
	}

	// TODO: how to mark them as required
	cmd.Flags().StringVar(&opts.name, "name", "", "name of the compute runtime")
	cmd.Flags().StringVar(&opts.cloud , "cloud", "", "name of the cloud")
	cmd.Flags().StringVar(&opts.region , "region", "us-west-2", "name of the region")
	cmd.Flags().StringVar(&opts.az , "az", "us-west-2a", "name of the sz")
	cmd.Flags().StringVar(&opts.headResourceCpu , "head-resource-cpu", "1", "name of the sz")
	cmd.Flags().StringVar(&opts.headResourceMemory , "head-resource-memory", "1g", "name of the sz")
	cmd.Flags().StringVar(&opts.workerResourceCpu , "worker-resource-cpu", "1", "name of the sz")
	cmd.Flags().StringVar(&opts.workerResourceMemory , "worker-resource-memory", "1g", "name of the sz")

	return cmd
}

func createComputeRuntime(opts CreateOptions) error {
	conn, err := cmdutil.GetGrpcConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// build gRPC client
	client := go_client.NewComputeRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var workerSpecs []*go_client.WorkerGroupSpec
	spec :=&go_client.WorkerGroupSpec{
		GroupName: "hehe",
		MaxReplicas: 1,
		MinReplicas: 1,
		Replicas: 1,
		Resource: &go_client.Resource{
			Cpu: 1,
			Memory: 1,
		},
	}
	workerSpecs = append(workerSpecs, spec)

	request := &go_client.CreateComputeRuntimeRequest{
		ComputeRuntime: &go_client.ComputeRuntime{
			Name: opts.name,
			Cloud: go_client.ComputeRuntime_ON_PREM,
			Region: opts.region,
			AvailabilityZone: opts.az,
			HeadGroupSpec: &go_client.HeadGroupSpec{
				Resource: &go_client.Resource{
					// TODO: change to use opts
					Cpu: 1,
					Memory: 1,
				},
			},
			WorkerGroupSepc: workerSpecs,
		},
	}

	r, err := client.CreateComputeRuntime(ctx, request)
	if err != nil {
		log.Fatalf("could not list compute runtime %v", err)
	}

	log.Printf("compute runtime %v is created", r.Id)
	return nil

}

