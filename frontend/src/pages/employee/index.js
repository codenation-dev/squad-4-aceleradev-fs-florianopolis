import React from "react";
import { useEffect, useState } from 'react';
import { listEmployee } from '../../services/employeeService';

const renderEmployeeList = (list) => {
    if (list.length == 0){
        return <td className="col text-center" colSpan="4">Nenhum funcionário encontrado!!!</td>
    }
    console.log(list)
    return list.map((customer, key) => {        
        return (
            <tr key={key} className="">
                <td className="col">{customer.complete_name}</td>
                <td className="col">{customer.function}</td>
                <td className="col">{customer.departament}</td>
                <td className="col text-right">{customer.wage.toLocaleString('pt-BR', {"minimumFractionDigits": 2})}</td>
            </tr>
        )
    })
}

const Employee = () => {
    const [nome, setNome] = useState('');
    const [cargo, setCargo] = useState('');
    const [orgao, setOrgao] = useState('');
    const [valor, setValor] = useState(0);
    const [page, setPage] = useState(0);
    const [search, setSearch] = useState(false);
    const [employeeList, setEmployeeList] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const result = await listEmployee(nome, cargo, orgao, valor, page);
            setEmployeeList(result.data);
        };
      
        fetchData();        
    }, [page, search]);

    const nextPage = () => {
        if (employeeList && employeeList.length > 0) {
          setPage(page + 1);
        }
      };
    
    const previousPage = () => {
        setPage(Math.max(page - 1, 0));
      };

    const pesquisar = () => {
        setPage(0)
        setSearch(!search)        
    }
      
    return (
        <div className="container">
            <h1>Lista de Funcionários</h1>
            <div className="row">
                <input value={nome}
                    onChange={e => setNome(e.target.value)}
                    className="form-control col-3 mr-2"
                    type="search"
                    placeholder="Nome"
                    aria-label="Nome" />
                <input value={cargo}
                    onChange={e => setCargo(e.target.value)}
                    className="form-control col-3 mr-2"
                    type="search"
                    placeholder="Cargo"
                    aria-label="Cargo" />
                <input value={orgao}
                    onChange={e => setOrgao(e.target.value)}
                    className="form-control col-3 mr-2"
                    type="search"
                    placeholder="Orgão"
                    aria-label="Orgão" />
                <input value={valor}
                    onChange={e => setValor(e.target.value)}
                    className="form-control col-2 mr-2"
                    type="search"
                    pattern="[0-9]*" inputmode="numeric"
                    placeholder="Valor"
                    aria-label="Valor" />

                <button className="" href="#" onClick={() => pesquisar()}>Pesquisar</button>
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
                    {renderEmployeeList(employeeList)}
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
                    onClick={() => previousPage()}
                    >
                    Previous
                    </button>
                </li>
                <li className="page-item">
                    <button id="next" className="page-link" href="#" onClick={() => nextPage()}>
                    Next
                    </button>
                </li>
                </ul>
            </nav>
            </div>
        </div>
    )
}

export default Employee;
  