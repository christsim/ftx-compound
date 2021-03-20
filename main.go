package main

import (
	"fmt"
	"github.com/christsim/ftx-compound/ftx"
	"os"
	"strconv"
	"strings"
)

func appendIfExists(arr []string, name string) []string {
	if len(os.Getenv(name)) > 0 {
		return append(arr, os.Getenv(name))
	}
	return arr
}

func main() {

	argsWithoutProg := os.Args[1:]

	var args []string
	args = appendIfExists(args, "API_KEY")
	args = appendIfExists(args, "API_SECRET")
	args = appendIfExists(args, "API_SUBACCOUNT")
	args = appendIfExists(args, "COIN")
	args = appendIfExists(args, "YEARLY_RATE")

	if len(args) == 0 {
		args = argsWithoutProg
	}

	if len(args) != 5 {
		fmt.Println("Usage: " + os.Args[0] + " apiKey apiSecret subAccount coin yearlyRate")
		os.Exit(1)
	}

	apiKey := args[0]
	apiSecret := args[1]
	subAccount := args[2]
	coin := strings.ToUpper(args[3])
	yearlyRate, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		fmt.Println("Invalid yearly rate", err)
		os.Exit(1)
	}

	// Connect to main account, use empty string for main account and subaccount name for a subaccount
	// API key, API secret, subaccount name
	client := ftx.New(apiKey, apiSecret, subAccount)

	{
		balance, _ := client.GetBalance(coin)
		fmt.Printf("Balance Before:  Coin: %s, Free: %.3f, Total: %.3f\n", balance.Result.Coin, balance.Result.Free, balance.Result.Total)

		{
			lendingInfo, _ := client.GetLendingInfo(coin)
			fmt.Printf("LendingInfo Before:  Coin: %s, Locked: %.3f, MinRate: %.3f, Offered: %.3f\n", lendingInfo.Result.Coin, lendingInfo.Result.Locked, lendingInfo.Result.Minrate, lendingInfo.Result.Offered)
		}

		if balance.Result.Free != balance.Result.Total {
			sumbitLendingOfferResp, _ := client.SubmitLendingOffer(coin, balance.Result.Total, yearlyRate)
			fmt.Println("SubmitLendingOffer Success: ", sumbitLendingOfferResp.Success)
		}

		{
			lendingInfo, _ := client.GetLendingInfo(coin)
			fmt.Printf("LendingInfo After:  Coin: %s, Locked: %.3f, MinRate: %.3f, Offered: %.3f\n", lendingInfo.Result.Coin, lendingInfo.Result.Locked, lendingInfo.Result.Minrate, lendingInfo.Result.Offered)
		}
	}

	{
		balance, _ := client.GetBalance(coin)
		fmt.Printf("Balance After:  Coin: %s, Free: %.3f, Total: %.3f\n", balance.Result.Coin, balance.Result.Free, balance.Result.Total)
	}
}
