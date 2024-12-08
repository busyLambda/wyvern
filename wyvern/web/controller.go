package web

import (
	"fmt"
	"net/http"

	"github.com/busyLambda/wyvern/wyvern"
)

func HandlerFromController(c Controller) wyvern.Handler {
	return c.Mount()
}

type MountFunc = func(w *http.ResponseWriter, r *http.Request, res *Resources)

func NewController(mountfunc MountFunc, res *Resources) Controller {
	return Controller{
		MountFunc:        mountfunc,
		reactiveHandlers: make(map[string]wyvern.Handler),
		res:              res,
	}
}

type Controller struct {
	MountFunc        func(w *http.ResponseWriter, r *http.Request, res *Resources)
	reactiveHandlers map[string]wyvern.Handler
	res              *Resources

	// Handler(w http.ResponseWriter, r *http.Request)
	// Mount(w http.ResponseWriter, r *http.Request)
	// AttachReactiveHandler()
}

func (c *Controller) Mount() wyvern.Handler {
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

func (c *Controller) SetRes(res *Resources) {
	if c.res != nil {
		c.res = res
	} else {
		warn := `
		[!Warning!]: Attempted to manually set web.Controller.res using SetRes()
		If you are not a developer of the library you shouldn't use this method
		it is only meant to be used once.
		`

		fmt.Print(warn)
	}
}
