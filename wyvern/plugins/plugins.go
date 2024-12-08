package plugins

import (
	"net/http"

	"github.com/busyLambda/wyvern/wyvern"
	"github.com/busyLambda/wyvern/wyvern/web"
)

type Plugins struct {
	plugins map[string]Plugin
}

func NewPlugins() *Plugins {
	return &Plugins{
		plugins: map[string]Plugin{},
	}
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

func (p *Plugins) AddPlugin(name string, plugin Plugin) {
	p.plugins[name] = plugin
}

func (p *Plugins) Get(name string) Plugin {
	return p.plugins[name]
}

type Plugin interface {
	Mount(r *http.ServeMux, res *web.Resources)
	Poll(res *web.Resources) func(next wyvern.Handler) wyvern.Handler
	IsPollable() bool
}
