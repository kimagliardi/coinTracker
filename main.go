package main

import (
	"fmt"

	ex "github.com/kimagliardi/coinTracker/exchanges"
)

func main() {
	fmt.Println()
	coin, _ := ex.CoinStats("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")

	fmt.Println(coin.UpdatedAt)
	fmt.Println(coin.Data.Name)
	fmt.Println(coin.Data.Price, "USD")
	fmt.Println(coin.Data.PriceBNB)
}
