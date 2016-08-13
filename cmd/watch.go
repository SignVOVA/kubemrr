package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mkokho/kubemrr/pkg"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch one or several Kubernetes API servers",
	Run: func(cmd *cobra.Command, args []string) {
		ro := pkg.MustRootOptions(cmd)
		pkg.RunWatch(ro)
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)
}
