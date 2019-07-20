import React from "react";
import { useEffect, useState } from 'react';
import { listEmployee } from '../../services/employeeService';
import Sidemenu from '../../components/sidemenu'
import RenderEmployee from '../../components/renderemployee'
import "./style.css"

const Employee = () => {
    const [nome, setNome] = useState('');
    const [cargo, setCargo] = useState('');
    const [orgao, setOrgao] = useState('');
    const [valor, setValor] = useState(0);
    const [employeeList, setEmployeeList] = useState([]);

    useEffect(() => {
        setEmployeeList(listEmployee(nome, cargo, orgao))
    }, [nome, cargo, orgao, valor]);

    return (
        <>
            <Sidemenu />
            <div className="content employes">
                <div className="title">
                    <h2>Lista de Funcionários</h2>
                </div>
                <div className="search row">
                    <div>
                        <input value={nome}
                            onChange={e => setNome(e.target.value)}
                            className="form-control"
                            type="search"
                            placeholder="Nome"
                            aria-label="Nome" />
                    </div>
                    <div>
                        <input value={cargo}
                            onChange={e => setCargo(e.target.value)}
                            className="form-control"
                            type="search"
                            placeholder="Cargo"
                            aria-label="Cargo" />
                    </div>
                    <div>

                        <input value={orgao}
                            onChange={e => setOrgao(e.target.value)}
                            className="form-control"
                            type="search"
                            placeholder="Orgão"
                            aria-label="Orgão" />
                    </div>
                    <div>
                        <input value={valor}
                            onChange={e => setValor(e.target.value)}
                            className="form-control"
                            type="search"
                            pattern="[0-9]*" inputmode="numeric"
                            placeholder="Valor"
                            aria-label="Valor" />
                    </div>
                </div>
                <table className="table table-striped table-dark table-hover ">
                    <thead>
                        <tr>
                            <th className="col">Nome</th>
                            <th className="col">Cargo</th>
                            <th className="col">Orgão</th>
                            <th className="col text-right">Salário</th>
                        </tr>
                    </thead>
                    <tbody>
                        <RenderEmployee list={employeeList}/>
                    </tbody>
                </table>

                <div className="d-flex justify-content-center">
                <nav>
                    <ul className="pagination">
                    <li className="page-item">
                        <button
                        id="prev"
                        className="page-link"
                        href="#"
                        onClick={() => this.previousPage()}
                        >
                        Previous
                        </button>
                    </li>
                    <li className="page-item">
                        <button id="next" className="page-link" href="#" onClick={() => this.nextPage()}>
                        Next
                        </button>
                    </li>
                    </ul>
                </nav>
                </div>
            </div>
        </>
    )
}

export default Employee;
  