package structs

type Block struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

type Message struct {
	Blocks []Block `json:"blocks"`
	Options []string `json:"options"`
}
