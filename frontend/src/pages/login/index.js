import React, { Component } from 'react'

class Login extends Component {
    constructor(props) {
        super(props)
    }

    render = () => (
        <form className="form-signin">
            <div className="text-center mb-4">
                <h1 className="h3 mb-3 font-weight-normal">Login / Register</h1>
            </div>
    
            <div className="form-label-group">
                <label htmlFor="inputEmail">Username</label>
                <input
                    name="username"
                    onChange={() => {}}
                    value={''}
                    className="form-control"
                    placeholder="Username"
                    required
                />
            </div>
    
            <div className="form-label-group mt-2">
                <label htmlFor="inputPassword">Password</label>
                <input
                    name="password"
                    onChange={() => {}}
                    value={''}
                    type="password"
                    className="form-control"
                    placeholder="Password"
                    required
                />
            </div>
    
            <div className="mt-5">
                <button className="login btn btn-lg btn-primary btn-block" type="submit">Login</button>
                <button className="register btn btn-lg btn-secondary btn-block" type="submit">Register</button>
            </div>
        </form>
    )
}

export default Login