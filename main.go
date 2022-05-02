package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	pancakeswap = "https://api.pancakeswap.info/api/v2/tokens/"
)

type Coin struct {
	UpdatedAt int64 `json:"updated_at"`
	Data      struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Price    string `json:"price"`
		PriceBNB string `json:"price_BNB"`
	} `json:"data"`
}

func getCoin(coinAddress string) (*Coin, error) {
	contractAddress := pancakeswap + coinAddress
	resp, err := http.Get(contractAddress)

	if err != nil {
		log.Fatal(err)
		return &Coin{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &Coin{}, err
	}

	coin := Coin{}
	err = json.Unmarshal(body, &coin)

	if err != nil {
		return &Coin{}, err
	}
	return &coin, nil
}

func main() {

	fmt.Println(getCoin("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e8"))
}
