package web

import (
	"net/http"
)

func HandlerFromController(c Controller) func(w http.ResponseWriter, r *http.Request) {
	return c.Mount()
}

type MountFunc = func(w *http.ResponseWriter, r *http.Request, res *Resources)

func NewController(mountfunc MountFunc, res *Resources) Controller {
	return Controller{
		MountFunc:        mountfunc,
		reactiveHandlers: make(map[string]func(w http.ResponseWriter, r *http.Request)),
		res:              res,
	}
}

type Controller struct {
	MountFunc        func(w *http.ResponseWriter, r *http.Request, res *Resources)
	reactiveHandlers map[string]func(w http.ResponseWriter, r *http.Request)
	res              *Resources

	// Handler(w http.ResponseWriter, r *http.Request)
	// Mount(w http.ResponseWriter, r *http.Request)
	// AttachReactiveHandler()
}

func (c *Controller) Mount() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c.MountFunc(&w, r, c.res)
	}
}

func (c *Controller) AttachReactiveHandler(name string, handler func(w http.ResponseWriter, r *http.Request, res *Resources)) {
	finalHandler := func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, c.res)
	}
	c.reactiveHandlers[name] = finalHandler
}
