package notifier

import "fmt"

type Notifier interface {
	Send(to string, message string) error
}

type ConsoleNotifier struct{}

func (c *ConsoleNotifier) Send(to string, messsage string) error {
	fmt.Printf("\nSending to %s: %s\n", to, messsage)
	return nil
}
