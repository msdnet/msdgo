package main

import (
	"caipiao/cmd"
	"os"

	log "github.com/sirupsen/logrus"
)

// "github.com/kataras/iris"

func main() {
	rootcmd := cmd.NewRootCmd()
	if err := rootcmd.Execute(); err != nil {
		log.WithError(err).Error(err.Error())
		os.Exit(1)
	}

	// app := iris.New()
	// app. Get("/", func(ctx iris.Context) {
	// 	ctx.Writef("hello world %s", "msdnet")
	// })
	// app.Run(iris.Addr(":8080"))
	// app.PartyFunc("/cpanel", func(child iris.Party) {
	// 	child.Get("/", func(ctx iris.Context) {})
	// })
}
