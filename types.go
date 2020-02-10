package main

// Config defines the services that notihub provides
type Config struct {
	Telegram bool `yaml:"telegram"`
}

// SendRequest defines the send payload
type SendRequest struct {
	Recipient string `json:"recipient"`
	Msg       string `json:"msg"`
}
