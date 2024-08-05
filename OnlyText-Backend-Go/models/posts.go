package models

import (
	"time"

	"github.com/dlsathvik04/OnlyTextBackendGo/db/store"
)

type Post struct {
	Postid     int
	Content    string
	Userid     int
	Public     bool
	Created_at time.Time
}

type PostsService interface {
	CreatePost(userid int, content string, public bool) (Post, error)
	GetPostByID(postid int) (Post, error)
	GetUserPosts(userid int) ([]Post, error)
}

func NewPostsService(ots *store.OnlyTextStorage) PostsService {
	return &PostsManager{ots}
}

type PostsManager struct {
	ots *store.OnlyTextStorage
}

func (um PostsManager) CreatePost(userid int, content string, public bool) (Post, error) {
	var post Post
	db := um.ots.GetDB()
	err := db.QueryRow(
		`INSERT INTO
    		Posts (userid, content, public)
		VALUES
    		($1, $2, $3)
		RETURNING *;`,
		userid, content, public,
	).Scan(
		&post.Postid, &post.Content, &post.Userid, &post.Public, &post.Created_at,
	)
	return post, err
}

func (um PostsManager) GetPostByID(postid int) (Post, error) {
	var post Post
	db := um.ots.GetDB()
	err := db.QueryRow(
		`SELECT * FROM Posts WHERE postid=$1;`,
		postid,
	).Scan(
		&post.Postid, &post.Content, &post.Userid, &post.Public, &post.Created_at,
	)
	return post, err
}

func (um PostsManager) GetUserPosts(userid int) ([]Post, error) {
	db := um.ots.GetDB()
	rows, err := db.Query(
		`SELECT
		    *
		FROM
		    Posts
		WHERE
		    userid = $1;`,
		userid,
	)
	posts := []Post{}
	for rows.Next() {
		var post Post
		rows.Scan(
			&post.Postid, &post.Content, &post.Userid, &post.Public, &post.Created_at,
		)
		posts = append(posts, post)
	}

	defer rows.Close()
	return posts, err
}

// func (um UserManager) GetLikedPostsByUserID(userid int)([]Post, error){

// }
