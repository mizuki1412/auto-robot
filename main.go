package main

import (
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/mizuki1412/go-core-kit/v2/service/restkit"
	"github.com/spf13/cobra"
)

func main() {
	cli.RootCMD(&cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {

			_ = restkit.Run()
		},
	})
	c1 := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	c1.Flags().String("test", "", "")
	c1.Flags().String("test1", "", "")
	cli.AddChildCMD(c1)

	cli.Execute()
}
