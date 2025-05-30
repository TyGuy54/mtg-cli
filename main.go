package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	s "strings"
)

type card_list struct {
	Cards []MTG_card `json:"cards"`
}

type MTG_card struct {
	Name      string `json:"name"`
	Mana_cost string `json:"manaCost"`
	Card_type string `json:"type"`
	Power     string `json:"power"`
	Tougness  string `json:"toughness"`
	Text      string `json:"text"`
}

func format_cards(filtered_data MTG_card) string {
	formated_data := filtered_data.Mana_cost + "\n" +
		filtered_data.Name + "\n" +
		filtered_data.Card_type + "\n" +
		filtered_data.Power + "/" + filtered_data.Tougness + "\n" +
		filtered_data.Text

	return formated_data
}

func filter_response_data(response_body []byte, arg string) MTG_card {
	var result card_list
	var data []MTG_card

	err := json.Unmarshal(response_body, &result)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}


	x := s.Replace(arg, "%20", " ", -1)
	for key, value := range result.Cards {
		if result.Cards[key].Name == x {
			data = append(data, value)
		}
	}

	return data[0]
}

func get_mtg_card(arg string) string {
	split_str := s.Replace(arg, "-", "%20", -1)
	http_request := fmt.Sprintf("https://api.magicthegathering.io/v1/cards?name=%v", split_str)
	response, err := http.Get(http_request)

	if err != nil {
		fmt.Println("No repsposne")
	}

	defer response.Body.Close()

	response_data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("No data was returned")
	}

	// fmt.Println(string(response_data))

	filtered_data := filter_response_data(response_data, split_str)
	formated_data := format_cards(filtered_data)

	return formated_data
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid command")
		os.Exit(1)
	}

	command := os.Args[1]
	arg := os.Args[2:]

	switch command {
	case "-srchf":
		if len(arg) > 0 {
			fmt.Println(get_mtg_card(arg[0]))
		}
	default:
		fmt.Println("command dose not exsit")
	}
}