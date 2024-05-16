package telegram

type SendMessage struct {
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
