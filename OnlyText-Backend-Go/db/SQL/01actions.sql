INSERT INTO
    Users (username, password, email, firstname, lastname)
VALUES
    ('dlsathvik04', 'one', 'dlsathvik04@gmail.com', 'Lekha Sathvik', 'Devabathini'),
    ('user2', 'two', 'two@gmail.com', 'firstname2', 'lastname2'),
    ('user3', 'two', 'two@gmail.com', 'firstname3', 'lastname3')
RETURNING *;

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
    Follow (userid, follower)
VALUES
    (1, 3),
    (1, 2);

INSERT INTO
    Comments (content, postid, userid)
VALUES
    ('comment on post 1 by user 2', 1, 2),
    ('comment on post 1 by user 3', 1, 3),
    ('comment on post 2 by user 2', 2, 2);

INSERT INTO
    Likes (userid, postid)
VALUES
    (2, 1),
    (3, 1);

INSERT INTO
    DirectMessages (sender, receiver, content) -- not needed when inserting all cols
VALUES
    (1,2,'hello from 1 to 2'),
    (2,1, 'hello from 2 to 1');

SELECT * FROM Users;
SELECT * FROM Posts;

-- get user by id
SELECT username, email,firstname, lastname,emailverified, verified FROM Users WHERE userid=1;
-- get user by username
SELECT username, email,firstname, lastname,emailverified, verified FROM Users WHERE username='dlsathvik04';

-- get post by id
SELECT * FROM Posts WHERE postid=1;

--  get comment by id
SELECT * FROM Comments WHERE commentid=1;

--  get user followers
SELECT
    username,
    email,
    firstname,
    lastname,
    emailverified,
    verified
FROM
    Users
WHERE
    Users.userid IN (
        SELECT
            follower
        FROM
            Follow
        WHERE
            userid = 1
    );

-- get user following
SELECT
    username,
    email,
    firstname,
    lastname,
    emailverified,
    verified
FROM
    Users
WHERE
    Users.userid IN (
        SELECT
            userid
        FROM
            Follow
        WHERE
            follower = 2
    );

-- get user posts
SELECT
    *
FROM
    Posts
WHERE
    userid = 1;

-- get USER liked posts
SELECT
    Likes.postid,
    Likes.created_at,
    Posts.content,
    Posts.userid,
    Posts.public,
    Posts.created_at
FROM
    Likes
    JOIN Posts ON Likes.postid = Posts.postid
WHERE
    Likes.userid = 2;

-- get user comments
SELECT
    *
FROM
    Comments
WHERE
    userid = 2;

-- get number of likes given post
SELECT
    COUNT(*) as likes_count
FROM
    Likes
WHERE
    Likes.postid = 1;

-- get direct messages between sender and reciver in reverse
SELECT
    *
FROM
    DirectMessages
WHERE
    (
        receiver = 1
        AND sender = 2
    );

-- get direct messages between two users in reverse
SELECT
    *
FROM
    DirectMessages
WHERE
    (
        receiver = 1
        AND sender = 2
    )or (receiver=2 AND sender=1);

-- get liked users given post
SELECT
    username,
    email,
    firstname,
    lastname,
    emailverified,
    verified
FROM
    Users
WHERE
    Users.userid IN (
        SELECT
            userid
        FROM
            Likes
        WHERE
            Likes.postid = 1
    );

-- get comments given post
SELECT
    Comments.commentid,
    Comments.content,
    Comments.userid,
    Posts.userid as postowner
FROM
    Comments
    JOIN Posts ON Comments.postid = Posts.postid
WHERE
    Posts.postid = 1;

-- Verify User email
UPDATE Users
    SET
        emailverified = TRUE
    WHERE
        userid=1 RETURNING userid;

-- #TODO get mutual followers given two users
-- #TODO get user feed given userid