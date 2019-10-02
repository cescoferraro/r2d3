package cmd

import (
	"github.com/cescoferraro/r2d3/util"
	"github.com/kyokomi/emoji"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// weekCmd represents the week command
var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verifyIfRequirementesAreMet()
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
		cheklist, err := util.WeekCheckList(now, card)
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, kk := range cheklist {
			log.Println(kk.Name)
			for index, ckItem := range kk.CheckItems {
				pizzaMessage := emoji.Sprint(spacer(index+1) + strconv.Itoa(index+1) + " " + util.State(ckItem.State) + "" + ckItem.Name + " " + ckItem.ID)
				hey := emoji.Sprint(pizzaMessage)
				log.Println(hey)
			}
		}
		return nil
	},
}

func spacer(i int) string {

	if i <= 9 {
		return " "
	}
	return ""
}

func init() {
	rootCmd.AddCommand(weekCmd)

}
