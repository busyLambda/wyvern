package defaultplugins

import (
	"net/http"

	"github.com/busyLambda/wyvern/wyvern"
	"github.com/busyLambda/wyvern/wyvern/web"
)

type StaticAssetPlugin struct {
	path string
}

func DefaultStaticAssetPlugin() *StaticAssetPlugin {
	return &StaticAssetPlugin{
		path: "./static",
	}
}

func NewStaticAssetPlugin(path string) *StaticAssetPlugin {
	return &StaticAssetPlugin{
		path,
	}
}

func (s *StaticAssetPlugin) Mount(r *http.ServeMux, res *web.Resources) {
	r.Handle("/assets", http.FileServer(http.Dir(s.path)))
}

func (s *StaticAssetPlugin) Poll(res *web.Resources) func(next wyvern.Handler) wyvern.Handler {
	return func(next wyvern.Handler) wyvern.Handler {
		return func(w http.ResponseWriter, r *http.Request) {

		}
	}
}

func (s *StaticAssetPlugin) IsPollable() bool {
	return false
}
