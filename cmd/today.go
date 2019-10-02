package cmd

import (
	"github.com/cescoferraro/r2d3/util"
	"github.com/kyokomi/emoji"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verifyIfRequirementesAreMet()
		now := time.Now()
		err := runTOdayCmd(now)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	},
}

func runTOdayCmd(now time.Time) error {
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
	log.Println(last.Name)
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
	for index, ckItem := range cheklist.CheckItems {
		pizzaMessage := emoji.Sprint(spacer(index+1) + strconv.Itoa(index+1) + " " + ckItem.ID + " " + util.State(ckItem.State) + "" + ckItem.Name)
		hey := emoji.Sprint(pizzaMessage)
		log.Println(hey)
	}
	return nil
}

func init() {
	var username string
	var token string
	var client string
	todayCmd.Flags().StringVarP(&client, "client", "c", "", "client, facultative if you have config file")
	_ = viper.BindPFlag("client", todayCmd.Flags().Lookup("client"))
	todayCmd.Flags().StringVarP(&token, "token", "t", "", "token, facultative if you have config file")
	_ = viper.BindPFlag("token", todayCmd.Flags().Lookup("token"))
	todayCmd.Flags().StringVarP(&username, "username", "u", "", "username, facultative if you have config file")
	_ = viper.BindPFlag("username", todayCmd.Flags().Lookup("username"))
	rootCmd.AddCommand(todayCmd)
}
