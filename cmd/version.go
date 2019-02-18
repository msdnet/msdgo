package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

func NewVersionCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the version information",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("janus %s\n", version)
		},
	}
}
