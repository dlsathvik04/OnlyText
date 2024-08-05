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
  content varchar NOT NULL
);
ALTER TABLE DirectMessages
    ADD CONSTRAINT dm_sender FOREIGN KEY (sender) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;
ALTER TABLE DirectMessages
    ADD CONSTRAINT dm_reciver FOREIGN KEY (receiver) 
    REFERENCES Users (userid) 
    ON DELETE CASCADE;

-- CREATE TABLE Users
-- (
--   userid        SERIAL  PRIMARY KEY,
--   username      VARCHAR NOT NULL,
--   password      VARCHAR NOT NULL,
--   email         VARCHAR NOT NULL,
--   firstname     varchar NOT NULL,
--   lastname      varchar,
--   emailverified boolean,
--   verified      boolean
-- );

-- #TODO insert new user

INSERT INTO
    Users (username, password, email, firstname, lastname)
VALUES
    ('dlsathvik04', 'one', 'dlsathvik04@gmail.com', 'Lekha Sathvik', 'Devabathini'),
    ('user2', 'two', 'two@gmail.com', 'firstname2', 'lastname2'),
    ('user3', 'two', 'two@gmail.com', 'firstname3', 'lastname3');

-- #TODO add post from an user
INSERT INTO
    Posts (userid, content)
VALUES
    (1, 'the first post'),
    (1, 'the second post'),
    (1, 'the third post'),
    (2, 'the first post 2'),
    (3, 'the first post 3');

INSERT INTO
    Posts (userid, content, public)
VALUES
    (1, 'the first PRIVATE post', FALSE),
    (1, 'the second PRIVATE post', FALSE),
    (1, 'the third PRIVATE post', FALSE),
    (2, 'the first PRIVATE post 2', FALSE),
    (3, 'the fPRIVATE irst post 3', FALSE);

INSERT INTO 
    Comments ()

-- #TODO
SELECT * FROM Users;
SELECT * FROM Posts;

-- #TODO get user by id
SELECT * FROM Users WHERE userid=1;
-- #TODO get user by username
SELECT * FROM Users WHERE username='dlsathvik04';

-- #TODO get post by id
SELECT * FROM Posts WHERE postid=1;
-- #TODO get comment by id
-- #TODO get user followers
-- #TODO get user following
-- #TODO get user posts
-- #TODO get liked posts
-- #TODO get user comments
-- #TODO get number of likes given post
-- #TODO get direct messages between sender and reciver in reverse chronological order
-- #TODO get liked users given post
-- #TODO get comments given post
-- #TODO get mutual followers given two users
-- #TODO get user feed given userid