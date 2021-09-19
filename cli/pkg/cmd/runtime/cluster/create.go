package cluster

import "github.com/spf13/cobra"

func newCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <command>",
		Short: "Create Cluster Runtime",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	return cmd
}