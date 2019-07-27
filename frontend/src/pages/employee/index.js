import { useEffect, useState } from React from "react";

import { listEmployee } from '../../services/employeeService';

const renderEmployeeList = list => {
  if (!list || list.length === 0) {
    return (
      <tr>
        <td className='col text-center' colSpan='4'>
          Nenhum funcionário encontrado!!!
        </td>
      </tr>
    )
  }
  console.log(list)
  return list.map((customer, key) => {
    return (
      <tr key={key} className=''>
        <td className='col'>{customer.complete_name}</td>
        <td className='col'>{customer.function}</td>
        <td className='col'>{customer.departament}</td>
        <td className='col text-right'>
          {customer.wage.toLocaleString('pt-BR', { minimumFractionDigits: 2 })}
        </td>
      </tr>
    )
  })
};

const Employee = () => {
  const [nome, setNome] = useState('')
  const [cargo, setCargo] = useState('')
  const [orgao, setOrgao] = useState('')
  const [valor, setValor] = useState(0)
  const [ehCliente, setEhCliente] = useState('A')
  const [campoOrdenacao, setCampoOrdenacao] = useState('complete_name')
  const [ordenacao, setOrdenacao] = useState(false)
  const [page, setPage] = useState(0)
  const [search, setSearch] = useState(false)
  const [employeeList, setEmployeeList] = useState([])

  const [AMBOS, EH_CLIENTE, NAO_CLIENTE] = ['A', 'yes', 'no']

  useEffect(() => {
    const fetchData = async () => {
      const result = await listEmployee(
        nome,
        cargo,
        orgao,
        valor,
        ehCliente,
        page,
        campoOrdenacao,
        ordenacao
      )
      setEmployeeList(result)
    };

    fetchData()
  }, [page, search])

  const nextPage = () => {
    if (employeeList && employeeList.length > 0) {
      setPage(page + 1)
    }
  }

  const previousPage = () => {
    setPage(Math.max(page - 1, 0))
  };

  const pesquisar = () => {
    setPage(0)
    setSearch(!search)
  };

  return (
    <div className='container'>
      <h1>Lista de Funcionários</h1>
      <div className='row'>
        <input
          value={nome}
          onChange={e => setNome(e.target.value)}
          className='form-control col-3 mr-2'
          type='search'
          placeholder='Nome'
          aria-label='Nome'
        />
        <input
          value={cargo}
          onChange={e => setCargo(e.target.value)}
          className='form-control col-3 mr-2'
          type='search'
          placeholder='Cargo'
          aria-label='Cargo'
        />
        <input
          value={orgao}
          onChange={e => setOrgao(e.target.value)}
          className='form-control col-3 mr-2'
          type='search'
          placeholder='Orgão'
          aria-label='Orgão'
        />
        <input
          value={valor}
          onChange={e => setValor(e.target.value)}
          className='form-control col-2 mr-2'
          type='search'
          pattern='[0-9]*'
          inputMode='numeric'
          placeholder='Valor'
          aria-label='Valor'
        />
      </div>
      <div className='row'>
        <div className='col-5 mr-2'>
          <legend className='col-form-label pt-0'>Já é Cliente?</legend>

          <div className='col'>
            <div className='form-check form-check-inline'>
              <label className='radio-inline p-2'>
                <input
                  type='radio'
                  className='form-check-input'
                  name='eh_cliente'
                  value={EH_CLIENTE}
                  checked={ehCliente === EH_CLIENTE}
                  onChange={evt => setEhCliente(evt.target.value)}
                />
                Sim
              </label>
            </div>
            <div className='form-check form-check-inline'>
              <label className='radio-inline p-2'>
                <input
                  type='radio'
                  className='form-check-input'
                  name='eh_cliente'
                  value={NAO_CLIENTE}
                  checked={ehCliente === NAO_CLIENTE}
                  onChange={evt => setEhCliente(evt.target.value)}
                />
                Não
              </label>
            </div>
            <div className='form-check form-check-inline'>
              <label className='radio-inline p-2'>
                <input
                  type='radio'
                  className='form-check-input'
                  name='eh_cliente'
                  value={AMBOS}
                  checked={ehCliente === AMBOS}
                  onChange={evt => setEhCliente(evt.target.value)}
                />
                Ambos
              </label>
            </div>
          </div>
        </div>

        <div className='col-5 mr-2'>
          <legend className='col-form-label pt-0'>Ordenação</legend>
          <label className='input-group mb-3'>
            <select
              className='form-control'
              value={campoOrdenacao}
              onChange={evt => setCampoOrdenacao(evt.target.value)}
            >
              <option value='function'>Cargo</option>
              <option value='complete_name'>Nome</option>
              <option value='departament'>Órgão</option>
              <option value='wage'>Salário</option>
              <option value='relevancia'>Relevância</option>
            </select>

            <select
              className='form-control'
              value={ordenacao}
              onChange={evt => setOrdenacao(evt.target.value)}
            >
              <option value='false'>Ascendente</option>
              <option value='true'>Descendente</option>
            </select>
          </label>
        </div>

        <div className='col-1 mr-2 center'>
          <button
            className='btn btn-info'
            type='button'
            href='#'
            onClick={() => pesquisar()}
          >
            Pesquisar
          </button>
        </div>
      </div>
      <table className='table table-striped table-dark table-hover '>
        <thead>
          <tr>
            <th className='col'>Nome</th>
            <th className='col'>Cargo</th>
            <th className='col'>Orgão</th>
            <th className='col text-right'>Salário</th>
          </tr>
        </thead>
        <tbody>{renderEmployeeList(employeeList)}</tbody>
      </table>

      <div className='d-flex justify-content-center'>
        <nav>
          <ul className='pagination'>
            <li className='page-item'>
              <button
                id='prev'
                className='page-link'
                href='#'
                onClick={() => previousPage()}
              >
                Previous
              </button>
            </li>
            <li className='page-item'>
              <button
                id='next'
                className='page-link'
                href='#'
                onClick={() => nextPage()}
              >
                Next
              </button>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  )
};

export default Employee
