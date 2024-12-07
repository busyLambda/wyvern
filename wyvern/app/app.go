package app

import (
	"fmt"
	"net/http"

	"github.com/busyLambda/wyvern/wyvern/plugins"
	"github.com/busyLambda/wyvern/wyvern/web"
)

type App struct {
	router    http.ServeMux
	plugins   *plugins.Plugins
	resources *web.Resources
}

func NewApp() *App {
	return &App{
		router:    *http.NewServeMux(),
		resources: &web.Resources{},
	}
}

type Pair = struct {
	String string
	web.Controller
}

func (a *App) Res() *web.Resources {
	return a.resources
}

func (a *App) Group(pattern string, controllers []Pair, plugs []string) {
	for _, pair := range controllers {
		controller := pair.Controller
		c_pattern := pair.String

		full_pattern := fmt.Sprintf("%s%s", pattern, c_pattern)
		handler := web.HandlerFromController(controller)

		middlewares := []Middleware{}

		for _, plugname := range plugs {
			plug := a.plugins.Get(plugname)
			middlewares = append(middlewares, plug.Poll)
		}

		finalHandler := chainMiddleware(http.HandlerFunc(handler), middlewares)
		a.router.HandleFunc(full_pattern, finalHandler)
	}
}

type Handler = func(w http.ResponseWriter, r *http.Request)
type Middleware = func(Handler) Handler

func chainMiddleware(handler Handler, middleware []func(Handler) Handler) Handler {
	if len(middleware) == 0 {
		return handler
	}

	return middleware[0](chainMiddleware(handler, middleware[1:]))
}

func (a *App) Run(addr string) {
	a.plugins.MountAll(&a.router, a.resources)

	if a.plugins != nil {
		fmt.Printf("• Loaded %v plugins ✅\n", a.plugins.Count())
	} else {
		fmt.Println("• No plugins set ❎")
	}

	fmt.Printf("\n✅ Running server at \033[1;34mhttp://locahost%s\033[0m\n", addr)

	http.ListenAndServe(addr, &a.router)
}
