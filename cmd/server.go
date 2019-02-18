package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

//NewServerCmd Command
func NewServerCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a login server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx)
		},
	}
	return cmd
}

//RunServerStart  begin
func RunServerStart(ctx context.Context) error {
	fmt.Println("server start!")
	initConfig()
	testJson()
	return nil
}

func testJson() {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	log.Println("json:", value.String())
}
