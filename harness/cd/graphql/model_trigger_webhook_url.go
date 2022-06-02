package graphql

type WebhookUrl string

func (u *WebhookUrl) String() string {
	return string(*u)
}
