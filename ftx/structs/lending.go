package structs

/**
coin	string	USD
lendable	number	10026.5	additional size you can lend
locked	number	100.0	size either in lending offers or not yet unlocked from lending offers
minRate	number	1e-6	minimum rate at which your offers will lend
offered	number	100.0	size in your lending offers
*/

type LendingInfoResult struct {
	Coin    string  `json:"coin"`
	Locked  float64 `json:"locked"`
	Minrate float64 `json:"minrate"`
	Offered float64 `json:"offered"`
}

type LendingInfo struct {
	Success bool              `json:"success"`
	Result  LendingInfoResult `json:"result"`
}

type LendingInfos struct {
	Success bool                `json:"success"`
	Result  []LendingInfoResult `json:"result"`
}

type SubmitLendingOrder struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
	Rate float64 `json:"rate"`
}

type SubmitLendingOrderResponse struct {
	Success bool  `json:"success"`
	Result  Order `json:"result"`
}
