package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordNotifier struct{}

func NewDiscordNotifier() *DiscordNotifier {
	return &DiscordNotifier{}
}

func (d *DiscordNotifier) Send(to, message string) error {
	payload := map[string]string{
		"content": message,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal Discord payload: %w", err)
	}
	resp, err := http.Post(to, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send Discord message: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Discord API returned status: %s", resp.Status)
	}
	fmt.Printf("[Discord] Message sent to %s\n", to)
	return nil
}

func (d *DiscordNotifier) Name() string {
	return "discord"
}
