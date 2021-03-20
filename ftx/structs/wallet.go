package structs

type BalanceResult struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}

type Balance struct {
	Success bool          `json:"success"`
	Result  BalanceResult `json:"result"`
}

type Balances struct {
	Success bool            `json:"success"`
	Result  []BalanceResult `json:"result"`
}
