import React from "react";
import { Link } from 'react-router-dom';
import './style.css'

const Sidemenu = () => {
    return (
        <div className="sidemenu">
            <Link to="/"><span>Dashboard</span></Link>
            <Link to="/alertas"><span>Alertas</span></Link>
            <Link to="/"><span>Importar</span></Link>
            <Link to="/customer"><span>Clientes</span></Link>
            <Link to="/employee"><span>Funcionarios</span></Link>
            <Link to="/login"><span>Cadastro de Usuário</span></Link>
        </div>
    )
}

export default Sidemenu;