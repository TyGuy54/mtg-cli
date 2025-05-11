package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type card_list struct {
	Cards []MTG_card `json:"cards"`
}

type MTG_card struct {
	Name      string `json:"name"`
	Mana_cose string `json:"manaCost"`
	Card_type string `json:"type"`
	Power     string `json:"power"`
	Tougness  string `json:"toughness"`
	Text      string `json:"text"`
}

func filter_response_data(body []byte) MTG_card {
	var result card_list

	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result.Cards[0]
}

func format_cards(filtered_data MTG_card) string {
	formated_data := filtered_data.Name + "\n" +
		filtered_data.Mana_cose + "\n" +
		filtered_data.Card_type + "\n" +
		filtered_data.Power + "\n" +
		filtered_data.Tougness + "\n" +
		filtered_data.Text

	return formated_data
}

func get_mtg_card(arg string) string {
	response, err := http.Get("https://api.magicthegathering.io/v1/cards?&name=" + arg)
	if err != nil {
		fmt.Println("No repsposne")
	}

	defer response.Body.Close()

	response_data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("No data was returned")
	}

	filtered_data := filter_response_data(response_data)
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
