import React, { useState } from 'react'
import {register, login} from '../../services/loginService'
import "./style.css"

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
        <form className="form-signin">
            <div className="title">
                <h1>Login / Register</h1>
            </div>
    
            <div className="inputs">
                <input
                    name="username"
                    className="form-control"
                    placeholder="Username"
                    value={username}
                    onChange={e => setUsername(e.target.value)}
                    required
                />
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
    
            <div className="buttons">
                <button className="button register" onClick={(evt) => registerUser(evt)}>Register</button>
                <button className="button login" onClick={(evt) => loginUser(evt)}>Login</button>
            </div>
        </form>
    )
}

export default Login