import React, { Component } from 'react'
import "./style.css"

class Login extends Component {
    constructor(props) {
        super(props)
    }

    render = () => (
        <form className="form-signin">
            <div className="title">
                <h1>Login / Register</h1>
            </div>
    
            <div className="inputs">
                <input
                    name="username"
                    className="form-control"
                    placeholder="Username"
                    required
                />
                <input
                    name="password"
                    type="password"
                    className="form-control"
                    placeholder="Password"
                    required
                />
            </div>
    
            <div className="buttons">
                <button className="button register" type="submit">Register</button>
                <button className="button login" type="submit">Login</button>
            </div>
        </form>
    )
}

export default Login