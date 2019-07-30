import React, { useEffect, useState } from "react";
import {logout} from '../../services/loginService';
import {
  listEmployee,
  statEmployee,
  departStatEmployee
} from "../../services/employeeService";
import { sendEmail } from "../../services/emailService";
import EmployeeBar from "../graphics/EmployeeBar";
import QtdDepart from "../graphics/qtdDepart";
import QtdCargo from "../graphics/qtdCargo";

const renderEmployeeList = result => {  

  if (!result || !result.list || result.list.length === 0) {
    return (
      <tr>
        <td className="col text-center" colSpan="12">
          Nenhum funcionário encontrado!!!
        </td>
      </tr>
    );
  }

  return result.list.map((customer, key) => {
    return (
      <tr key={key}>
        <td scope="row" colSpan="6">{customer.complete_name}</td>
        <td scope="row" colSpan="3">{customer.function}</td>
        <td scope="row" colSpan="2">{customer.departament}</td>
        <td scope="row" className="text-right">
          {customer.wage.toLocaleString("pt-BR", {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
          })}
        </td>
      </tr>
    );
  });
};

const Employee = () => {
  const [nome, setNome] = useState("");
  const [cargo, setCargo] = useState("");
  const [orgao, setOrgao] = useState("");
  const [valor, setValor] = useState(0);
  const [showList, setShowList] = useState(true);
  const [ehCliente, setEhCliente] = useState("A");
  const [relevancia, setRelevancia] = useState(0);
  const [campoOrdenacao, setCampoOrdenacao] = useState("complete_name");
  const [ordenacao, setOrdenacao] = useState(false);
  const [page, setPage] = useState(0);
  const [search, setSearch] = useState(false);
  const [employeeList, setEmployeeList] = useState([]);
  const [employeeStat, setEmployeeStat] = useState([]);
  const [departStat, setDepartStat] = useState([]);

  const [AMBOS, EH_CLIENTE, NAO_CLIENTE] = ["A", "yes", "no"];

  const filterRelevancia = [{id: 0, name: 'Sem Relevância'}, {id: 1, name: '1'}, {id: 2, name: '2'}
  , {id: 3, name: '3'}, {id: 4, name: '4'}, {id: 5, name: '5'}, {id: 6, name: '6'}, {id: 7, name: '7'}
  , {id: 8, name: '8'}, {id: 9, name: '9'}, {id: 10, name: '10'}]

  const makeRelevanciaFilter = (item) => <option key={item.id} value={item.id}>{item.name}</option>

  useEffect(() => {
    const fetchData = async () => {
      const result = await listEmployee(
        nome,
        cargo,
        orgao,
        valor,
        ehCliente,
        relevancia,
        page,
        campoOrdenacao,
        ordenacao
      );
      
      setEmployeeList(result);
    };

    const fetchStatData = async () => {
      const resultList = await listEmployee(
        nome,
        cargo,
        orgao,
        valor,
        ehCliente,
        relevancia,
        page,
        campoOrdenacao,
        ordenacao
      );      
      const result = await statEmployee(nome, cargo, orgao, valor, ehCliente, relevancia);
      const departResult = await departStatEmployee(
        nome,
        cargo,
        orgao,
        valor,
        ehCliente, relevancia
      );

      setEmployeeList({stats: result.info || {}})
      setDepartStat(departResult || {});
      setEmployeeStat(result.stats || {});
    };

    if (showList) {
      fetchData();
    } else {
      fetchStatData();
    }
  }, [page, search, showList]);

  const nextPage = () => {
    if (employeeList.list && employeeList.list.length > 0) {
      setPage(page + 1);
    }
  };

  const previousPage = () => {
    setPage(Math.max(page - 1, 0));
  };

  const pesquisar = () => {
    setPage(0);
    setSearch(!search);
  };

  const renderListInfo = () => {
    return (
      <div>
        <table className="table table-striped table-dark table-hover">
          <thead>
            <tr>
              <th scope="col" colSpan="6">Nome</th>
              <th scope="col" colSpan="3">Cargo</th>
              <th scope="col" colSpan="2">Orgão</th>
              <th scope="col" className="text-right">Salário</th>
            </tr>
          </thead>
          <tbody>{renderEmployeeList(employeeList)}</tbody>
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
                <button
                  id="next"
                  className="page-link"
                  href="#"
                  onClick={() => nextPage()}
                >
                  Next
                </button>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    );
  };
  const renderEstatisticaInfo = () => {
    return (
      <div>
        <EmployeeBar data={employeeStat} />
        <QtdDepart data={departStat.por_orgao} />
        <QtdCargo data={departStat.por_cargo} />
      </div>
    );
  };
  const renderInfo = () =>
    showList ? renderListInfo() : renderEstatisticaInfo();

  return (
    <div className="container">
      
      <nav className="navbar navbar-expand-sm navbar-dark bg-dark">
      <ul className="nav navbar-nav mr-auto">
          <li className="nav-item active">
            <span className="nav-link">Leads</span>
          </li>
        </ul>
        <ul className="nav navbar-nav ml-auto">
          <li className="nav-item">
            <button type="button" className="btn btn-default" onClick={() => logout()}>
              <span className="glyphicon glyphicon-log-out"></span> Sair
            </button>
          </li>
        </ul>
      </nav>

      <div className="card my-2">
        <div className="row">
          <div className="col-10">
        <div className="card-body mr-0 pr-0 row">
          <input
            value={nome}
            onChange={e => setNome(e.target.value)}
            className="form-control col-3 mr-2"
            type="search"
            placeholder="Nome"
            aria-label="Nome"
          />
          <input
            value={cargo}
            onChange={e => setCargo(e.target.value)}
            className="form-control col-3 mr-2"
            type="search"
            placeholder="Cargo"
            aria-label="Cargo"
          />
          <input
            value={orgao}
            onChange={e => setOrgao(e.target.value)}
            className="form-control col-3 mr-2"
            type="search"
            placeholder="Orgão"
            aria-label="Orgão"
          />
          <input
            value={valor}
            onChange={e => setValor(e.target.value)}
            className="form-control col mr-2"
            type="search"
            pattern="[0-9]*"
            inputMode="numeric"
            placeholder="Valor"
            aria-label="Valor"
          />            
        
          <div className="col-4 mr-2">
            <legend className="col-form-label pt-0">Já é Cliente?</legend>

            <div className="col">
              <div className="form-check form-check-inline">
                <label className="radio-inline p-2">
                  <input
                    type="radio"
                    className="form-check-input"
                    name="eh_cliente"
                    value={EH_CLIENTE}
                    checked={ehCliente === EH_CLIENTE}
                    onChange={evt => setEhCliente(evt.target.value)}
                  />
                  Sim
                </label>
              </div>
              <div className="form-check form-check-inline">
                <label className="radio-inline p-2">
                  <input
                    type="radio"
                    className="form-check-input"
                    name="eh_cliente"
                    value={NAO_CLIENTE}
                    checked={ehCliente === NAO_CLIENTE}
                    onChange={evt => setEhCliente(evt.target.value)}
                  />
                  Não
                </label>
              </div>
              <div className="form-check form-check-inline">
                <label className="radio-inline p-2">
                  <input
                    type="radio"
                    className="form-check-input"
                    name="eh_cliente"
                    value={AMBOS}
                    checked={ehCliente === AMBOS}
                    onChange={evt => setEhCliente(evt.target.value)}
                  />
                  Ambos
                </label>
              </div>
            </div>
          </div>

          <div className="col-2 mr-2 ml-0">
          <legend className="col-form-label pt-0">Relevância</legend>
            <label className="input-group mb-3">
              <select
                className="form-control"
                value={relevancia}
                onChange={evt => setRelevancia(evt.target.value)}
              >
                {filterRelevancia.map(makeRelevanciaFilter)}                
              </select>
              </label>
          </div>
          <div className="col-5 mr-2 ml-0">
            <legend className="col-form-label pt-0">Ordenação</legend>
            <label className="input-group mb-3">
              <select
                className="form-control"
                value={campoOrdenacao}
                onChange={evt => setCampoOrdenacao(evt.target.value)}
              >
                <option value="function">Cargo</option>
                <option value="complete_name">Nome</option>
                <option value="departament">Órgão</option>
                <option value="wage">Salário</option>
                <option value="relevancia">Relevância</option>
              </select>

              <select
                className="form-control"
                value={ordenacao}
                onChange={evt => setOrdenacao(evt.target.value)}
              >
                <option value="false">Ascendente</option>
                <option value="true">Descendente</option>
              </select>
            </label>
          </div>          
        </div>
      </div>
      
      <div className="col-2 card-body d-flex flex-column">
      <button
          className="btn btn-info"
          type="button"
          href="#"
          onClick={() => pesquisar()}
        >
          Pesquisar
        </button>
            {showList ? (
              <button
                className="btn btn-info my-2"
                type="button"
                href="#"
                onClick={() => setShowList(!showList)}
              >
                Gráficos
              </button>
            ) : (
              <button
                className="btn btn-info my-2"
                type="button"
                href="#"
                onClick={() => setShowList(!showList)}
              >
                Lista
              </button>
            )}

              <button
                className="btn btn-info"
                type="button"
                href="#"
                onClick={async() => {
                  const result = await sendEmail(nome,
                    cargo,
                    orgao,
                    valor,
                    ehCliente,
                    relevancia)
                  alert("Email enviado com sucesso!")
                  console.log(result)
                }}
              >
                Enviar Email
              </button>
      </div>
      </div>
      </div>

      <div className="card mb-2">
        <div className="card-body row">
          <div className="col-3">
            Quantidade{" "}
            <h2 className="font-weight-bold text-info">
              {(employeeList.stats && employeeList.stats.total) || 0}
            </h2>
          </div>
          <div className="col-3">
            Média Salarial
            <h2 className="font-weight-bold text-info">
              {(
                (employeeList.stats && employeeList.stats.media) ||
                0
              ).toLocaleString("pt-BR", {
                minimumFractionDigits: 2,
                maximumFractionDigits: 2
              })}
            </h2>
          </div>
          <div className="col-3">
            Menor Salário
            <h2 className="font-weight-bold text-info">
              {(
                (employeeList.stats && employeeList.stats.menor) ||
                0
              ).toLocaleString("pt-BR", {
                minimumFractionDigits: 2,
                maximumFractionDigits: 2
              })}
            </h2>
          </div>
          <div className="col-3">
            Maior Salário
            <h2 className="font-weight-bold text-info">
              {(
                (employeeList.stats && employeeList.stats.maior) ||
                0
              ).toLocaleString("pt-BR", {
                minimumFractionDigits: 2,
                maximumFractionDigits: 2
              })}
            </h2>
          </div>
        </div>
      </div>

      {renderInfo()}
    </div>
  );
};

export default Employee;
