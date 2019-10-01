package cmd

import (
	"github.com/kyokomi/emoji"
	"log"
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
		now := time.Now()
		token := "c6a1826b020982e644517f56d0f29c88d0c65dbb67c00fdab346a420c1bcfc36"
		board, err := d3Board(token)
		if err != nil {
			log.Fatal(err)
			return err
		}
		last, err := currentWeekList(now, board)
		if err != nil {
			log.Fatal(err)
			return err
		}
		card, err := currentUserCard("Cesco", last)
		if err != nil {
			log.Fatal(err)
			return err
		}
		cheklist, err := weekCheckList(now, token, card)
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, kk := range cheklist {
			log.Println(kk.Name)
			for _, ckItem := range kk.CheckItems {
				pizzaMessage := emoji.Sprint(state(ckItem.State) + "" + ckItem.Name + " " + ckItem.ID)
				hey := emoji.Sprint(pizzaMessage)
				log.Println(hey)
			}
		}
		return nil
	},

}

func init() {
	rootCmd.AddCommand(weekCmd)

}
