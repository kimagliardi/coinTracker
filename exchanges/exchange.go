package exchanges

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func CoinStats(coinAddress string) (*Coin, error) {
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
