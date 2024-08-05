// import React from 'react'

import Post from "../Post/post";
import "./postlist.css";

const PostList = () => {
    const posts = [
        {
            content: "content",
            username: "username",
            created: "created",
            likes: 9,
            comments: 9,
        },
    ];

    return (
        <div className="postlist-root">
            {posts.map((value, index) => (
                <Post
                    key={index}
                    comments={value.comments}
                    likes={value.likes}
                    content={value.content}
                    username={value.username}
                    created={value.created}
                />
            ))}
        </div>
    );
};

export default PostList;
