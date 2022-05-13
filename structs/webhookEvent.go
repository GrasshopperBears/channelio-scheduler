package structs

type entity struct {
	ChatKey   string `json:"chatKey"`
	ChatId    string `json:"chatId"`
	Id        string `json:"id"`
	PlainText string `json:"plainText"`
	CreatedAt int    `json:"createdAt"`
	ChatType  string `json:"chatType"`
	PersonId  string `json:"personId"`
}

type WebhookEvent struct {
	Entity entity `json:"entity"`
}
