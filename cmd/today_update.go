package cmd

import (
	"errors"
	"fmt"
	"github.com/cescoferraro/r2d3/util"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("update called")
		if !(viper.GetString("msg") != "" || viper.GetString("state") != "") {
			err := errors.New("need stuff")
			return err
		}
		log.Println(viper.GetString("msg"))
		log.Println(viper.GetString("state"))
		verifyIfRequirementesAreMet()
		if len(args) == 0 || len(args) > 2 {
			return errors.New("need explnation")
		}
		id := args[0]
		msg := viper.GetString("msg")
		log.Println(id, " ", msg)
		now := time.Now()
		board, err := util.D3Board()
		if err != nil {
			return err
		}
		last, err := util.CurrentWeekList(now, board)
		if err != nil {
			return err
		}
		log.Println(last.Name)
		card, err := util.CurrentUserCard(last)
		if err != nil {
			return err
		}
		cheklist, err := util.CurrentDayCheckList(now, card)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println(cheklist.Name)
		for index, ckItem := range cheklist.CheckItems {
			yay, err := calculate(index, id, ckItem.ID)
			if err != nil {
				return err
			}
			if yay {
				//pizzaMessage := emoji.Sprint(spacer(index+1) + strconv.Itoa(index+1) + " " + ckItem.ID + " " + util.State(ckItem.State) + "" + ckItem.Name)
				//hey := emoji.Sprint(pizzaMessage)
				log.Println("current state " + ckItem.State)
				client := &http.Client{}
				state := viper.GetString("state")
				if state == "" {
					state = ckItem.State
				}
				if msg == "" {
					msg = ckItem.Name
				}

				token := viper.GetString("token")
				s := "[" + viper.GetString("client") + "]%20-%20" + strings.Replace(msg, " ", "%20", -1)
				// marshal User to json
				//i := struct {
				//	Name  string
				//	Token string
				//	Key   string
				//}{
				//	Name:  msg,
				//	Key:   util.AppKey,
				//	Token: token,
				//}
				//json, err := json.Marshal(i)
				//if err != nil {
				//	panic(err)
				//}
				req, err := http.NewRequest(http.MethodPut,
					"https://api.trello.com/1"+
						"/cards/"+card.ID+
						"/checklist/"+cheklist.ID+
						"/checkItem/"+ckItem.ID+
						"?key="+util.AppKey+
						"&token="+token+
						"&state="+state+
						"&name="+s, nil)
				log.Println(req)
				resp, err := client.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()
				if err != nil {
					return err
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
					return err
				}
				log.Println(string(body))
			}
		}
		return nil
	},
}

func calculate(index int, id string, ckID string) (bool, error) {
	isHex := bson.IsObjectIdHex(id)
	if isHex {
		return id == ckID, nil
	}
	rindex, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	return rindex == index+1, nil
}

func init() {

	var username string
	var token string
	var msg string
	var client string
	var state string
	updateCmd.Flags().StringVarP(&state, "state", "s", "", "mandatory")
	_ = viper.BindPFlag("state", updateCmd.Flags().Lookup("state"))
	updateCmd.Flags().StringVarP(&msg, "msg", "m", "", "mandatory")
	_ = viper.BindPFlag("msg", updateCmd.Flags().Lookup("msg"))
	updateCmd.Flags().StringVarP(&client, "client", "c", "", "facultative if you have config file")
	_ = viper.BindPFlag("client", updateCmd.Flags().Lookup("client"))
	updateCmd.Flags().StringVarP(&token, "token", "t", "", "facultative if you have config file")
	_ = viper.BindPFlag("token", updateCmd.Flags().Lookup("token"))
	updateCmd.Flags().StringVarP(&username, "username", "u", "", "facultative if you have config file")
	_ = viper.BindPFlag("username", updateCmd.Flags().Lookup("username"))
	todayCmd.AddCommand(updateCmd)

}

//https://trello.com/1/cards/5d92036b0f3a750cb67db11c/checklist/5d92036c56f1894ad4b20fd4/checkItem/5d938207fc99aa8833ae9092
//https://trello.com/1/cards/5d92036b0f3a750cb67db11c/checklist/5d92036c56f1894ad4b20fd4/checkItem/5d938207fc99aa8833ae9092
