package cmd

import (
	"encoding/json"
	"errors"
	"github.com/cescoferraro/trello"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func weekCheckList(now time.Time, token string, card *trello.Card) ([]CheckList, error) {
	result, err := getChelistsFromCard(card, token)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}
func currentDayCheckList(now time.Time, token string, card *trello.Card) (CheckList, error) {
	var result CheckList
	cheklists, err := getChelistsFromCard(card, token)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	for _, day := range cheklists {
		// hoje
		if day.Name == funcName(now) {
			log.Println(day.Name)
			return day, nil
		}
	}
	return result, errors.New("not found")
}

func currentUserCard(user string, list *trello.List) (*trello.Card, error) {
	var result *trello.Card
	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		log.Fatal(err)
		return result, nil
	}
	for _, card := range cards {
		if card.Name == user {
			return card, nil
		}
	}
	return result, errors.New("not found")
}
func currentWeekList(now time.Time, board *trello.Board) (*trello.List, error) {
	var result *trello.List
	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return result, err
	}
	if len(lists) == 0 {
		return result, err
	}
	result = lists[0]
	log.Println(result.Name)
	return result, nil
}

func d3Board(token string) (*trello.Board, error) {
	client := trello.NewClient(appKey, token)
	d3BoardID := "538f872d42bdfee638a6b839"
	board, err := client.GetBoard(d3BoardID, trello.Defaults())
	if err != nil {
		return board, err
	}
	return board, err
}

func state(s string) string {
	switch s {
	case "imcomplete":
		return ":no_entry:"
	case "complete":
		return ":white_check_mark:"
	default:
		return ":heavy_multiplication_x:"
	}
}

func getChelistsFromCard(card *trello.Card, token string) ([]CheckList, error) {
	var result []CheckList
	resp, err := http.Get(`https://api.trello.com/1/card/` + card.ID + `/checklists?key=` + appKey + `&token=` + token)
	if err != nil {
		// handle err
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return result, err
	}
	return result, nil
}

func funcName(data time.Time) string {
	switch data.Weekday() {
	case time.Sunday:
		return "Sexta"
	case time.Monday:
		return "Segunda"
	case time.Tuesday:
		return ("Terça")
	case time.Wednesday:
		return ("Quarta")
	case time.Thursday:
		return "Quinta"
	case time.Friday:
		return "Sexta"
	case time.Saturday:
		return "Sexta"
	default:
		return "Domingo"
	}
}

type CheckList struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	IDBoard    string       `json:"idBoard"`
	IDCard     string       `json:"idCard"`
	CheckItems []CheckItems `json:"checkItems"`
}

type CheckItems struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
