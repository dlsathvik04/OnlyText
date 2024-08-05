package store

import (
	"database/sql"
	"fmt"
)

type OnlyTextDB interface {
	GetUserByID(id string)
}

type OnlyTextStorage struct {
	Store Store
}

func (ots *OnlyTextStorage) GetDB() *sql.DB {
	return ots.Store.GetDB()
}

func NewOnlyTextStorage(db *sql.DB) *OnlyTextStorage {
	fmt.Println("######### configurng DB #########")
	store := NewStore(db, []func(*sql.DB) error{
		dropAllTables,
		createUsersTable,
		createPostsTable,
		createFollowTable,
		createCommentsTable,
		createLikesTable,
		createDirectMessagesTable,
	})
	st := OnlyTextStorage{Store: *store}
	fmt.Println("########## SUCCESS #############")
	return &st
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users
(
  userid        SERIAL  PRIMARY KEY,
  username      VARCHAR NOT NULL UNIQUE,
  password      VARCHAR NOT NULL,
  email         VARCHAR NOT NULL,
  firstname     varchar NOT NULL,
  lastname      varchar,
  emailverified boolean,
  verified      boolean
);
ALTER TABLE Users ALTER COLUMN emailverified SET DEFAULT FALSE;
ALTER TABLE Users ALTER COLUMN verified SET DEFAULT FALSE;`)
	if err == nil {
		fmt.Println("created Users table")
	}
	return err
}

func createPostsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE Posts
(
  postid  SERIAL  PRIMARY KEY,
  content VARCHAR NOT NULL,
  userid  INTEGER  NOT NULL,
  public  BOOLEAN NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE Posts
    ADD CONSTRAINT postauthor FOREIGN KEY (userid) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE Posts ALTER COLUMN public SET DEFAULT TRUE;`)
	if err == nil {
		fmt.Println("created Posts table")
	}
	return err
}

func createFollowTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Follow
(
  userid   INTEGER NOT NULL,
  follower INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (userid, follower)
);
ALTER TABLE Follow
    ADD CONSTRAINT following FOREIGN KEY (userid) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE Follow
    ADD CONSTRAINT follower FOREIGN KEY (userid) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;`)
	if err == nil {
		fmt.Println("created Follow table")
	}
	return err
}

func createCommentsTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Comments
(
  commentid SERIAL PRIMARY KEY,
  content   VARCHAR,
  postid    INTEGER NOT NULL,
  userid    INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE Comments
    ADD CONSTRAINT comment_author FOREIGN KEY (userid) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE Comments
    ADD CONSTRAINT commented_on FOREIGN KEY (postid) 
    REFERENCES Posts (postid) 
    ON DELETE CASCADE;`)
	if err == nil {
		fmt.Println("created Comments table")
	}
	return err
}

func createLikesTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Likes
(
  userid  INTEGER NOT NULL,
  postid  INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE Likes
    ADD CONSTRAINT liked_user FOREIGN KEY (userid) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE Likes
    ADD CONSTRAINT post_liked FOREIGN KEY (postid) 
    REFERENCES Posts (postid) 
    ON DELETE CASCADE;`)
	if err == nil {
		fmt.Println("created Likes table")
	}
	return err
}

func createDirectMessagesTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS DirectMessages
(
  sender    INTEGER  NOT NULL,
  receiver  INTEGER  NOT NULL,
  content varchar NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE DirectMessages
    ADD CONSTRAINT dm_sender FOREIGN KEY (sender) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE DirectMessages
    ADD CONSTRAINT dm_reciver FOREIGN KEY (receiver) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;`)
	if err == nil {
		fmt.Println("created DirectMessages table")
	}
	return err
}

func dropAllTables(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS Follow;
DROP TABLE IF EXISTS DirectMessages;
DROP TABLE IF EXISTS Likes;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS Posts;
DROP TABLE IF EXISTS Users;`)
	if err == nil {
		fmt.Println("created DirectMessages table")
	}
	return err
}
