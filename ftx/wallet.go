package ftx

import (
	"errors"
	"github.com/christsim/ftx-compound/ftx/structs"
	"log"
	"strings"
)

type Balances structs.Balances
type Balance structs.Balance

func (client *FtxClient) GetBalances() (Balances, error) {
	var balances Balances
	resp, err := client._get("wallet/balances", []byte(""))
	if err != nil {
		log.Printf("Error GetBalances", err)
		return balances, err
	}
	err = _processResponse(resp, &balances)
	return balances, err
}

func (client *FtxClient) GetBalance(coin string) (Balance, error) {
	var balances, err = client.GetBalances()

	if err != nil {
		log.Printf("Error GetBalance", err)
		return Balance{}, err
	}

	if !balances.Success {
		log.Printf("Error GetBalance, not success")
		return Balance{}, err
	}

	for _, v := range balances.Result {
		if strings.EqualFold(v.Coin, coin) {
			return Balance{
				Success: true,
				Result:  v,
			}, nil
		}
	}

	return Balance{}, errors.New("Balance currency not found.")
}
