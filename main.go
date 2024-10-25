package main

import (
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/spf13/cobra"
	"mizuki/project/auto-robot/mod/chromerob"
)

func main() {
	cli.RootCMD(&cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {
			//_ = restkit.Run()
		},
	})
	c1 := &cobra.Command{
		Use: "web",
		Run: func(cmd *cobra.Command, args []string) {
			chromerob.Start()
		},
	}
	cli.AddChildCMD(c1)
	cli.Execute()
}
