package compute


import "github.com/spf13/cobra"

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <command>",
		Short: "Delete Compute Runtime",
		Long:  ``,
		Annotations: map[string]string{
			"IsCore": "true",
		},
	}

	return cmd
}