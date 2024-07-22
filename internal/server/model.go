package server

type ReceivingBody struct {
	From    string `json:"from"`
	To      string `json:"to"`
	BodyMsg string `json:"bodyMsg"`
}
