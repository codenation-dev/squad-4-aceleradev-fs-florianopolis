import React, { useState } from 'react'
import {register, login} from '../../services/loginService'

const Login = (props) => {   
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('')

    const registerUser = (evt) => {
        evt.preventDefault();

        const doRegister = async () => {
            try {
                await register(username, password)
                await login(username, password);
                props.history.push("/");
              } catch (ex) {
                alert(ex.message);
              }            
        };
      
        return doRegister();                 
    }

    const loginUser = (evt) => {
        evt.preventDefault();

        const doLogin = async () => {
            try {
                await login(username, password);
                props.history.push("/");
              } catch (ex) {
                alert(ex.message);
              }            
        };
      
        return doLogin();                 
    }

    return (
        <div className="container h-100">            
            <div className="row h-100">
                <div className="card card-body col-4 mb-2 m-auto">
                    <form className="form-signin">
                        <div className="text-center mb-4">
                            <h1 className="h3 mb-3 font-weight-normal">Login / Register</h1>
                        </div>
                
                        <div className="form-label-group">
                            <label htmlFor="inputEmail">Username</label>
                            <input
                                name="username"
                                className="form-control"
                                placeholder="Username"
                                value={username}
                                onChange={e => setUsername(e.target.value)}
                                required
                            />
                        </div>
                
                        <div className="form-label-group mt-2">
                            <label htmlFor="inputPassword">Password</label>
                            <input
                                name="password"
                                type="password"
                                className="form-control"
                                placeholder="Password"
                                value={password}
                                onChange={e => setPassword(e.target.value)}
                                required
                            />
                        </div>
                
                        <div className="mt-5">
                            <button className="login btn btn-lg btn-primary btn-block" onClick={(evt) => loginUser(evt)}>Login</button>
                            <button className="register btn btn-lg btn-secondary btn-block" onClick={(evt) => registerUser(evt)}>Register</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Login