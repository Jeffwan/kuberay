package cluster

import "github.com/spf13/cobra"

func newCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <command>",
		Short: "Get Cluster Runtime",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	return cmd
}