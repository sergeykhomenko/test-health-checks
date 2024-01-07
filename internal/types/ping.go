package types

type PingReport struct {
	Url    string
	Status PingStatus
}

type PingResponse struct {
	Content map[string]any
	Status  int
}
