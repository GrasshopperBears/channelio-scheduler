package structs

type entity struct {
	ChatKey string `json:"chatKey"`
	Id string `json:"id"`
	PlainText string `json:"plainText"`
	CreatedAt int `json:"createdAt"`
}

type WebhookEvent struct {
	Entity entity `json:"entity"`
}
