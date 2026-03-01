package notifier

import "fmt"

type Registry struct {
	providers map[string]Notifier
}

func NewRegistry(notifiers []Notifier) *Registry {
	provs := make(map[string]Notifier)
	for _, n := range notifiers {
		provs[n.Name()] = n
	}
	return &Registry{providers: provs}
}

func (r *Registry) Get(name string) (Notifier, error) {
	p, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider not found")
	}
	return p, nil
}
