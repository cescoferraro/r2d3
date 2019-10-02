/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/cescoferraro/r2d3/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"time"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verifyIfRequirementesAreMet()
		if len(args) == 0 || len(args) > 1 {
			return errors.New("need explnation")
		}
		now := time.Now()
		board, err := util.D3Board()
		if err != nil {
			log.Fatal(err)
			return err
		}
		last, err := util.CurrentWeekList(now, board)
		if err != nil {
			log.Fatal(err)
			return err
		}
		card, err := util.CurrentUserCard(last)
		if err != nil {
			log.Fatal(err)
			return err
		}
		cheklist, err := util.CurrentDayCheckList(now, card)
		if err != nil {
			log.Fatal(err)
			return err
		}
		token := viper.GetString("token")
		s := "[" + viper.GetString("client") + "]%20-%20" + strings.Replace(args[0], " ", "%20", -1)
		resp, err := http.Post("https://api.trello.com/1/checklists/"+
			cheklist.ID+
			"/checkItems?name="+s+
			"&key="+
			util.AppKey+"&token="+token, "", nil)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer resp.Body.Close()

		err = runTOdayCmd(now)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	},
}

func init() {
	todayCmd.AddCommand(addCmd)

	var username string
	var token string
	var client string
	addCmd.Flags().StringVarP(&client, "client", "c", "", "client, facultative if you have config file")
	_ = viper.BindPFlag("client", addCmd.Flags().Lookup("client"))
	addCmd.Flags().StringVarP(&token, "token", "t", "", "token, facultative if you have config file")
	_ = viper.BindPFlag("token", addCmd.Flags().Lookup("token"))
	addCmd.Flags().StringVarP(&username, "username", "u", "", "username, facultative if you have config file")
	_ = viper.BindPFlag("username", addCmd.Flags().Lookup("username"))
	rootCmd.AddCommand(addCmd)
	//addCmd.PersistentFlags().String("token", "2323", "A help for foo")
}
