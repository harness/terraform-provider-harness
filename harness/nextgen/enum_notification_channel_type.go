package nextgen

type CVNGNotificationChannelType string

var CVNGNotificationChannelTypes = struct {
	Email     CVNGNotificationChannelType
	Slack     CVNGNotificationChannelType
	PagerDuty CVNGNotificationChannelType
	MsTeams   CVNGNotificationChannelType
}{
	Email:     "Email",
	Slack:     "Slack",
	PagerDuty: "PagerDuty",
	MsTeams:   "MsTeams",
}

var CVNGNotificationChannelTypesSlice = []string{
	CVNGNotificationChannelTypes.Email.String(),
	CVNGNotificationChannelTypes.Slack.String(),
	CVNGNotificationChannelTypes.PagerDuty.String(),
	CVNGNotificationChannelTypes.MsTeams.String(),
}

func (c CVNGNotificationChannelType) String() string {
	return string(c)
}
