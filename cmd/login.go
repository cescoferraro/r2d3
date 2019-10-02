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
	"fmt"
	"github.com/chzyer/readline"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		askForField("token", viper.GetString("token"))
		askForField("client", viper.GetString("client"))
		askForField("username", viper.GetString("username"))
		s := configFileNameFullPath()
		err := os.Remove(s)
		if err != nil {
			//log.Fatal(err)
		}
		log.Println("===========")
		for _, rule := range viper.AllKeys() {
			log.Println(rule + " " + viper.GetString(rule))
		}
		log.Println("===========")
		_, err = os.Create(s)
		if err != nil {
			log.Fatal(err)
		}
		err = viper.WriteConfigAs(s)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func configFileNameFullPath() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := home + "/" + confiFileName() + ".yaml"
	return s
}

func askForField(field string, value string) {
	rl, err := readline.New("> " + field + " [" + value + "] ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		if line != "" {
			viper.Set(field, line)
			break
		} else {
			viper.Set(field, viper.GetString(field))
			break

		}
	}
	// Do Stuff Here
}

func init() {
	var username string
	var token string
	var client string
	loginCmd.Flags().StringVarP(&client, "client", "c", "", "client, facultative if you have config file")
	_ = viper.BindPFlag("client", loginCmd.Flags().Lookup("client"))
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "token, facultative if you have config file")
	_ = viper.BindPFlag("token", loginCmd.Flags().Lookup("token"))
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "username, facultative if you have config file")
	_ = viper.BindPFlag("username", loginCmd.Flags().Lookup("username"))
	rootCmd.AddCommand(loginCmd)
}
