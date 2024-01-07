package types

type UrlsDto struct {
	Urls     []string `json:"urls" binding:"required,unique"`
	Strategy string   `json:"strategy" binding:"required,oneof=first_to_fall at_least_one"`
}

type PingStrategy string

const (
	PingStrategyFirstToFall PingStrategy = "first_to_fall"
	PingStrategyAtLeastOne  PingStrategy = "at_least_one"
)

type PingStatus string

const (
	PingStatusActive     PingStatus = "active"
	PingStatusInactive   PingStatus = "inactive"
	PingStatusTimeout    PingStatus = "timeout"
	PingStatusRedirect   PingStatus = "redirect"
	PingStatusTerminated PingStatus = "terminated"
	PingStatusError      PingStatus = "error"
)
