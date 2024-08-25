package api

import (
	"net/http"

	"github.com/dlsathvik04/golibs/jsonresponse"
)

type GreetApiManager struct {
}

func NewGreetApiManager() *GreetApiManager {
	return &GreetApiManager{}
}

func (gam *GreetApiManager) Register(server *http.ServeMux) {
	server.HandleFunc("/", gam.handleRootGreet)
}

func (gam *GreetApiManager) handleRootGreet(w http.ResponseWriter, r *http.Request) {
	jsonresponse.RespondWithJson(w, http.StatusOK, struct{ Greetings string }{"Hello from ots"})
}
