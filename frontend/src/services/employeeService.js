import api from './api';
import {logout} from './loginService';

export const listEmployee = (
    nome, funcao, orgao, wage, ehCliente, page, campoOrdenacao, ordenacao) => {
    
    return api.get(`public_func?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&page=${page}&sortby=${campoOrdenacao}&desc=${ordenacao}`)
    .then(response => response.data)
    .catch(error => [])
}

export const statEmployee = (
  nome, funcao, orgao, wage, ehCliente) => {
  
  return api.get(`public_func/stats?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}`)
  .then(response => response.data)
  .catch(error => [])
}

export const departStatEmployee = (
  nome, funcao, orgao, wage, ehCliente) => {
  
  return api.get(`public_func/distribution?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}`)
  .then(response => response.data)
  .catch(error => [])
}

export default {
  listEmployee,
  statEmployee,
  departStatEmployee,
}