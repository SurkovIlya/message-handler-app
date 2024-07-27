package model

type Message struct {
	ID    uint32
	Value string `json:"value"`
}

type KafkaParams struct {
	Host     string
	Topic    string
	MaxBytes int
}
