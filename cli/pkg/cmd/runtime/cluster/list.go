package cluster

import "github.com/spf13/cobra"

func newCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List Cluster Runtime",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	return cmd
}