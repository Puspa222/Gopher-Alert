package notifier

type DiscordNotifier struct{}

func NewDiscordNotifier() *DiscordNotifier {
	return &DiscordNotifier{}
}

func (d *DiscordNotifier) Send(to, message string) error {
	return nil
}
