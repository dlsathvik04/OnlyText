import { useState } from "react";

const Register = () => {
    const [firstName, setFirstName] = useState();
    const [lastName, setLastName] = useState();
    const [username, setUserName] = useState();
    const [email, setEmail] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log({ username, password, firstName, lastName, email });
    };
    return (
        <div className="login-root">
            <div className="form ta-l">
                <form action="none">
                    <label htmlFor="firstname" className="d-block">
                        First Name:
                        <input
                            className=" d-block"
                            type="text"
                            name="firstname"
                            id="firstname"
                            onChange={(e) => setFirstName(e.target.value)}
                        />
                    </label>

                    <label htmlFor="firstname" className=" d-block">
                        Last Name:
                        <input
                            className=" d-block"
                            type="text"
                            name="lastname"
                            id="lastname"
                            onChange={(e) => setLastName(e.target.value)}
                        />
                    </label>
                    <label htmlFor="username" className=" d-block">
                        Username:
                        <input
                            className=" d-block"
                            type="text"
                            name="username"
                            id="username"
                            onChange={(e) => setUserName(e.target.value)}
                        />
                    </label>
                    <label htmlFor="email" className=" d-block">
                        Email:
                        <input
                            className=" d-block"
                            type="email"
                            name="email"
                            id="email"
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </label>
                    <label htmlFor="password" className=" d-block">
                        Password:
                        <input
                            className=" d-block"
                            type="password"
                            name="password"
                            id="password"
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </label>
                    <div className="submit_btn ta-c">
                        <button onClick={handleSubmit}>Submit</button>
                    </div>
                    <a href="/login" className="ta-c d-block">
                        Login Here
                    </a>
                </form>
            </div>
        </div>
    );
};

export default Register;
