package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dlsathvik04/OnlyTextBackendGo/api"
	"github.com/dlsathvik04/OnlyTextBackendGo/db/store"
	"github.com/dlsathvik04/OnlyTextBackendGo/models"
	"github.com/dlsathvik04/golibs/hasher"
	"github.com/dlsathvik04/golibs/jwt"
	_ "github.com/lib/pq"

	"github.com/dlsathvik04/golibs/dotenv"
)

// type  Storage interface {

// }

func main() {

	//environment setup
	dotenv.LoadDotEnv(".env", false)
	dbUrl := os.Getenv("DB_URL")
	serverSecret := os.Getenv("SERVER_SECRET")
	serverProvider := os.Getenv("SERVER_PROVIDER")
	serverPort := os.Getenv("PORT")

	//database connection
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	//dependencies
	jwtMan := jwt.NewJWTManager(time.Minute*10, serverSecret, serverProvider)
	hasher := hasher.NewHasher(serverSecret)
	ots := store.NewOnlyTextStorage(db)

	//services
	us := models.NewUserService(ots)
	ps := models.NewPostsService(ots)

	//server setup
	mux := http.NewServeMux()
	server := api.NewOnlyTextServer(mux, ":"+serverPort, hasher, jwtMan, us, ps)
	server.ListenAndServe()
}
