package models

type ASFBotInfo struct {
	Result map[string]ASFBotInfoResult `json:"Result"`
}
type ASFBotInfoResult struct {
	CardsFarmer struct {
		GamesToFarm []struct {
			CardsRemaining int `json:"CardsRemaining"`
		} `json:"GamesToFarm"`
	} `json:"CardsFarmer"`
	IsConnectedAndLoggedOn bool `json:"IsConnectedAndLoggedOn"`
}

// ASFStatsResponse summarized response from ASF API
type ASFStatsResponse struct {
	Total          int `json:"total"`
	Online         int `json:"online"`
	CardsRemaining int `json:"cards_remaining"`
}
