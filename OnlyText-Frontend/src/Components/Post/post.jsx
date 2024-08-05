import { useState } from "react";
import "./post.css";
const Post = ({ content, username, created, likes, comments }) => {
    const [newComment, setNewComment] = useState("");
    const [displayComments, setDisplayComments] = useState(false);
    const comments_list = ["comment 1", "comment 2", "comment 3"];

    const handleAddNewComment = () => {
        console.log(newComment);
    };

    const handleDisplayCommentsToggle = () => {
        if (displayComments) {
            setDisplayComments(false);
        } else {
            setDisplayComments(true);
        }
        console.log(displayComments);
    };
    return (
        <div className="post-root">
            <div className="post-header">
                <div className="user large-bold">
                    <p>{username}</p>
                </div>
                <div className="created">
                    <p>{created}</p>
                </div>
            </div>
            <p>{content}</p>
            <div className="post-footer">
                <button> likes - {likes}</button>
                <button onClick={handleDisplayCommentsToggle}>
                    comments - {comments}
                </button>
            </div>
            {displayComments && (
                <div className="comment-section">
                    <p className="large-bold">Commnets</p>
                    <div className="new-comment">
                        <textarea
                            onChange={(e) => setNewComment(e.target.value)}
                            name="newcomment"
                            placeholder="Add your comment"
                            id="newcomment"
                        ></textarea>
                        <button onClick={handleAddNewComment}>Add</button>
                    </div>
                    {comments_list.map((value, index) => {
                        return <p key={index}>{value}</p>;
                    })}
                </div>
            )}
        </div>
    );
};

export default Post;
