package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dlsathvik04/OnlyTextBackendGo/db/store"
	"github.com/dlsathvik04/OnlyTextBackendGo/internal/dotenv"
	"github.com/dlsathvik04/OnlyTextBackendGo/models"

	_ "github.com/lib/pq"
)

// type  Storage interface {

// }

func main() {
	dotenv.LoadDotEnv(".env")
	url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	ots := store.NewOnlyTextStorage(db)
	um := models.NewUserService(ots)
	usr, err := um.CreateUser("dlsathvik04", "one", "dlsathvik04@gmail.com", "Lekha Sathvik", "Devabathini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr)
	usr, err = um.CreateUser("dlsathvik04", "one", "dlsathvik04@gmail.com", "Lekha Sathvik", "Devabathini")
	if err != nil {
		fmt.Println(strings.Contains(err.Error(), "duplicate"))
	}
	fmt.Println(usr)
	id, err := um.VerifyUserEmail(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	usr, err = um.GetUserByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr)

	pm := models.NewPostsService(ots)
	post, err := pm.CreatePost(1, "hello", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)

	post, err = pm.CreatePost(1, "hello hello", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)

	posts, err := pm.GetUserPosts(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
	// n := 20
	// fmt.Println(n / 3)

	fmt.Println(LCM(10))
	// fmt.Println(GCD(98, 56))
}
func GCD(n, m int) int {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	if n == 1 || m == 1 {
		return 1
	}
	return GCD(m, n%m)
}

func LCM(n int) int {
	res := 1
	for n > 1 {
		if res%n != 0 {
			res = (res * n) / GCD(res, n)
		}

		n--
	}
	return res
}
