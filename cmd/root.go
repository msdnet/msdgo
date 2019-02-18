package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// "github.com/asaskevich/govalidator"

func Run() {
	// b := "12a"
	// isfloat := govalidator.IsFloat(b)
	// fmt.Println("hello", isfloat)
}

var configFile string

// NewRootCmd creates a new instance of the root command
func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	cmd := &cobra.Command{
		Use:   "janus",
		Short: "Janus is an API Gateway",
		Long: `
This is a lightweight API Gateway and Management Platform that enables you
to control who accesses your API, when they access it and how they access it.
API Gateway will also record detailed analytics on how your users are interacting
with your API and when things go wrong.
Complete documentation is available at https://hellofresh.gitbooks.io/janus`,
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Config file (default is $PWD/config.toml)")

	// cmd.AddCommand(NewCheckCmd(ctx))
	cmd.AddCommand(NewVersionCmd(ctx))
	cmd.AddCommand(NewServerCmd(ctx))

	return cmd
}
