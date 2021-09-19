package cluster

import "github.com/spf13/cobra"

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <command>",
		Short: "Delete Cluster Runtime",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	return cmd
}