package domain

type PushNotificationMessage struct {
	To         []string          `json:"to"`
	Body       string            `json:"body"`
	Data       map[string]string `json:"data,omitempty"`
	Sound      string            `json:"sound,omitempty"`
	Title      string            `json:"title,omitempty"`
	TTLSeconds int               `json:"ttl,omitempty"`
	Expiration int64             `json:"expiration,omitempty"`
	Priority   string            `json:"priority,omitempty"`
	Badge      int               `json:"badge,omitempty"`
	ChannelID  string            `json:"channelId,omitempty"`
}
