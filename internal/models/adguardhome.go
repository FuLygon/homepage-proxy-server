package models

// AdguardHomeStatsResponse minimalized response from AdGuard Home API
type AdguardHomeStatsResponse struct {
	NumDNSQueries           int     `json:"num_dns_queries"`
	NumBlockedFiltering     int     `json:"num_blocked_filtering"`
	NumReplacedSafeBrowsing int     `json:"num_replaced_safebrowsing"`
	NumReplacedSafeSearch   int     `json:"num_replaced_safesearch"`
	NumReplacedParental     int     `json:"num_replaced_parental"`
	AvgProcessingTime       float64 `json:"avg_processing_time"`
}
