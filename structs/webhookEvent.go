package structs

type entity struct {
	ChatKey string `json:"chatKey"`
	ChatId string `json:"chatId"`
	Id string `json:"id"`
	PlainText string `json:"plainText"`
	CreatedAt int `json:"createdAt"`
	ChatType string `json:"chatType"`
}

type WebhookEvent struct {
	Entity entity `json:"entity"`
}
