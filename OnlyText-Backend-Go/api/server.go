package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dlsathvik04/OnlyTextBackendGo/models"
	"github.com/dlsathvik04/golibs/hasher"
	"github.com/dlsathvik04/golibs/jwt"
)

type OnlyTextServer struct {
	server *http.ServeMux
	addr   string
	hasher hasher.Hasher
	jwtMan jwt.JWTManager
	us     models.UserService
	ps     models.PostsService
}

func NewOnlyTextServer(
	server *http.ServeMux,
	addr string,
	hasher hasher.Hasher,
	jwtMan jwt.JWTManager,
	us models.UserService,
	ps models.PostsService,
) *OnlyTextServer {
	ots := OnlyTextServer{server, addr, hasher, jwtMan, us, ps}
	ots.SetUp()
	return &ots
}

func (ots *OnlyTextServer) ListenAndServe() {
	fmt.Println("Trying to listen on :", ots.addr, "....")
	err := http.ListenAndServe(ots.addr, ots.server)
	if err != nil {
		log.Fatal(err)
	}
}

func (ots *OnlyTextServer) SetUp() {

	//manage root greetings - no authorization --- "/"
	greeter := NewGreetApiManager()
	greeter.Register(ots.server)

	// manage the user route - no authorization --- "/user"
	userApiMan := NewUserApiManger(ots.us, ots.hasher, ots.jwtMan)
	userApiMan.Register(ots.server)

}
