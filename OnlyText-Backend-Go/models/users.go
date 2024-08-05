package models

import (
	"github.com/dlsathvik04/OnlyTextBackendGo/db/store"
)

type User struct {
	Userid        int
	Username      string
	Password      string `json:"-"`
	Email         string
	Firstname     string
	Lastname      string
	Emailverified bool
	Verified      bool
}

type UserService interface {
	CreateUser(username, password, email, firstname, lastname string) (User, error)
	GetUserByID(id int) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserFollowers(userid int) ([]User, error)
	GetUserFollowing(userid int) ([]User, error)
	GetLikedUsersByPostID(postid int) ([]User, error)
	VerifyUserEmail(userid int) (int, error)
}

type UsersManager struct {
	ots *store.OnlyTextStorage
}

func NewUserService(ots *store.OnlyTextStorage) UserService {
	return UsersManager{ots}
}

func (um UsersManager) CreateUser(username, password, email, firstname, lastname string) (User, error) {
	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	err := db.QueryRow(`INSERT INTO
    Users (username, password, email, firstname, lastname)
	VALUES
    ($1, $2, $3, $4, $5) RETURNING *;`,
		username, password, email, firstname, lastname,
	).Scan(
		&Userid,
		&Username,
		&Password,
		&Email,
		&Firstname,
		&Lastname,
		&Emailverified,
		&Verified)
	user := User{
		Userid,
		Username,
		Password,
		Email,
		Firstname,
		Lastname,
		Emailverified,
		Verified,
	}
	return user, err

}

func (um UsersManager) GetUserByID(id int) (User, error) {

	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	err := db.QueryRow(`SELECT * FROM Users WHERE userid=$1;`,
		id,
	).Scan(
		&Userid,
		&Username,
		&Password,
		&Email,
		&Firstname,
		&Lastname,
		&Emailverified,
		&Verified)
	user := User{
		Userid,
		Username,
		Password,
		Email,
		Firstname,
		Lastname,
		Emailverified,
		Verified,
	}
	return user, err
}

func (um UsersManager) GetUserByUsername(username string) (User, error) {

	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	err := db.QueryRow(`SELECT * FROM Users WHERE username=$1;`,
		username,
	).Scan(
		&Userid,
		&Username,
		&Password,
		&Email,
		&Firstname,
		&Lastname,
		&Emailverified,
		&Verified)
	user := User{
		Userid,
		Username,
		Password,
		Email,
		Firstname,
		Lastname,
		Emailverified,
		Verified,
	}
	return user, err
}

func (um UsersManager) GetUserFollowers(userid int) ([]User, error) {
	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	rows, err := db.Query(`SELECT * FROM Users
	WHERE
	    Users.userid IN (
	        SELECT
	            follower
	        FROM
	            Follow
	        WHERE
	            userid = $1
	    );`, userid)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()
	res := []User{}
	for rows.Next() {
		err := rows.Scan(&Userid,
			&Username,
			&Password,
			&Email,
			&Firstname,
			&Lastname,
			&Emailverified,
			&Verified)
		if err != nil {
			return []User{}, err
		}
		user := User{
			Userid,
			Username,
			Password,
			Email,
			Firstname,
			Lastname,
			Emailverified,
			Verified,
		}
		res = append(res, user)
	}
	return res, nil
}

func (um UsersManager) GetUserFollowing(userid int) ([]User, error) {
	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	rows, err := db.Query(`SELECT * FROM Users
	WHERE
	    Users.userid IN (
	        SELECT
	            userid
	        FROM
	            Follow
	        WHERE
	            follower = $1
	    );`, userid)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	res := []User{}
	for rows.Next() {
		err := rows.Scan(&Userid,
			&Username,
			&Password,
			&Email,
			&Firstname,
			&Lastname,
			&Emailverified,
			&Verified)
		if err != nil {
			return []User{}, err
		}
		user := User{
			Userid,
			Username,
			Password,
			Email,
			Firstname,
			Lastname,
			Emailverified,
			Verified,
		}
		res = append(res, user)
	}
	return res, nil
}

func (um UsersManager) GetLikedUsersByPostID(postid int) ([]User, error) {
	var Userid int
	var Username string
	var Password string
	var Email string
	var Firstname string
	var Lastname string
	var Emailverified bool
	var Verified bool
	db := um.ots.GetDB()
	rows, err := db.Query(
		`SELECT
		    *
		FROM
		    Users
		WHERE
		    Users.userid IN (
		        SELECT
		            userid
		        FROM
		            Likes
		        WHERE
		            Likes.postid = $1
		    );`,
		postid,
	)

	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	res := []User{}
	for rows.Next() {
		err := rows.Scan(&Userid,
			&Username,
			&Password,
			&Email,
			&Firstname,
			&Lastname,
			&Emailverified,
			&Verified)
		if err != nil {
			return []User{}, err
		}
		user := User{
			Userid,
			Username,
			Password,
			Email,
			Firstname,
			Lastname,
			Emailverified,
			Verified,
		}
		res = append(res, user)
	}
	return res, nil
}

func (um UsersManager) VerifyUserEmail(userid int) (int, error) {
	var userid_up int
	err := um.ots.GetDB().QueryRow(`UPDATE Users
    SET
        emailverified = TRUE
    WHERE
        userid=$1 RETURNING userid;`, userid).Scan(&userid_up)
	return userid_up, err
}
