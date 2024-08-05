// import React from 'react'

import NavBar from "../../Components/Navbar/navbar";
import PostList from "../../Components/PostList/postlist";
import "./home.css";

const Home = () => {
    return (
        <div className="home-root">
            <div className="nav">
                <NavBar />
            </div>
            <div className="postlist">
                <PostList />
            </div>
        </div>
    );
};

export default Home;
