import React, { useEffect } from "react";
import { Link } from 'react-router-dom';
import {getCustomerById} from '../../services/customerService'

const Dashboard = () => {    
    
    useEffect(() => {
        console.log("registro")
        //console.log(getCustomerById(98));
    });

    return (
        <>
            <h1>Dashboard</h1>
            <Link to="/alertas"><h3>Alerta</h3></Link>
            <Link to="/login"><h3>Login</h3></Link>
            <Link to="/employee"><h3>Funcionario</h3></Link>
            <Link to="/customer"><h3>Clientes</h3></Link>
        </>
    )
}

export default Dashboard;
  