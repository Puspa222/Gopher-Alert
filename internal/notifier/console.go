package notifier

import "fmt"

type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
	return &ConsoleNotifier{}
}

func (c *ConsoleNotifier) Send(to string, messsage string) error {
	fmt.Printf("\nSending to %s: %s\n", to, messsage)
	return nil
}
