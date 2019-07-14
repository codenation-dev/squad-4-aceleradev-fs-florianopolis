import React, { useEffect } from "react";
import { Link } from 'react-router-dom';
import {register} from '../../services/loginService'

const Dashboard = () => {    
    
    useEffect(() => {
        console.log("registro")
        console.log(register("jorge@teste.com.br", "123"));
    });

    return (
        <>
            <h1>Dashboard</h1>
            <Link to="/alertas"><h3>Alerta</h3></Link>
            <Link to="/login"><h3>Login</h3></Link>
        </>
    )
}

export default Dashboard;
  