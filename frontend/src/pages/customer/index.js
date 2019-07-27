import React from "react";
import { useEffect, useState } from 'react';
import { listCustomer } from '../../services/customerService';

const renderCustomerList = (list) => {
    if (!list || list.length === 0){
        return <tr><td className="col text-center">Nenhum cliente encontrado!!!</td></tr>
    }

    return list.map((customer, key) => {
        const { name } = customer
        return (
            <tr key={key}>
                <td>{name}</td>
            </tr>
        )
    })
}

const Customer = () => {
    const [nome, setNome] = useState('');
    const [page, setPage] = useState(0);
    const [search, setSearch] = useState(false);
    const [customerList, setCustomerList] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const result = await listCustomer(nome, page);
            console.log(result);
            setCustomerList(result.data)
        };
      
        fetchData();     
    }, [page, search]);

    const nextPage = () => {
        if (customerList && customerList.length > 0) {
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


                <button className="" href="#" onClick={() => pesquisar()}>Pesquisar</button>
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

export default Customer;
  