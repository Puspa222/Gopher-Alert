package notifier

type Notifier interface {
	Send(to string, message string) error
	Name() string
}
