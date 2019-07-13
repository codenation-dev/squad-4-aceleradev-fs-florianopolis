import React from 'react';
import { Link } from 'react-router-dom';


const Dashboard = () => {    
    return (
        <>
            <h1>Dashboard</h1>
            <Link to="/alertas"><h3>Alerta</h3></Link>
            <Link to="/login"><h3>Login</h3></Link>
        </>
    )
}

export default Dashboard;
  