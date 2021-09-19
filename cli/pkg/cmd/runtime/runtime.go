package runtime

import (
	"github.com/ray-project/kuberay/cli/pkg/cmd/runtime/cluster"
	"github.com/ray-project/kuberay/cli/pkg/cmd/runtime/compute"
	"github.com/spf13/cobra"
)

func NewCmdRuntime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime <command>",
		Short: "Manage runtimes(cluster, compute)",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	cmd.AddCommand(cluster.NewCmdClusterRuntime())
	cmd.AddCommand(compute.NewCmdComputeRuntime())

	return cmd
}