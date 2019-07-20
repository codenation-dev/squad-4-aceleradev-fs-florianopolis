import React from "react";
import { useEffect, useState } from 'react';
import { listCustomer } from '../../services/customerService';

const renderCustomerList = (list) => {
    if (list.length === 0){
        return <td className="col text-center">Nenhum cliente encontrado!!!</td>
    }

    return list.map((customer, key) => {
        const { id, nome } = customer
        return (
            <tr key={id}>
                <td>{nome}</td>
            </tr>
        )
    })
}

const Customer = () => {
    const [nome, setNome] = useState('');
    const [customerList, setCustomerList] = useState([]);

    useEffect(() => {
        setCustomerList(listCustomer(nome))
    }, [nome]);

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
            </div>
            
            <table className="table table-striped table-dark table-hover ">
                <thead>
                    <tr>
                        <th className="col">Nome</th>
                    </tr>
                </thead>
                <tbody>
                    {renderCustomerList(customerList)}
                </tbody>
            </table>
        </div>
    )
}

export default Customer;
  