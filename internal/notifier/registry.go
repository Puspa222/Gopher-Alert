package notifier

import "fmt"

type Registry struct {
	providers map[string]Notifier
}

func NewRegistry(console Notifier) *Registry {
	return &Registry{
		providers: map[string]Notifier{
			"console": console,
		},
	}
}

func (r *Registry) Get(name string) (Notifier, error) {
	p, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider not found")
	}
	return p, nil
}
