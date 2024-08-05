import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();
    const navigate = useNavigate();
    const handleSubmit = (event) => {
        event.preventDefault();
        console.log({ username, password });
        navigate("/home");
    };
    return (
        <div className="login-root">
            <div className="form ta-l">
                <form action="none">
                    <label htmlFor="username" className="d-block">
                        Username:
                        <input
                            className="d-block"
                            type="text"
                            name="username"
                            id="username"
                            onChange={(e) => setUserName(e.target.value)}
                        />
                    </label>
                    <label htmlFor="password" className="d-block">
                        Password:
                        <input
                            className="d-block"
                            type="password"
                            name="password"
                            id="password"
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </label>
                    <div className="submit_btn ta-c">
                        <button onClick={handleSubmit}>Submit</button>
                    </div>
                    <a href="/register" className="ta-c d-block">
                        Regiser Here
                    </a>
                </form>
            </div>
        </div>
    );
};

export default Login;
