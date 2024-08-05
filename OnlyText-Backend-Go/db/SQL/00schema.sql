DROP TABLE IF EXISTS Follow;
DROP TABLE IF EXISTS DirectMessages;
DROP TABLE IF EXISTS Likes;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS Posts;
DROP TABLE IF EXISTS Users;

CREATE TABLE Users
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
ALTER TABLE Users ALTER COLUMN verified SET DEFAULT FALSE;

-- Posts TABLE
CREATE TABLE Posts
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
ALTER TABLE Posts ALTER COLUMN public SET DEFAULT TRUE;

CREATE TABLE Follow
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
    ON DELETE CASCADE;

CREATE TABLE Comments
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
    ON DELETE CASCADE;

CREATE TABLE Likes
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
    ON DELETE CASCADE;

CREATE TABLE DirectMessages
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
    ON DELETE CASCADE;