package cmd

import (
	"github.com/cescoferraro/r2d3/util"
	"github.com/spf13/cobra"
)

var appKey = "61be8d74f176d471b67764842058d542"

var LoginCMD = &cobra.Command{
	Use:   "login",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		util.Openbrowser("https://trello.com/1/authorize?expiration=1day&name=MyPersonalToken&scope=read&response_type=token&key=" + appKey)
	},
}
