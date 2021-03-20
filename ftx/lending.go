package ftx

import (
	"encoding/json"
	"errors"
	"github.com/christsim/ftx-compound/ftx/structs"
	"log"
	"strings"
)

type LendingInfo = structs.LendingInfo
type LendingInfos = structs.LendingInfos
type SubmitLendingOrder = structs.SubmitLendingOrder
type SubmitLendingOrderResponse = structs.SubmitLendingOrderResponse

func (client *FtxClient) GetLendingInfos() (LendingInfos, error) {
	var lendingInfos LendingInfos
	resp, err := client._get("spot_margin/lending_info", []byte(""))
	if err != nil {
		log.Printf("Error GetLendingInfos", err)
		return lendingInfos, err
	}
	err = _processResponse(resp, &lendingInfos)
	return lendingInfos, err
}

func (client *FtxClient) GetLendingInfo(coin string) (LendingInfo, error) {
	var lendingInfos, err = client.GetLendingInfos()

	if err != nil {
		log.Printf("Error GetLendingInfo", err)
		return LendingInfo{}, err
	}

	if !lendingInfos.Success {
		log.Printf("Error GetLendingInfo, not success")
		return LendingInfo{}, err
	}

	for _, v := range lendingInfos.Result {
		if strings.EqualFold(v.Coin, coin) {
			return LendingInfo{
				Success: true,
				Result:  v,
			}, nil
		}
	}

	return LendingInfo{}, errors.New("Balance currency not found.")
}

/**
coin	string	USD
size	number	10.0
rate	number	1e-6
*/
func (client *FtxClient) SubmitLendingOffer(coin string, size float64, yearlyRate float64) (SubmitLendingOrderResponse, error) {
	var submitLendingOrderResponse SubmitLendingOrderResponse

	var hourlyRate = yearlyRate / 100 / 365 / 24

	var submitLendingOrder = SubmitLendingOrder{
		Coin: coin,
		Size: size,
		Rate: hourlyRate,
	}

	requestBody, err := json.Marshal(submitLendingOrder)
	if err != nil {
		log.Printf("Error SubmitLendingOffer", err)
		return submitLendingOrderResponse, err
	}

	resp, err := client._post("spot_margin/offers", requestBody)
	if err != nil {
		log.Printf("Error SubmitLendingOffer", err)
		return submitLendingOrderResponse, err
	}
	err = _processResponse(resp, &submitLendingOrderResponse)
	return submitLendingOrderResponse, err
}
