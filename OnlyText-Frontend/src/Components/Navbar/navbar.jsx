// import React from "react";
import { useState } from "react";
import "./navbar.css";

const NavBar = () => {
    const [dd, setdd] = useState("nav-list");
    const showDropdown = () => {
        if (dd.includes("active")) {
            setdd("nav-list");
        } else {
            setdd("nav-list active");
        }
    };
    return (
        <div className="nav-root">
            <div className="nav-head">
                <a href="/home">
                    <h3>OnlyText</h3>
                </a>
                <button onClick={showDropdown}>Menu</button>
            </div>
            <ul className={dd}>
                <li>
                    <a href="">Profile</a>
                </li>
                <li>
                    <a href="">Connections</a>
                </li>
                <li>
                    <a href="">Messages</a>
                </li>
                <li>
                    <a href="/login">Logout</a>
                </li>
            </ul>
        </div>
    );
};

export default NavBar;
