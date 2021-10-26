package utils

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type BinDetails struct {
	Number struct {
		Length int  `json:"length"`
		Luhn   bool `json:"luhn"`
	} `json:"number"`
	Scheme  string `json:"scheme"`
	Type    string `json:"type"`
	Brand   string `json:"brand"`
	Prepaid bool   `json:"prepaid"`
	Country struct {
		Numeric   string `json:"numeric"`
		Alpha2    string `json:"alpha2"`
		Name      string `json:"name"`
		Emoji     string `json:"emoji"`
		Currency  string `json:"currency"`
		Latitude  int    `json:"latitude"`
		Longitude int    `json:"longitude"`
	} `json:"country"`
	Bank struct {
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"bank"`
}

func GetBinInfo(number int) (BinDetails, error) {

	url := fmt.Sprintf("https://lookup.binlist.net/%d", number)

	method := "GET"

	var binDetail BinDetails

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Info("Request Error", err)
		return BinDetails{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Info("Request Client Error")
		return BinDetails{}, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Info("Request Body Error")
		log.Fatalln(err)
	}

	log.Info(string(b))

	if err := json.Unmarshal(b, &binDetail); err != nil {   // Parse []byte to go struct pointer
		log.Info("Can not unmarshal JSON")
	}

	return binDetail, nil
}
