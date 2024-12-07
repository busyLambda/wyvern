package web

type Resources struct {
	items map[string]interface{}
}

func (r *Resources) Get(name string) interface{} {
	return r.items[name]
}
