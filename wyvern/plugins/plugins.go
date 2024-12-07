package plugins

import (
	"net/http"

	"github.com/busyLambda/wyvern/wyvern/web"
)

type Plugins struct {
	plugins map[string]Plugin
}

func (p *Plugins) Count() int {
	return len(p.plugins)
}

func (p *Plugins) MountAll(r *http.ServeMux, res *web.Resources) {
	if p != nil {
		for _, plug := range p.plugins {
			plug.Mount(r, res)
		}
	}
}

func (p *Plugins) Get(name string) Plugin {
	return p.plugins[name]
}

type Plugin interface {
	Mount(r *http.ServeMux, res *web.Resources)
	Poll(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)
}
