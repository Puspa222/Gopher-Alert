package notifier

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

func (r *Registry) Get(name string) (Notifier, bool) {
	p, ok := r.providers[name]
	return p, ok
}
