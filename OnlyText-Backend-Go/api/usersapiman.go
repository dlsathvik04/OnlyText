package api

import (
	"encoding/json"
	"net/http"

	"github.com/dlsathvik04/OnlyTextBackendGo/models"
	"github.com/dlsathvik04/golibs/hasher"
	jsonResponse "github.com/dlsathvik04/golibs/jsonresponse"
	"github.com/dlsathvik04/golibs/jwt"
)

type UserApiManager struct {
	us     models.UserService
	hasher hasher.Hasher
	jwtMan jwt.JWTManager
}

func NewUserApiManger(us models.UserService, hasher hasher.Hasher, jwtMan jwt.JWTManager) *UserApiManager {
	return &UserApiManager{us, hasher, jwtMan}
}

// register user routes here
func (uam *UserApiManager) Register(router *http.ServeMux) {
	router.HandleFunc("POST /users", uam.handleRegister)
	router.HandleFunc("POST /users/login", uam.handleLogin)

}

func (uam *UserApiManager) handleRegister(w http.ResponseWriter, r *http.Request) {

	var u struct {
		Username, Password, Email, Firstname, Lastname string
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusBadRequest, "Invalid Request Body : "+err.Error())
		return
	}
	user, err := uam.us.CreateUser(u.Username, uam.hasher.Hash(u.Password), u.Email, u.Firstname, u.Lastname)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusBadRequest, "Error during db op : "+err.Error())
		return
	}
	jsonResponse.RespondWithJson(w, http.StatusOK, user)

}

func (uam *UserApiManager) handleLogin(w http.ResponseWriter, r *http.Request) {
	var b struct {
		Username, Password string
	}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusBadRequest, "Invalid body structure : "+err.Error())
		return
	}
	user, err := uam.us.GetUserByUsername(b.Username)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if uam.hasher.Compare(b.Password, user.Password) {
		token := uam.jwtMan.GenerateToken(struct {
			userId   int
			verified bool
		}{user.UserId, user.Verified})
		jsonResponse.RespondWithJson(w, http.StatusOK, struct{ Token string }{token})

	} else {
		jsonResponse.RespondWithError(w, http.StatusBadRequest, "Invalid credentials")
	}
}

func (uam UserApiManager) Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := w.Header().Get("authtoken")
		if token == "" {
			jsonResponse.RespondWithError(w, http.StatusUnauthorized, "Authorization failed")
		}
		next(w, r)
	}
}
